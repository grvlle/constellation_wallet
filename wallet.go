package main

import (
	"time"
)

// NewWallet initates a new wallet object
func (a *WalletApplication) NewWallet() *Wallet {

	return &a.wallet
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
func (a *WalletApplication) passKeysToFrontend(walletAddress string) {
	go func() {
		for {
			a.RT.Events.Emit("wallet_keys", "Support for a Mnemonic Seed coming soon", a.paths.EncPrivKeyFile, walletAddress)
			time.Sleep(5 * time.Second)
		}
	}()
	a.WidgetRunning.PassKeysToFrontend = true
}
