package app

import (
	"os"
	"runtime"
	"strings"

	"github.com/grvlle/constellation_wallet/backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginError takes a string and pushes it to the login screen as an errror
func (a *WalletApplication) LoginError(errMsg string) {
	if errMsg != "" {
		a.RT.Events.Emit("login_error", errMsg, true)
	}
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
		a.log.Warn("Cannot login to another wallet during a pending transaction.")
		a.LoginError("Cannot login to another wallet during a pending transaction.")
		return false
	}

	if keystorePath == "" {
		a.LoginError("Please provide a path to the KeyStore file.")
		return false
	}

	if !a.passwordsProvided(keystorePassword, keyPassword, alias) {
		a.log.Warnln("One or more passwords were not provided.")
		return false
	}

	os.Setenv("CL_STOREPASS", keystorePassword)
	os.Setenv("CL_KEYPASS", keyPassword)

	a.wallet.WalletAlias = alias

	if err := a.DB.First(&a.wallet, "wallet_alias = ?", alias).Error; err != nil {
		a.log.Errorln("Unable to query database object for existing wallet. Reason: ", err)
		a.LoginError("Access Denied. Alias not found.")
		return false
	}

	if !a.WalletKeystoreAccess() {
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

		// os.Setenv("CL_STOREPASS", keystorePassword)
		// os.Setenv("CL_KEYPASS", keyPassword)

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

// LogOut will reset the wallet UI and clear the wallet objects
func (a *WalletApplication) LogOut() bool {
	if a.TransactionFinished {
		a.UserLoggedIn = false
		a.wallet = models.Wallet{}
		return true
	}
	a.sendWarning("Cannot log out while transaction is processing. Please try again.")
	return false
}

// ImportKey is called from the frontend when browsing the fs for a keyfile
func (a *WalletApplication) ImportKey() string {
	a.paths.EncPrivKeyFile = a.RT.Dialog.SelectFile()
	if a.paths.EncPrivKeyFile == "" {
		a.LoginError("Access Denied. No key path detected.")
		return ""
	}

	if a.paths.EncPrivKeyFile[len(a.paths.EncPrivKeyFile)-4:] != ".p12" {
		a.LoginError("Access Denied. Not a key file.")
		return ""
	}
	a.log.Info("Path to imported key: " + a.paths.EncPrivKeyFile)
	return a.paths.EncPrivKeyFile
}

// SelectDirToStoreKey is called from the FE when creating a new keyfile
func (a *WalletApplication) SelectDirToStoreKey() string {
	a.paths.EncPrivKeyFile = a.RT.Dialog.SelectSaveFile()

	if len(a.paths.EncPrivKeyFile) <= 0 {
		a.LoginError("No valid path were provided. Please try again.")
		return ""
	}
	if a.paths.EncPrivKeyFile[len(a.paths.EncPrivKeyFile)-4:] != ".p12" {
		a.paths.EncPrivKeyFile = a.paths.EncPrivKeyFile + ".p12"
		return a.paths.EncPrivKeyFile
	}
	return a.paths.EncPrivKeyFile
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
