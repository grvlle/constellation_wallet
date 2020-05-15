package app

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

/* AssemblyWrapper contains the code that is interacting with the wallet assembly provided
by the Constellation Engineering team. The code is interfacing with the wallet assembly using
the CLI */

func (a *WalletApplication) runWalletCMD(tool string, scalaFunc string, scalaArgs ...string) error {
	var main string

	if runtime.GOOS == "windows" {
		main = a.paths.Java
	} else {
		main = "java"
	}

	cmds := []string{"-jar", a.paths.DAGDir + "/cl-" + tool + ".jar", scalaFunc}
	args := append(cmds, scalaArgs...)
	cmd := exec.Command(main, args...)
	a.log.Infoln("Running command: ", cmd)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out    // Captures STDOUT
	cmd.Stderr = &stderr // Captures STDERR

	err := cmd.Run()
	if err != nil {
		errFormatted := fmt.Sprint(err) + ": " + stderr.String()
		return errors.New(errFormatted)
	}
	fmt.Println(out.String())
	a.log.Debugln(cmd)

	return nil
}

// WalletKeystoreAccess is true if the user can unlock the .p12 keystore
// and key using storepass and keypass
func (a *WalletApplication) WalletKeystoreAccess() bool {
	a.log.Infoln("Checking Keystore Access...")

	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		a.log.Errorln("Unable to pipe STDOUT, Reason: ", err)
		a.sendError("Unable to pipe STDOUT, Reason: ", err)
	}
	os.Stdout = w
	err = a.runWalletCMD("wallet", "show-address", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+a.wallet.WalletAlias, "--env_args=true")
	if err != nil {
		a.log.Warn("KeyStore Access Rejected!")
		a.LoginError("Access Denied. Please make sure that you have typed in the correct credentials.")
		a.KeyStoreAccess = false
		return a.KeyStoreAccess
	}

	// STDOUT is captured here

	w.Close()
	dagAddress, err := ioutil.ReadAll(r)
	if err != nil {
		a.log.Errorln("Unable to read address from STDOUT", err)
		a.sendError("Unable to read address from STDOUT", err)
	}
	// if STDOUT prefix of show-address output isn't DAG

	if err == nil && a.wallet.Address != "" && string(dagAddress[:40]) == a.wallet.Address {
		a.KeyStoreAccess = true
		a.log.Info("KeyStore Access Granted!")
		return a.KeyStoreAccess
	}
	os.Stdout = rescueStdout

	a.KeyStoreAccess = false
	a.log.Warn("KeyStore Access Rejected!")
	a.LoginError("Access Denied. Please make sure that you have typed in the correct credentials.")
	return a.KeyStoreAccess
}

// GenerateDAGAddress generates a new wallet address and stores it in memory
// java -cp constellation-assembly-1.0.12.jar org.constellation.util.wallet.GenerateAddress --pub_key_str=<base64 hash of pubkey> --store_path=<path to file where address will be stored>
func (a *WalletApplication) GenerateDAGAddress() string {
	a.log.Infoln("Creating DAG Address from Public Key...")

	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		a.log.Errorln("Unable to pipe STDOUT, Reason: ", err)
		a.sendError("Unable to pipe STDOUT, Reason: ", err)
		return ""
	}
	os.Stdout = w

	err = a.runWalletCMD("wallet", "show-address", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+a.wallet.WalletAlias, "--env_args=true")
	if err != nil {
		a.sendError("Unable to generate wallet address. Reason:", err)
		a.log.Errorf("Unable to generate wallet address. Reason: %s", err.Error())
		return ""
	}

	w.Close()
	dagAddress, err := ioutil.ReadAll(r)
	if err != nil {
		a.log.Errorln("Unable to read address from STDOUT", err)
		a.sendError("Unable to read address from STDOUT", err)
	}
	os.Stdout = rescueStdout
	a.wallet.Address = string(dagAddress[:40])

	return a.wallet.Address
}

