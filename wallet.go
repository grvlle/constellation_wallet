package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Wallet holds all wallet information.
type Wallet struct {
	Balance          int    `json:"balance"`
	AvailableBalance int    `json:"available_balance"`
	Nonce            int    `json:"nonce"`
	TotalBalance     int    `json:"total_balance"`
	Delegated        int    `json:"delegated"`
	Deposit          int    `json:"deposit"`
	Address          []byte `json:"address"`
	TokenPrice       struct {
		DAG struct {
			BTC float64 `json:"BTC,omitempty"`
			USD float64 `json:"USD,omitempty"`
			EUR float64 `json:"EUR,omitempty"`
		} `json:"DAG"`
	} `json:"token_price"`
	PrivateKey struct {
		Key string `json:"key"`
	} `json:"privateKey"`
	PublicKey struct {
		Key string `json:"key"`
	} `json:"publicKey"`
}

// NewWallet initates a new wallet object
func (a *WalletApplication) NewWallet() *Wallet {

	a.Wallet = &Wallet{
		Balance:          1024155,
		AvailableBalance: 1012233,
		Nonce:            420,
		TotalBalance:     1012420,
		Delegated:        42,
		Deposit:          0,
		Address:          []byte{0x00},
	}
	a.Wallet.PrivateKey.Key, a.Wallet.PublicKey.Key = a.getKeys()

	return a.Wallet
}

// getKeys will parse key files, base64 encode them and remove the decrypted files.
func (a *WalletApplication) getKeys() (string, string) {

	a.newKeyPair()
	PrivKey, err := a.getFileContents(a.paths.KeyFile)
	if err != nil {
		a.sendError("Unable to parse PrivKey. Reason: ", err)
		a.log.Warnf("Unable to parse PrivKey file. Reason: %s", err)
	}
	PubKey, err := a.getFileContents(a.paths.PubKeyFile)
	if err != nil {
		a.sendError("Unable to parse PubKey. Reason: ", err)
		a.log.Warnf("Unable to parse PubKey file. Reason: %s", err)
	}
	a.removeKeyArtifacts()
	a.log.Info("Keys successfully created")

	return base64.StdEncoding.EncodeToString(PrivKey), base64.StdEncoding.EncodeToString(PubKey)
}

// newKeyPair is used to generate a new pub/priv key using ECDSA. This
// function is called when a NewWallet() is created.
func (a *WalletApplication) newKeyPair() {

	// newKeys will check if keys exist and create new ones if not
	newKeys := "java -cp bcprov-jdk15on-1.62.jar:constellation-assembly-1.0.12.jar org.constellation.GetOrCreateKeys fakepassword true"
	parts := strings.Fields(newKeys)
	head := parts[0]
	parts = parts[1:len(parts)]

	os.Setenv("PATH", "/usr/bin:/sbin") // This is neccessary when interacting with the CLI on RedHat and other linux distros
	cmd := exec.Command(head, parts...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out    // Captures STDOUT
	cmd.Stderr = &stderr // Captures STDERR

	err := cmd.Run()
	if err != nil {
		err := fmt.Sprint(err) + ": " + stderr.String()
		a.log.Errorf("Unable to create/locate wallet. Reason:", err)
	}
	a.log.Info(out.String())
}

func (a *WalletApplication) removeKeyArtifacts() error {
	err := os.Remove(a.paths.KeyFile)
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
func (a *WalletApplication) passKeysToFrontend() {
	go func() {
		for {
			a.RT.Events.Emit("wallet_keys", a.Wallet.PrivateKey.Key, a.Wallet.PublicKey.Key)
			time.Sleep(5 * time.Second)
		}
	}()
}

func (a *WalletApplication) getFileContents(filePath string) ([]byte, error) {
	path := filepath.Join(filePath)
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return fileContents, nil
}
