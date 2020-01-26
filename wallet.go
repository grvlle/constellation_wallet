package main

import (
	"os"
	"time"
)

/* Database Model is located in models.go */

// CreateUser is called when creating a new wallet in frontend component Login.vue
func (a *WalletApplication) CreateWallet(keystorePath, keystorePassword, keyPassword, alias string) bool {

	// if keystorePath == "" {
	// 	keystorePath = a.wallet.KeyStorePath

	// }
	// if alias == "" {
	// 	alias = a.wallet.WalletAlias
	// }

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

	if err := a.DB.Where(&Wallet{WalletAlias: alias}).FirstOrCreate(&Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed, WalletAlias: alias}).Error; err != nil {
		a.log.Errorf("Unable to create database object for new wallet. Reason: ", err)
		a.sendError("Unable to create database object for new wallet. Reason: ", err)
		return false
	}

	if err := a.DB.Where("wallet_alias = ?", alias).First(&a.wallet).Updates(&Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed, WalletAlias: alias}).Error; err != nil {
		a.log.Errorf("Unable to query database object for new wallet. Reason: ", err)
		a.sendError("Unable to query database object for new wallet. Reason: ", err)
		return false
	}

	err = a.initNewWallet()
	if err != nil {
		a.log.Errorf("Unable to initialize wallet object. Reason: ", err)
		a.sendError("Unable to initialize wallet object. Reason: ", err)
	}

	a.KeyStoreAccess = a.WalletKeystoreAccess()
	a.UserLoggedIn = false
	a.NewUser = true

	return true
}

// initWallet initializes a new wallet. This is called from login.vue
// only when a new wallet is created.
func (a *WalletApplication) initNewWallet() error {

	a.CreateEncryptedKeyStore()
	a.wallet.Address = a.GenerateDAGAddress()
	a.wallet.KeyStorePath = a.paths.EncPrivKeyFile

	if err := a.DB.Model(&a.wallet).Update("Address", a.wallet.Address).Error; err != nil {
		a.log.Errorf("Unable to query database object for new wallet. Reason: ", err)
		a.sendError("Unable to query database object for new wallet. Reason: ", err)
	}

	//a.initTransactionHistory()
	a.passKeysToFrontend()

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}

	a.log.Infoln("A New wallet has been created successfully!")

	return nil
}

// initExistingWallet queries the database for the user wallet and pushes
// the information to the front end components.
func (a *WalletApplication) initExistingWallet(keystorePath string) {

	a.paths.EncPrivKeyFile = keystorePath

	if !a.WidgetRunning.DashboardWidgets {
		a.initDashboardWidgets()
	}
	if !a.WidgetRunning.PassKeysToFrontend {
		a.passKeysToFrontend()
	}

	a.log.Infoln("User has logged into the wallet")

}

func (a *WalletApplication) initDashboardWidgets() {
	// Initializes a struct containing all Chart Data on the dashboard
	chartData := a.ChartDataInit()

	// Below methods are continously updating the client side modules.
	a.nodeStats(chartData)
	a.txStats(chartData)
	a.networkStats(chartData)
	a.blockAmount()
	a.tokenAmount()
	a.pricePoller()

	a.WidgetRunning.DashboardWidgets = true
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