// CheckAndFetchWalletCLI will download the cl-wallet dependencies from
// the official Constellation Repo
func (a *WalletApplication) CheckAndFetchWalletCLI() {
	keytoolPath := a.paths.DAGDir + "/cl-keytool.jar"
	walletPath := a.paths.DAGDir + "/cl-wallet.jar"

	keytoolExists := a.fileExists(keytoolPath)
	walletExists := a.fileExists(walletPath)

	if keytoolExists && walletExists {
		a.RT.Events.Emit("downloading_dependencies", false)
	} else {
		a.RT.Events.Emit("downloading_dependencies", true)
	}

	if keytoolExists {
		a.log.Info(keytoolPath + " file exists. Skipping downloading")
	} else {
		if err := a.fetchWalletJar("cl-keytool.jar", keytoolPath); err != nil {
			a.log.Errorln("Unable to fetch or store cl-keytool.jar", err)
		}
	}

	if walletExists {
		a.log.Info(walletPath + " file exists. Skipping downloading")
	} else {
		if err := a.fetchWalletJar("cl-wallet.jar", walletPath); err != nil {
			a.log.Errorln("Unable to fetch or store cl-wallet.jar", err)
		}
	}

	if a.fileExists(keytoolPath) && a.fileExists(walletPath) {
		a.RT.Events.Emit("downloading_dependencies", false)
	}

}

// produceTXObject will put an actual transaction on the network. This is called from the
// transactions.go file, more specifically the sendTransaction func. This in turn is triggered
// from the frontend (Transactions.vue) and the tx func. note you can either pass a priv key like
// or pass in a path to an encrypted .p12 file

// java -jar cl-wallet.jar create-transaction --keystore testkey.p12 --alias alias --storepass storepass --keypass keypass -d DAG6o9dcxo2QXCuJS8wnrR944YhFBpwc2jsh5j8f -p prev_tx -f new_tx --fee 0 --amount 1
func (a *WalletApplication) produceTXObject(amount int64, fee int64, address, newTX, prevTX string) {

	// Convert to string
	amountStr := strconv.FormatInt(amount, 10)
	// feeStr := strconv.FormatInt(fee, 10)

	// amountNorm, err := normalizeAmounts(amount)
	// if err != nil {
	// 	a.log.Errorln("Unable to normalize amounts when producing tx object. Reason: ", err)
	// 	a.sendError("Unable to send transaction. Don't worry, your funds are safe. Please report this issue. Reason: ", err)
	// 	return
	// }
	feeNorm, err := normalizeAmounts(fee)
	if err != nil {
		a.log.Errorln("Unable to normalize amounts when producing tx object. Reason: ", err)
		a.sendError("Unable to send transaction. Don't worry, your funds are safe. Please report this issue. Reason: ", err)
		return
	}

	// newTX is the full command to sign a new transaction
	err = a.runWalletCMD("wallet", "create-transaction", "--keystore="+a.paths.EncPrivKeyFile, "--normalized", "--alias="+a.wallet.WalletAlias, "--amount="+amountStr, "--fee="+feeNorm, "-d="+address, "-f="+newTX, "-p="+prevTX, "--env_args=true")
	if err != nil {
		a.sendError("Unable to send transaction. Don't worry, your funds are safe. Please report this issue. Reason: ", err)
		a.log.Errorln("Unable to send transaction. Reason: ", err)
		return
	}
	time.Sleep(1 * time.Second) // Will sleep for 1 sec between TXs to prevent spamming.
}

// CreateEncryptedKeyStore is called ONLY when a NEW wallet is created. This
// will create a new password protected encrypted keypair stored in user selected location
// java -jar cl-keytool.jar --keystore testkey.p12 --alias alias --storepass storepass --keypass keypass
func (a *WalletApplication) CreateEncryptedKeyStore() error {
	err := a.runWalletCMD("keytool", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+a.wallet.WalletAlias, "--env_args=true")
	if err != nil {
		a.LoginError("Unable to write encrypted keys to filesystem.")
		a.log.Errorf("Unable to write encrypted keys to filesystem. Reason: %s", err.Error())
		return err
	}
	return nil
}
