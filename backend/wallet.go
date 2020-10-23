package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/grvlle/constellation_wallet/backend/models"
	"github.com/zalando/go-keyring"
)


func (a *WalletApplication) MigrateWallet(keystorePath, keystorePassword, keyPassword, alias string) bool {

    alias = strings.ToLower(alias)

    if runtime.GOOS == "windows" && !a.javaInstalled() {
        a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
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

    return a.produceKeystoreMigrateV2(keystorePath, alias);

}

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

	a.wallet = models.Wallet{
		WalletAlias: address,     //PrimaryKey
		Address: address}

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

func (a *WalletApplication) CreateKeyStoreFile(fileName, jsonKey string) string {

    var fullFilePath = a.paths.HomeDir + "/" + fileName;

		// Check if a file with the same name exists
		i := 0
		suffix := ""
		for {
			if (i != 0) {
				suffix = strconv.Itoa(i)
			}
			if _, err := os.Stat(fullFilePath + suffix + "-key.json"); os.IsNotExist(err) {
				break
			}
			i ++
		}

		fullFilePath += suffix + "-key.json"
    err := WriteStringToFile(fullFilePath, jsonKey)
    if err != nil {
        a.log.Errorln("Unable to write file. Reason: ", err)
        a.sendError("Unable to write file. Reason: ", err)
        return ""
    }

    return fullFilePath
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
	a.passKeysToFrontend()

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}
	a.log.Infoln("A New wallet has been created successfully!")
}

// initExistingWallet queries the database for the user wallet and pushes
// the information to the front end components.
func (a *WalletApplication) initWallet(keystorePath string) error {

	if a.NewUser {
		err := a.initTXFromBlockExplorer()
		if err != nil {
			return err
		}
		a.StoreImagePathInDB("faces/face-0.jpg")
	} else if !a.NewUser {
		err := a.initTXFromBlockExplorer()
		if err != nil {
			return err
		}
		// a.initTXFromDB()   // Disregard upon import
		a.initTXFilePath() // Update paths from DB.
	}

	a.RT.Events.Emit("wallet_init", a.wallet.TermsOfService, a.wallet.Currency)

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
	if (err != nil) {
		a.log.Warnln("Unable to create your keychain.")
		a.LoginError("Unable to create your keycahin.")
		return false
	}
	
	return true
}

func (a *WalletApplication) deleteKeychain(service, account string) error {
	_, err := keyring.Get(service, account)
	if (err != nil) {
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
	a.RT.Events.Emit("update_tx_history", allTX) // Pass the tx to the frontend as a new transaction.
}

// initTXFromBlockExplorer is called when an existing wallet is imported.
func (a *WalletApplication) initTXFromBlockExplorer() error {
	a.log.Info("Sending API call to block explorer on: " + a.Network.BlockExplorer.URL + "/address/" + a.wallet.Address + "/transaction")

	// resp, err := http.Get(a.Network.BlockExplorer.URL + a.Network.BlockExplorer.Handles.CollectTX + a.wallet.Address)
	resp, err := http.Get(a.Network.BlockExplorer.URL + "/address/" + a.wallet.Address + "/transaction")
	if err != nil {
		a.log.Errorln("Failed to send HTTP request. Reason: ", err)
		a.LoginError("Unable to collect previous transactions from blockexplorer.")
		return err
	}
	defer resp.Body.Close()

	if resp.Body != nil {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			a.LoginError("Unable to collect previous transactions from blockexplorer. Try again later.")
			a.log.Errorln("Unable to collect previous transactions from blockexplorer. Reason: ", err)
			return err
		}

		ok, error := a.verifyAPIResponse(bodyBytes)
		// Blockexplorer returns below string when no previous transactions are found
		if !ok && error != "Cannot find transactions for sender" {
			a.log.Errorln("API returned the following error", error)
			a.LoginError("The wallet import failed. Please check your internet connection and try again.")
			return errors.New(error)
		}

		// If no previous transactions for imported wallet - proceed
		if !ok && error == "Cannot find transactions for sender" {
			a.log.Info("Unable to detect any previous transactions.")
			return nil
		}

		allTX := []models.TXHistory{}

		err = json.Unmarshal(bodyBytes, &allTX)
		if err != nil {
			a.log.Errorln("Unable to fetch TX history from block explorer. Reason: ", err)
			a.sendError("Unable to fetch TX history from block explorer. Reason: ", err)
			return err
		}

		// Reverse order
		for i := len(allTX)/2 - 1; i >= 0; i-- {
			opp := len(allTX) - 1 - i
			allTX[i], allTX[opp] = allTX[opp], allTX[i]
		}

		a.log.Infof("Successfully collected %d previous transactions. Updating local state...", len(allTX))

		for i, tx := range allTX {

			t, _ := time.Parse(time.RFC3339, tx.Timestamp)

			txData := &models.TXHistory{
				Amount:   tx.Amount,
				Receiver: tx.Receiver,
				Fee:      tx.Fee,
				Hash:     tx.Hash,
				TS:       t.In(t.Local().Location()).Format("Jan _2 15:04:05"),
				Status:   "Complete",
				Failed:   false,
			}
			if a.NewUser {
				a.storeTX(txData)
			}
			a.RT.Events.Emit("new_transaction", txData)

			if i+1 == len(allTX) {

				err := a.rebuildTxChainState(tx.Hash)
				if err != nil {
					a.log.Errorln(err)
					// If unable to import previous transactions, remove wallet from DB and logout.
					//TODO: logout
					if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Delete(&a.wallet).Error; err != nil {
						a.log.Errorln("Unable to delete wallet upon failed import. Reason: ", err)
						return err
					}
					a.log.Panicln("Unable to import previous transactions")
					a.LoginError("Unable to collect previous TX's from blockexplorer. Please try again later.")
				}
			}
		}

	} else {
		a.log.Info("Unable to detect any previous transactions.")
		return nil
	}
	return nil

}

// PassKeysToFrontend emits the keys to the settings.Vue component on a
// 5 second interval
//TODO - @vito why is 5s interval necessary?
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
		err = errors.New("Received empty response from the Token Balance API")
		a.log.Warnln("Unable to update token balance. Reason: ", err)
		return 0, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
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
		err = errors.New("Unable to parse balance")
		a.log.Warnln("Unable to update token balance. Reason: ", err)
		return 0, err
	}
	a.log.Infoln("Parsed the following balance: ", s)

	a.log.Infoln("Returning the following balance: ", b)

	return b, nil
}
