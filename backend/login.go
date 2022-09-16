package app

import (
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/grvlle/constellation_wallet/backend/models"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/bcrypt"
)

// LoginError takes a string and pushes it to the login screen as an errror
func (a *WalletApplication) LoginError(errMsg string) {
	if errMsg != "" {
		a.RT.Events.Emit("login_error", errMsg, true)
	}
}

// LoginKeychain - login with the password of the existing keychain
func (a *WalletApplication) LoginKeychain(keystorePassword string) string {
	user, err := user.Current()

	if err != nil {
		a.log.Warnln("Unable to detect your username.")
		a.LoginError("Unable to detect your username.")
		return ""
	}

	account := user.Username

	secret, err := keyring.Get(ServiceLogin, account)

	a.log.Warnln("secret - " + secret)

	if err != nil {
		a.log.Warnln("Your login keychain doesn't exist.")
		a.LoginError("Your login keychain doesn't exist.")
		return ""
	}

	if secret != keystorePassword {
		a.log.Warnln("Invalid password")
		a.LoginError("Invalid password")
		return ""
	}

	pkey, err := keyring.Get(ServicePKey, account)

	a.log.Warnln("pkey - " + pkey)

	if err == nil {
		return pkey
	}

	return "" //Unable to find a private key for this account, must import one
}

// Login is called from the FE when a user logs in with a wallet object
// already in the DB
func (a *WalletApplication) Login(keystorePath, keystorePassword, keyPassword, alias string) bool {

	alias = strings.ToLower(alias)

	if runtime.GOOS == "windows" && !a.javaInstalled() {
		a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
		return false
	}

	if !a.TransactionFinished {
		a.log.Warnln("Cannot login to another wallet during a pending transaction.")
		a.LoginError("Cannot login to another wallet during a pending transaction.")
		return false
	}

	if keystorePath == "" {
		a.log.Warnln("The provided path to the keystore file is empty.")
		a.LoginError("Please provide a path to the KeyStore file.")
		return false
	}

	if !a.passwordsProvided(keystorePassword, keyPassword, alias) {
		a.log.Warnln("One or more passwords were not provided.")
		a.LoginError("One or more passwords were not provided.")
		return false
	}

	os.Setenv("CL_STOREPASS", keystorePassword)
	os.Setenv("CL_KEYPASS", keyPassword)

	if err := a.DB.First(&a.wallet, "wallet_alias = ?", alias).Error; err != nil {
		a.log.Errorln("Unable to query database object for existing wallet. Reason: ", err)
		return a.ImportWallet(keystorePath, keystorePassword, keyPassword, alias)
	}

	if !a.WalletKeystoreAccess(keystorePath, alias) {
		a.LoginError("Access Denied. Please make sure that you have typed in the correct credentials.")
		return false
	}

	if !a.NewUser {
		a.DB.Model(&a.wallet).Update("KeystorePath", keystorePath)
		a.log.Infoln("PrivateKey path: ", keystorePath)
	}

	// Check password strings against salted hashes stored in DB. Also make sure KeyStore has been accessed.
	if a.CheckAccess(keystorePassword, a.wallet.KeystorePasswordHash) && a.CheckAccess(keyPassword, a.wallet.KeyPasswordHash) && a.KeyStoreAccess {
		a.UserLoggedIn = true
	} else {
		a.UserLoggedIn = false
		a.LoginError("Access Denied. Please make sure that you have typed in the correct credentials.")
	}

	if a.UserLoggedIn && a.KeyStoreAccess && !a.NewUser {

		err := a.initWallet(keystorePath)
		if err != nil {
			a.UserLoggedIn = false
		}
	}

	a.NewUser = false

	return a.UserLoggedIn
}

// CheckTermsOfService is called from the FE to check the termsOfService has been accepted
func (a *WalletApplication) CheckTermsOfService() bool {
	return a.wallet.TermsOfService
}

