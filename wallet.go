package main

import (
	"bytes"
	"encoding/json"
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

// getKeys will parse the wallet file(key) and store in the wallet
// data type.
func (a *WalletApplication) getKeys() (string, string) {
	a.newKeyPair()

	path := filepath.Join(a.paths.KeyFile)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		a.log.Warnf("Unable to parse wallet file. Reason: %s", err)
	}

	err = json.Unmarshal(f, &a.Wallet)
	if err != nil {
		a.log.Warnf("Unable to parse contents of acct. Reason: %s", err)
	}

	a.log.Info(a.Wallet.PrivateKey.Key)
	a.log.Info(a.Wallet.PublicKey.Key)

	return a.Wallet.PrivateKey.Key, a.Wallet.PublicKey.Key

}

// newKeyPair is used to generate a new pub/priv key using ECDSA. This
// function is called when a NewWallet() is created.
func (a *WalletApplication) newKeyPair() {

	// newKeys will check if keys exist and create new ones if not
	newKeys := "java -cp constellation-assembly-1.0.12.jar org.constellation.GetOrCreateKeys"
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
