package app

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/grvlle/constellation_wallet/backend/models"
	"github.com/zalando/go-keyring"
)

// ImportWallet is triggered when a user logs into a new Molly wallet for the first time
func (a *WalletApplication) ImportWallet(keystorePath, keystorePassword, keyPassword, alias string) bool {

	alias = strings.ToLower(alias)
	a.wallet = models.Wallet{
		KeyStorePath: keystorePath,
		WalletAlias:  alias,
		Currency:     "USD"}

	if runtime.GOOS == "windows" && !a.javaInstalled() {
		a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
		return false
	}

	if !a.TransactionFinished {
		a.log.Warn("Cannot Import wallet in a pending transaction.")
		a.LoginError("Cannot import a new wallet while there's a pending transaction.")
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

	a.wallet.Address = a.GenerateDAGAddress()
	a.KeyStoreAccess = a.WalletKeystoreAccess(keystorePath, alias)

	if a.KeyStoreAccess {
		if !a.DB.NewRecord(&a.wallet) {
			keystorePasswordHashed, err := a.GenerateSaltedHash(keystorePassword)
			if err != nil {
				a.log.Errorln("Unable to generate password hash. Reason: ", err)
				a.LoginError("Unable to generate password hash.")
				return false
			}

			keyPasswordHashed, err := a.GenerateSaltedHash(keyPassword)
			if err != nil {
				a.log.Errorln("Unable to generate password hash. Reason: ", err)
				a.LoginError("Unable to generate password hash.")
				return false
			}

			if err := a.DB.Create(&a.wallet).Error; err != nil {
				a.log.Errorln("Unable to create database object for the imported wallet. Reason: ", err)
				a.LoginError("Unable to create database object for the imported wallet. Maybe it has already been imported? Try to login.")
				return false
			}

			a.paths.LastTXFile = a.TempFileName("tx-")
			a.paths.PrevTXFile = a.TempFileName("tx-")
			a.paths.EmptyTXFile = a.TempFileName("tx-")

			err = a.createTXFiles()
			if err != nil {
				a.log.Fatalln("Unable to create TX files. Check fs permissions. Reason: ", err)
				a.sendError("Unable to create TX files. Check fs permissions. Reason: ", err)
			}

			if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("Path", models.Path{LastTXFile: a.paths.LastTXFile, PrevTXFile: a.paths.PrevTXFile, EmptyTXFile: a.paths.EmptyTXFile}).Error; err != nil {
				a.log.Errorln("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
				a.sendError("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
			}

			if err := a.DB.Where("wallet_alias = ?", a.wallet.WalletAlias).First(&a.wallet).Updates(&models.Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed}).Error; err != nil {
				a.log.Errorln("Unable to query database object for the imported wallet. Reason: ", err)
				a.LoginError("Unable to query database object for the imported wallet.")
				return false
			}

			a.UserLoggedIn = false
			a.NewUser = true
			a.WalletImported = true
			err = a.initWallet(keystorePath)
			if err != nil {
				a.log.Errorln("Failed to initialize wallet. Reason: ", err)
				// If unable to import previous transactions, remove wallet from DB and logout.
				if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Delete(&a.wallet).Error; err != nil {
					a.log.Errorln("Unable to delete wallet upon failed import. Reason: ", err)
					return false
				}
				return false
			}

			return true

		} else if a.DB.NewRecord(&a.wallet) { // There may already be an existing DB record in some circumstances.
			a.DB.First(&a.wallet)

			a.UserLoggedIn = false
			a.NewUser = false
			a.WalletImported = true
			err := a.initWallet(a.wallet.KeyStorePath)
			if err != nil {
				a.log.Errorln("Faled to initialize wallet. Reason: ", err)
				return false
			}
			return true
		}
	}

	return false
}

// CreateWallet is called when creating a new wallet in frontend component Login.vue
func (a *WalletApplication) CreateOrInitWalletV2(address string) bool {

	//If user comes back to login page, make sure to reset everything
	a.UserLoggedIn = false
	a.NewUser = false

	a.wallet = models.Wallet{
		WalletAlias: address, //PrimaryKey
		Address:     address}

	a.RT.Events.Emit("wallet_keys", a.wallet.Address)

	if a.sendCampaignStatus() {
		a.sendCampaignClaim()
	}

	//Check if any record with WalletAlias exist
	if err := a.DB.Take(&a.wallet).Error; err != nil {

		//Create new record
		if err := a.DB.Create(&a.wallet).Error; err != nil {
			a.log.Errorln("Unable to create database object for new wallet. Reason: ", err)
			a.LoginError("Unable to create new wallet.")
			return false
		}

		a.KeyStoreAccess = true

		if a.KeyStoreAccess {
			a.paths.LastTXFile = a.TempFileName("tx-")
			a.paths.PrevTXFile = a.TempFileName("tx-")
			a.paths.EmptyTXFile = a.TempFileName("tx-")

			err := a.createTXFiles()
			if err != nil {
				a.log.Fatalln("Unable to create TX files. Check fs permissions. Reason: ", err)
				a.sendError("Unable to create TX files. Check fs permissions. Reason: ", err)
			}

			if err := a.DB.Where("wallet_alias = ?", a.wallet.WalletAlias).First(&a.wallet).Update("Path", models.Path{LastTXFile: a.paths.LastTXFile, PrevTXFile: a.paths.PrevTXFile, EmptyTXFile: a.paths.EmptyTXFile}).Error; err != nil {
				a.log.Errorln("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
				a.sendError("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
			}

			a.UserLoggedIn = true
			a.FirstTX = true
			a.NewUser = false

			a.initNewWallet()

			return true
		}
	} else if !a.UserLoggedIn {
		a.UserLoggedIn = true
		a.FirstTX = false
		a.NewUser = false

		err := a.initWallet("")
		if err != nil {
			a.UserLoggedIn = false
		}

		return true
	}

	return true
}

// CreateWallet is called when creating a new wallet in frontend component Login.vue
func (a *WalletApplication) CreateWallet(keystorePath, keystorePassword, keyPassword, alias, label string) bool {

	alias = strings.ToLower(alias)

	if runtime.GOOS == "windows" && !a.javaInstalled() {
		a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
		return false
	}

	if !a.TransactionFinished {
		a.log.Warn("Cannot Create wallet in a pending transaction.")
		a.LoginError("Cannot create a new wallet while there's a pending transaction.")
		return false
	}

	if keystorePath == "" {
		a.LoginError("Please provide a path to store the KeyStore file.")
		return false
	}

	if !a.passwordsProvided(keystorePassword, keyPassword, alias) {
		a.log.Warnln("One or more passwords were not provided.")
		return false
	}

	if alias == "" {
		alias = a.wallet.WalletAlias
	}

	os.Setenv("CL_STOREPASS", keystorePassword)
	os.Setenv("CL_KEYPASS", keyPassword)

	keystorePasswordHashed, err := a.GenerateSaltedHash(keystorePassword)
	if err != nil {
		a.log.Errorln("Unable to generate password hash. Reason: ", err)
		a.sendError("Unable to generate password hash. Reason: ", err)
		return false
	}

	keyPasswordHashed, err := a.GenerateSaltedHash(keyPassword)
	if err != nil {
		a.log.Errorln("Unable to generate password hash. Reason: ", err)
		a.sendError("Unable to generate password hash. Reason: ", err)
		return false
	}

	a.wallet = models.Wallet{
		KeyStorePath:         keystorePath,
		KeystorePasswordHash: keystorePasswordHashed,
		KeyPasswordHash:      keyPasswordHashed,
		WalletAlias:          alias,
		WalletTag:            label}

	if !a.DB.NewRecord(&a.wallet) {
		if err := a.DB.Create(&a.wallet).Error; err != nil {
			a.log.Errorln("Unable to create database object for new wallet. Reason: ", err)
			a.LoginError("Unable to create new wallet. Alias already exists.")
			return false
		}

		if err := a.DB.Where("wallet_alias = ?", alias).First(&a.wallet).Updates(&models.Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed}).Error; err != nil {
			a.log.Errorln("Unable to query database object for new wallet after wallet creation. Reason: ", err)
			a.sendError("Unable to query database object for new wallet after wallet creation. Reason: ", err)
			return false
		}

		err = a.CreateEncryptedKeyStore()
		if err != nil {
			return false
		}

		a.wallet.Address = a.GenerateDAGAddress()

		if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("Address", a.wallet.Address).Error; err != nil {
			a.log.Errorln("Unable to update db object new wallet, with the DAG address. Reason: ", err)
			a.sendError("Unable to update db object new wallet, with the DAG address. Reason. Reason: ", err)
		}
		a.KeyStoreAccess = a.WalletKeystoreAccess(keystorePath, alias)

		if a.KeyStoreAccess {
			a.paths.LastTXFile = a.TempFileName("tx-")
			a.paths.PrevTXFile = a.TempFileName("tx-")
			a.paths.EmptyTXFile = a.TempFileName("tx-")

			err := a.createTXFiles()
			if err != nil {
				a.log.Fatalln("Unable to create TX files. Check fs permissions. Reason: ", err)
				a.sendError("Unable to create TX files. Check fs permissions. Reason: ", err)
			}

			if err := a.DB.Where("wallet_alias = ?", a.wallet.WalletAlias).First(&a.wallet).Update("Path", models.Path{LastTXFile: a.paths.LastTXFile, PrevTXFile: a.paths.PrevTXFile, EmptyTXFile: a.paths.EmptyTXFile}).Error; err != nil {
				a.log.Errorln("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
				a.sendError("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
			}

			a.UserLoggedIn = false
			a.FirstTX = true
			a.NewUser = true

			a.initNewWallet()

			return true
		}
	} else {
		a.LoginError("Unable to create new wallet. Alias already exists.")
	}

	return false
}

// initWallet initializes a new wallet. This is called from login.vue
// only when a new wallet is created.
func (a *WalletApplication) initNewWallet() {

	a.StoreImagePathInDB("faces/face-0.jpg")

	//a.initTransactionHistory()
	//a.passKeysToFrontend()

	a.initTXFromBlockExplorer()

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}
	a.log.Infoln("A New wallet has been created successfully!")
}

// initExistingWallet queries the database for the user wallet and pushes
// the information to the front end components.
func (a *WalletApplication) initWallet(keystorePath string) error {

	if a.NewUser {
		a.initTXFromBlockExplorer()
		a.StoreImagePathInDB("faces/face-0.jpg")
	} else if !a.NewUser {
		a.initTXFromBlockExplorer()
		// a.initTXFromDB()   // Disregard upon import
		a.initTXFilePath() // Update paths from DB.
	}

	a.RT.Events.Emit("wallet_init", a.wallet.TermsOfService, a.wallet.Currency)

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}
	// 	if !a.WidgetRunning.PassKeysToFrontend {
	// 		a.passKeysToFrontend()
	// 	}

	a.log.Infoln("User has logged into the wallet")

	return nil
}

func (a *WalletApplication) initDashboardWidgets() {
	// Initializes a struct containing all Chart Data on the dashboard
	chartData := a.ChartDataInit()

	// Below methods are continously updating the client side modules.
	a.pollTokenBalance()
	a.nodeStats(chartData)
	a.txStats(chartData)
	a.networkStats(chartData)
	a.blockAmount()
	a.pricePoller()

	a.WidgetRunning.DashboardWidgets = true
}

// SavePasswordToKeychain is for saving password to new keychain
func (a *WalletApplication) SavePasswordToKeychain(keystorePassword string) bool {
	return a.saveInfoToKeychain(ServiceLogin, keystorePassword)
}

// SavePhraseandPKeyToKeychain is for saving password to new keychain
func (a *WalletApplication) SavePhraseandPKeyToKeychain(seedPhrase, privateKey string) bool {
	return a.saveInfoToKeychain(ServiceSeed, seedPhrase) && a.saveInfoToKeychain(ServicePKey, privateKey)
}

// InitKeychains is for initializing of all your existing keychains
func (a *WalletApplication) InitKeychains() bool {
	user, err := user.Current()
	if err != nil {
		a.log.Warnln("Unable to detect your username.")
		a.LoginError("Unable to detect your username.")
		return false
	}

	account := user.Username

	a.deleteKeychain(ServiceLogin, account)
	a.deleteKeychain(ServiceSeed, account)
	a.deleteKeychain(ServicePKey, account)

	return true
}

func (a *WalletApplication) saveInfoToKeychain(service, info string) bool {
	user, err := user.Current()
	if err != nil {
		a.log.Warnln("Unable to detect your username.")
		a.LoginError("Unable to detect your username.")
		return false
	}

	account := user.Username

	err = keyring.Set(service, account, info)
	if err != nil {
		a.log.Warnln("Unable to create your keychain.")
		a.LoginError("Unable to create your keycahin.")
		return false
	}

	return true
}

func (a *WalletApplication) deleteKeychain(service, account string) error {
	_, err := keyring.Get(service, account)
	if err != nil {
		return nil
	}
	err = keyring.Delete(service, account)
	if err != nil {
		a.log.Warnln("Unable to delete your existing keychain: ", service)
		a.LoginError("Unable to delete your existing keychain.")
		return err
	}
	return nil
}

func (a *WalletApplication) createTXFiles() error {
	files := []string{a.paths.LastTXFile, a.paths.PrevTXFile, a.paths.EmptyTXFile}

	for _, f := range files {
		file, err := os.Create(f)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// ImportKeys is called from the frontend to use a file dialog to select key file.
func (a *WalletApplication) ImportKeys() error {
	filename := a.RT.Dialog.SelectFile()
	a.log.Info("Path to keys that user wants to import: " + filename)
	return nil
}

// ExportKeys is called from the frontend to use a file dialog to select directory
// where user wants to export the keys to.
func (a *WalletApplication) ExportKeys() error {
	filename := a.RT.Dialog.SelectDirectory()
	a.log.Info("File user wants to save to: " + filename)
	return nil
}

func (a *WalletApplication) initTXFilePath() {
	paths := &a.wallet.Path
	if err := a.DB.Model(&a.wallet).Where("alias = ?", a.wallet.WalletAlias).Association("Path").Find(&paths).Error; err != nil {
		a.log.Fatal("Unable to initialize TX filepaths. Reason: ", err)
		a.sendError("Unable to initialize TX filepaths. Reason: ", err)
		return
	}
	if a.wallet.Path.LastTXFile == "" && a.wallet.Path.PrevTXFile == "" {
		a.log.Fatal("Unable to initialize TX filepaths. Both are empty after DB query.")
	}
	a.paths.LastTXFile = a.wallet.Path.LastTXFile
	a.paths.PrevTXFile = a.wallet.Path.PrevTXFile
	a.paths.EmptyTXFile = a.wallet.Path.EmptyTXFile

}

func (a *WalletApplication) initTXFromDB() {

	if err := a.DB.Model(&a.wallet).Where("alias = ?", a.wallet.WalletAlias).Association("TXHistory").Find(&a.wallet.TXHistory).Error; err != nil {
		a.log.Error("Unable to initialize historic transactions from DB. Reason: ", err)
		a.sendError("Unable to initialize historic transactions from DB. Reason: ", err)
		return
	}

	allTX := []models.TXHistory{}

	for i, tx := range a.wallet.TXHistory {
		allTX = append([]models.TXHistory{tx}, allTX...) // prepend to reverse list for FE

		if a.wallet.TXHistory[i].Status == "Pending" {
			a.TxPending(a.wallet.TXHistory[i].Hash)
		}
	}
	a.RT.Events.Emit("update_tx_history", allTX)
}

func (a *WalletApplication) resyncTXHistory() ([]models.TXHistory, map[string]bool) {

	walletTxHistory := a.DB.Model(&a.wallet).Where("alias = ?", a.wallet.WalletAlias).Association("TXHistory")

	if err := walletTxHistory.Find(&a.wallet.TXHistory).Error; err != nil {
		a.log.Error("Unable to initialize historic transactions from DB. Reason: ", err)
		a.sendError("Unable to initialize historic transactions from DB. Reason: ", err)
		return nil, nil
	}

	//a.log.Infof("Before Count - ", walletTxHistory.Count())

	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	allTX := []models.TXHistory{}
	TXHistory := a.wallet.TXHistory

	for _, tx := range TXHistory {
		if !encountered[tx.Hash] {
			// Record this element as an encountered element.
			encountered[tx.Hash] = true
			// prepend to reverse list for FE
			allTX = append([]models.TXHistory{tx}, allTX...)

			if err := walletTxHistory.Append(tx).Error; err != nil {
				a.log.Errorln("Unable to update the DB record with the new TX. Reason: ", err)
				a.sendError("Unable to update the DB record with the new TX. Reason: ", err)
			}

			if tx.Status == "Pending" {
				a.TxPending(tx.Hash)
			}

			//a.log.Infoln("Keeping tx in db - " + tx.Hash + ", " + tx.Status)
		} else {
			a.log.Infoln("Duplicate tx found - " + tx.Hash)
			walletTxHistory.Delete(&tx)
		}
	}

	//a.log.Infof("After Count - ", walletTxHistory.Count())

	return allTX, encountered
}

func (a *WalletApplication) initTXFromBlockExplorer() {

	hasError := true
	beTxList := []models.TXHistory{}

	for i := 0; i < 10 && hasError; i++ {

		if i > 0 {
			time.Sleep(2 * time.Second)
		}

		hasError = false

		a.log.Info("Sending API call to block explorer on: " + a.Network.BlockExplorer.URL + "/address/" + a.wallet.Address + "/transaction")

		resp, err := http.Get(a.Network.BlockExplorer.URL + "/address/" + a.wallet.Address + "/transaction")
		if err != nil {
			hasError = true
			continue
		}
		defer resp.Body.Close()

		if resp.Body == nil {
			a.log.Info("Unable to detect any previous transactions.")
			return
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			hasError = true
			continue
		}

		ok, error := a.verifyAPIResponse(bodyBytes)
		// Blockexplorer returns below string when no previous transactions are found
		if !ok && error != "Cannot find transactions for sender" {
			a.log.Errorln("API returned the following error", error)
			hasError = true
			continue
		}

		// If no previous transactions for imported wallet - proceed
		if !ok && error == "Cannot find transactions for sender" {
			a.log.Info("Unable to detect any previous transactions.")
			return
		}

		err = json.Unmarshal(bodyBytes, &beTxList)
		if err != nil {
			hasError = true
		}
	}

	if hasError {
		a.log.Errorln("Unable to fetch TX history from block explorer.")
		a.sendError("Unable to fetch TX history from block explorer.", errors.New("unreachable Block Explorer"))
		return
	}

	// Reverse order
	for i := len(beTxList)/2 - 1; i >= 0; i-- {
		opp := len(beTxList) - 1 - i
		beTxList[i], beTxList[opp] = beTxList[opp], beTxList[i]
	}

	a.log.Infof("Successfully collected %d previous transactions. Updating local state...", len(beTxList))

	walletTxHistory := a.DB.Model(&a.wallet).Where("alias = ?", a.wallet.WalletAlias).Association("TXHistory")
	allTX, encountered := a.resyncTXHistory()

	for _, tx := range beTxList {

		if !encountered[tx.Hash] {

			t, _ := time.Parse(time.RFC3339, tx.Timestamp)

			txData := models.TXHistory{
				Amount:   tx.Amount,
				Sender:   tx.Sender,
				Receiver: tx.Receiver,
				Fee:      tx.Fee,
				Hash:     tx.Hash,
				TS:       t.In(t.Local().Location()).Format("Jan _2 15:04:05"),
				Status:   "Complete",
				Failed:   false,
			}

			if err := walletTxHistory.Append(&txData).Error; err != nil {
				a.log.Errorln("Unable to update the DB record with the new TX. Reason: ", err)
				a.sendError("Unable to update the DB record with the new TX. Reason: ", err)
			}

			//a.RT.Events.Emit("new_transaction", txData)
			allTX = append([]models.TXHistory{txData}, allTX...) // prepend to reverse list for FE

			//a.log.Infoln("TX added to db - " + tx.Hash)
		} else {
			//a.log.Infoln("TX already in db - " + tx.Hash)
		}

	}

	a.RT.Events.Emit("update_tx_history", allTX)
}

// PassKeysToFrontend emits the keys to the settings.Vue component on a
// 5 second interval
func (a *WalletApplication) passKeysToFrontend() {

	if a.wallet.Address != "" {
		go func() {
			for {
				a.RT.Events.Emit("wallet_keys", a.wallet.Address)
				time.Sleep(5 * time.Second)
			}
		}()
		a.WidgetRunning.PassKeysToFrontend = true
	} else {
		a.WidgetRunning.PassKeysToFrontend = false
	}
}

func (a *WalletApplication) passwordsProvided(keystorePassword, keyPassword, alias string) bool {
	if keystorePassword == "" {
		a.LoginError("Please provide a Key Store password.")
		return false
	} else if keyPassword == "" {
		a.LoginError("Please provide a Key Password.")
		return false
	} else if alias == "" {
		a.LoginError("An Alias has not been provided.")
		return false
	}
	return true
}

// GetTokenBalance polls and parses the token balance of a wallet and returns it as a float64.
func (a *WalletApplication) GetTokenBalance() (float64, error) {
	a.log.Debug("Contacting mainnet on: " + a.Network.URL + a.Network.Handles.Balance + " Sending the following payload: " + a.wallet.Address)

	resp, err := http.Get(a.Network.URL + a.Network.Handles.Balance + a.wallet.Address)
	if err != nil {
		a.log.Warnln("Failed to send HTTP request. Reason: ", err)
		return 0, err
	}
	if resp == nil {
		err = errors.New("received empty response from the Token Balance API")
		a.log.Warnln("Unable to update token balance. Reason: ", err)
		return 0, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		a.log.Warnln("Unable to update token balance. Reason: ", err)
		return 0, err
	}

	if string(bodyBytes) == "null" { // null body is returned when a wallet has zero balance.
		return 0.0, nil
	}

	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return 0, err
	}

	a.log.Infoln(result) //TEMP

	s := result["balance"]
	if s == "" {
		s = "0" // Empty means zero
	}

	b, ok := s.(float64)
	if !ok {
		err = errors.New("unable to parse balance")
		a.log.Warnln("Unable to update token balance. Reason: ", err)
		return 0, err
	}
	a.log.Infoln("Parsed the following balance: ", s)

	a.log.Infoln("Returning the following balance: ", b)

	return b, nil
}
