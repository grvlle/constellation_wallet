package main

import (
	"encoding/base64"
	"os"
	"time"
)

// NewWallet initates a new wallet object
func (a *WalletApplication) NewWallet() *Wallet {

	return a.Wallet
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

// getKeys will parse key files, base64 encode them and remove the decrypted files.
func (a *WalletApplication) getKeys() (string, string) {
	a.newKeys()
	PrivKey, err := a.getFileContents(a.paths.DecKeyFile)
	if err != nil {
		a.sendError("Unable to parse PrivKey. Reason: ", err)
		a.log.Warnf("Unable to parse PrivKey file. Reason: %s", err)
	}
	PubKey, err := a.getFileContents(a.paths.PubKeyFile)
	if err != nil {
		a.sendError("Unable to parse PubKey. Reason: ", err)
		a.log.Warnf("Unable to parse PubKey file. Reason: %s", err)
	} else {
		a.log.Info("Keys successfully created")
	}
	// TEMPORARY. DO NOT REMOVE
	// err = a.removeKeyArtifacts()
	// if err != nil {
	// 	a.sendError("Unable to remove Key artifacts. Reason: ", err)
	// 	a.log.Warnf("Unable to remove Key artifacts. Reason: %s", err)
	// }

	return base64.StdEncoding.EncodeToString(PrivKey), base64.StdEncoding.EncodeToString(PubKey)
}

func (a *WalletApplication) newKeys() {
	a.decryptKeyPair("alias", "storepass", "keypass")
}

func (a *WalletApplication) removeKeyArtifacts() error {
	err := os.Remove(a.paths.DecKeyFile)
	if err != nil {
		return err
	}
	err = os.RemoveAll(a.paths.EncryptedDir)
	if err != nil {
		return err
	}
	return nil
}

// PassKeysToFrontend emits the keys to the settings.Vue component on a
// 5 second interval
func (a *WalletApplication) passKeysToFrontend(privateKey, publicKey, walletAddress string) {
	go func() {
		for {
			a.RT.Events.Emit("wallet_keys", privateKey, publicKey, walletAddress)
			time.Sleep(5 * time.Second)
		}
	}()
	a.WidgetRunning.PassKeysToFrontend = true
}
