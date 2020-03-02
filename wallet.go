package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

func (a *WalletApplication) TempPrintCreds() {
	fmt.Println("address: ", a.wallet.Address, "alias: ", a.wallet.WalletAlias, "keyStorePass: ", os.Getenv("CL_STOREPASS"), "keyPass: ", os.Getenv("CL_KEYPASS"), "key: ", a.paths.EncPrivKeyFile)
}

/* Database Model is located in models.go */

func (a *WalletApplication) ImportWallet(keystorePath, keystorePassword, keyPassword, alias string) bool {

	alias = strings.ToLower(alias)

	if runtime.GOOS == "windows" && !a.javaInstalled() {
		a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
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

	a.wallet = Wallet{
		KeyStorePath: keystorePath,
		WalletAlias:  alias}

	a.wallet.Address = a.GenerateDAGAddress()
	a.KeyStoreAccess = a.WalletKeystoreAccess()

	if a.KeyStoreAccess {
		if !a.DB.NewRecord(&a.wallet) {
			keystorePasswordHashed, err := a.GenerateSaltedHash(keystorePassword)
			if err != nil {
				a.log.Errorf("Unable to generate password hash. Reason: ", err)
				a.LoginError("Unable to generate password hash.")
				return false
			}

			keyPasswordHashed, err := a.GenerateSaltedHash(keyPassword)
			if err != nil {
				a.log.Errorf("Unable to generate password hash. Reason: ", err)
				a.LoginError("Unable to generate password hash.")
				return false
			}

			if err := a.DB.Create(&a.wallet).Error; err != nil {
				a.log.Errorf("Unable to create database object for the imported wallet. Reason: ", err)
				a.LoginError("Unable to create database object for the imported wallet. Maybe it has already been imported? Try to login.")
				return false
			}

			a.paths.LastTXFile = a.TempFileName("tx-", "")
			a.paths.PrevTXFile = a.TempFileName("tx-", "")
			a.paths.EmptyTXFile = a.TempFileName("tx-", "")

			err = a.createTXFiles()
			if err != nil {
				a.log.Fatalln("Unable to create TX files. Check fs permissions. Reason: ", err)
				a.sendError("Unable to create TX files. Check fs permissions. Reason: ", err)
			}

			if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("Path", Path{LastTXFile: a.paths.LastTXFile, PrevTXFile: a.paths.PrevTXFile, EmptyTXFile: a.paths.EmptyTXFile}).Error; err != nil {
				a.log.Errorf("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
				a.sendError("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
			}

			if err := a.DB.Where("wallet_alias = ?", a.wallet.WalletAlias).First(&a.wallet).Updates(&Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed}).Error; err != nil {
				a.log.Errorf("Unable to query database object for the imported wallet. Reason: ", err)
				a.LoginError("Unable to query database object for the imported wallet.")
				return false
			}

			a.UserLoggedIn = false
			a.NewUser = true
			err = a.initWallet(keystorePath)
			if err != nil {
				a.log.Errorln("Faled to initialize wallet. Reason: ", err)
				return false
			}

			return true

		} else if a.DB.NewRecord(&a.wallet) { // There may already be an existing DB record in some circumstances.
			a.DB.First(&a.wallet)

			a.UserLoggedIn = false
			a.NewUser = false
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

// CreateUser is called when creating a new wallet in frontend component Login.vue
func (a *WalletApplication) CreateWallet(keystorePath, keystorePassword, keyPassword, alias string) bool {

	alias = strings.ToLower(alias)

	if runtime.GOOS == "windows" && !a.javaInstalled() {
		a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
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
		a.log.Errorf("Unable to generate password hash. Reason: ", err)
		a.sendError("Unable to generate password hash. Reason: ", err)
		return false
	}

	keyPasswordHashed, err := a.GenerateSaltedHash(keyPassword)
	if err != nil {
		a.log.Errorf("Unable to generate password hash. Reason: ", err)
		a.sendError("Unable to generate password hash. Reason: ", err)
		return false
	}

	a.wallet = Wallet{
		KeyStorePath:         keystorePath,
		KeystorePasswordHash: keystorePasswordHashed,
		KeyPasswordHash:      keyPasswordHashed,
		WalletAlias:          alias}

	if !a.DB.NewRecord(&a.wallet) {
		if err := a.DB.Create(&a.wallet).Error; err != nil {
			a.log.Errorf("Unable to create database object for new wallet. Reason: ", err)
			a.LoginError("Unable to create new wallet. Alias already exists.")
			return false
		}

		if err := a.DB.Where("wallet_alias = ?", alias).First(&a.wallet).Updates(&Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed}).Error; err != nil {
			a.log.Errorf("Unable to query database object for new wallet after wallet creation. Reason: ", err)
			a.sendError("Unable to query database object for new wallet after wallet creation. Reason: ", err)
			return false
		}

		a.CreateEncryptedKeyStore()

		a.wallet.Address = a.GenerateDAGAddress()

		if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("Address", a.wallet.Address).Error; err != nil {
			a.log.Errorf("Unable to update db object new wallet, with the DAG address. Reason: ", err)
			a.sendError("Unable to update db object new wallet, with the DAG address. Reason. Reason: ", err)
		}
		a.KeyStoreAccess = a.WalletKeystoreAccess()

		if a.KeyStoreAccess {
			a.paths.LastTXFile = a.TempFileName("tx-", "")
			a.paths.PrevTXFile = a.TempFileName("tx-", "")
			a.paths.EmptyTXFile = a.TempFileName("tx-", "")

			err := a.createTXFiles()
			if err != nil {
				a.log.Fatalln("Unable to create TX files. Check fs permissions. Reason: ", err)
				a.sendError("Unable to create TX files. Check fs permissions. Reason: ", err)
			}

			if err := a.DB.Where("wallet_alias = ?", a.wallet.WalletAlias).First(&a.wallet).Update("Path", Path{LastTXFile: a.paths.LastTXFile, PrevTXFile: a.paths.PrevTXFile, EmptyTXFile: a.paths.EmptyTXFile}).Error; err != nil {
				a.log.Errorf("Unable to update the DB record with the tmp tx-paths. Reason: ", err)
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
	a.passKeysToFrontend()

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}
	a.log.Infoln("A New wallet has been created successfully!")
}

// initExistingWallet queries the database for the user wallet and pushes
// the information to the front end components.
func (a *WalletApplication) initWallet(keystorePath string) error {

	a.paths.EncPrivKeyFile = keystorePath

	if a.NewUser {
		err := a.initTXFromBlockExplorer()
		if err != nil {
			return err
		}
		a.StoreImagePathInDB("faces/face-0.jpg")
	} else if !a.NewUser {
		a.initTXFromDB()   // Disregard upon import
		a.initTXFilePath() // Update paths from DB.
	}

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}
	if !a.WidgetRunning.PassKeysToFrontend {
		a.passKeysToFrontend()
	}

	a.log.Infoln("User has logged into the wallet")
	return nil

}

func (a *WalletApplication) initDashboardWidgets() {
	// Initializes a struct containing all Chart Data on the dashboard
	chartData := a.ChartDataInit()

	// Below methods are continously updating the client side modules.
	a.nodeStats(chartData)
	a.txStats(chartData)
	a.networkStats(chartData)
	a.pollTokenBalance()
	a.blockAmount()
	a.tokenAmount()
	a.pricePoller()

	a.WidgetRunning.DashboardWidgets = true
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
	transactions := &a.wallet.TXHistory
	if err := a.DB.Model(&a.wallet).Where("alias = ?", a.wallet.WalletAlias).Association("TXHistory").Find(&transactions).Error; err != nil {
		a.log.Error("Unable to initialize historic transactions from DB. Reason: ", err)
		a.sendError("Unable to initialize historic transactions from DB. Reason: ", err)
		return
	}

	for i := range a.wallet.TXHistory {
		if !a.wallet.TXHistory[i].Failed {
			a.RT.Events.Emit("new_transaction", &a.wallet.TXHistory[i]) // Pass the tx to the frontend as a new transaction.
		}
	}

}

// initTXFromBlockExplorer is called when an existing wallet is imported.
func (a *WalletApplication) initTXFromBlockExplorer() error {
	a.log.Info("Sending API call to block explorer on: " + a.Network.BlockExplorer.URL + a.Network.BlockExplorer.Handles.CollectTX + a.wallet.Address)

	resp, err := http.Get(a.Network.BlockExplorer.URL + a.Network.BlockExplorer.Handles.CollectTX + a.wallet.Address)
	if err != nil {
		a.log.Errorln("Failed to send HTTP request. Reason: ", err)
		a.sendError("Unable to collect previous TX's from blockexplorer. Please check your internet connection. Reason: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.Body != nil {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			a.log.Fatal(err)
		}
		ok, error := a.verifyAPIResponse(bodyBytes)
		// Blockexplorer returns below string when no previous transactions are found
		if !ok && error != "Cannot find transactions for sender" {
			a.log.Errorln("API returned the following error", error)
			// If unable to import previous transactions, remove wallet from DB and logout.
			if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Delete(&a.wallet).Error; err != nil {
				a.log.Errorln("Unable to delete wallet upon failed import. Reason: ", err)
				return err
			}
			a.LoginError("The wallet import failed. Please check your internet connection and try again.")
			return err
		}

		// If no previous transactions for imported wallet - proceed
		if !ok && error == "Cannot find transactions for sender" {
			a.log.Info("Unable to detect any previous transactions.")
			return nil
		}

		allTX := []TXHistory{}

		err = json.Unmarshal(bodyBytes, &allTX)
		if err != nil {
			a.log.Errorln("Unable to fetch TX history from block explorer. Reason: ", err)
			a.sendError("Unable to fetch TX history from block explorer. Reason: ", err)
			return err
		}

		go func() {
			for _, tx := range allTX {
				txData := &TXHistory{
					Amount: tx.Amount,
					Sender: tx.Sender,
					Fee:    tx.Fee,
					Hash:   tx.Hash,
					TS:     time.Now().Format("Mon Jan _2 15:04:05 2006"),
					Failed: false,
				}
				a.storeTX(txData)
				a.RT.Events.Emit("new_transaction", txData)

				ptx := a.loadTXFromFile(a.paths.PrevTXFile)
				ltx := a.loadTXFromFile(a.paths.LastTXFile)

				ptxObj, ltxObj := a.convertToTXObject(ptx, ltx)

				a.rebuildImportChain(tx.Amount, tx.Fee, tx.Receiver, ptxObj, ltxObj)
			}
		}()
	} else {
		a.log.Info("Unable to detect any previous transactions.")
		return nil
	}
	return nil

}

// PassKeysToFrontend emits the keys to the settings.Vue component on a
// 5 second interval
func (a *WalletApplication) passKeysToFrontend() {
	if a.paths.EncPrivKeyFile != "" && a.wallet.Address != "" {
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
	if a.paths.EncPrivKeyFile == "" {
		a.LoginError("Please provide a valid path to your KeyStore file.")
		return false
	} else if keystorePassword == "" {
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