// LogOut will reset the wallet UI and clear the wallet objects
func (a *WalletApplication) LogOut() bool {
	if a.TransactionFinished {
		a.UserLoggedIn = false
		a.wallet = models.Wallet{}
		a.Network.URL = MainnetLoadBalancerURL // Reset to default network upon every logout
		a.Network.BlockExplorer.URL = MainnetBlockExplorerURL
		//a.log.Infoln("Connected to: Main Constellation Network\n", a.Network.URL+"\n", a.Network.BlockExplorer.URL)
		return true
	}
	a.sendWarning("Cannot log out while transaction is processing. Please try again.")
	return false
}

// ImportKey is called from the frontend when browsing the fs for a keyfile
func (a *WalletApplication) ImportKey() string {
	var keyfile = a.RT.Dialog.SelectFile()
	if keyfile == "" {
		a.LoginError("Access Denied. No key path detected.")
		return ""
	}

	if keyfile[len(keyfile)-4:] != ".p12" {
		a.LoginError("Access Denied. Not a P12 file.")
		return ""
	}
	a.log.Info("Path to imported key: " + keyfile)
	return keyfile
}

func (a *WalletApplication) BrowseJsonFile() string {
	var keyfile = a.RT.Dialog.SelectFile()
	if keyfile == "" {
		a.LoginError("Access Denied. No key path detected.")
		return ""
	}

	if keyfile[len(keyfile)-5:] != ".json" {
		a.LoginError("Access Denied. Not a JSON file.")
		return ""
	}
	a.log.Info("Path to imported key: " + keyfile)
	return keyfile
}

// SelectNetwork is triggered from the login page and will change the network an loadbalancer endpoints
func (a *WalletApplication) SelectNetwork(testnet bool) bool {
	if testnet {
		// Ceres Test Network
		a.Network.URL = TestnetLoadBalancerURL
		a.Network.BlockExplorer.URL = TestnetBlockExplorerURL
		a.log.Infoln("Connected to: Ceres Test Network\n", a.Network.URL+"\n", a.Network.BlockExplorer.URL)
	}

	if !testnet {
		a.Network.URL = MainnetLoadBalancerURL
		a.Network.BlockExplorer.URL = MainnetBlockExplorerURL
		a.log.Infoln("Connected to: Main Constellation Network\n", a.Network.URL+"\n", a.Network.BlockExplorer.URL)
	}

	// Clear old TX history before initialization
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Association("TXHistory").Delete(&models.TXHistory{}).Error; err != nil {
		a.log.Errorln("Unable to update the DB record with the new TX. Reason: ", err)
		a.sendError("Unable to update the DB record with the new TX. Reason: ", err)
	}

	a.RT.Events.Emit("update_tx_history", []models.TXHistory{}) // Clear TX history

	// Force re-initialization and token balance update upon network switch
	a.initWallet(a.wallet.KeyStorePath)
	err := a.updateTokenBalance()
	if err != nil {
		a.log.Errorln("unable to manually update token balance upon network switch: ", err)
	}

	return testnet
}

// SelectDirToStoreKey is called from the FE when creating a new keyfile
func (a *WalletApplication) SelectDirToStoreKey() string {

	var keyfile = a.RT.Dialog.SelectSaveFile()

	if len(keyfile) <= 0 {
		a.LoginError("No valid path were provided. Please try again.")
		return ""
	}
	if keyfile[len(keyfile)-4:] != ".p12" {
		keyfile = keyfile + ".p12"
		return keyfile
	}
	return keyfile
}

// GenerateSaltedHash converts plain text to a salted hash
func (a *WalletApplication) GenerateSaltedHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes[:])
	return hash, nil
}

// CheckAccess verifies that the user has entered the correct password
func (a *WalletApplication) CheckAccess(password, passwordHash string) bool {
	err := a.Compare(password, passwordHash)
	if err != nil {
		a.log.Warnln("User tried to login with the wrong credentials!")
		return false
	}
	a.log.Infoln("Password check OK")
	return true
}

// Compare compares a string with a salted hash
func (a *WalletApplication) Compare(s, hash string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
