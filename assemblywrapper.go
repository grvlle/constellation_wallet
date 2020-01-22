package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

/* AssemblyWrapper contains the code that is interacting with the wallet assembly provided
by the Constellation Engineering team. The code is interfacing with the wallet assembly using
the CLI */

func (a *WalletApplication) runWalletCMD(scalaFunc string, scalaArgs ...string) error {

	main := "java"
	cmds := []string{"-jar", "cl-wallet.jar", scalaFunc}
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
		return fmt.Errorf(errFormatted)
	}
	fmt.Println(out.String())
	a.log.Debugln(cmd)

	return nil
}

func (a *WalletApplication) runKeyToolCMD(scalaFunc string, scalaArgs ...string) error {

	main := "java"
	cmds := []string{"-jar", "cl-keytool.jar", scalaFunc}
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
		return fmt.Errorf(errFormatted)
	}
	fmt.Println(out.String())

	return nil
}

// createAddressFromPublicKey takes the pubKey hash and writes the DAG formatted
// address to a file on the filesystem.

// java -cp constellation-assembly-1.0.12.jar org.constellation.util.wallet.GenerateAddress --pub_key_str=<base64 hash of pubkey> --store_path=<path to file where address will be stored>
func (a *WalletApplication) createAddressFromPublicKey(alias string) string {
	a.log.Infoln("Creating DAG Address from Public Key...")

	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		a.log.Errorf("Unable to pipe STDOUT, Reason: ", err)
		a.sendError("Unable to pipe STDOUT, Reason: ", err)
	}
	os.Stdout = w

	fmt.Println("Hello, playground") // this gets captured

	err = a.runWalletCMD("show-address", "--keystore="+a.paths.EncPrivKeyFile, "--storepass=$CL_STOREPASS", "--alias="+alias, "--keypass=$CL_KEYPASS", "--env_args=true")
	if err != nil {
		a.sendError("Unable to generate wallet address. Reason:", err)
		a.log.Errorf("Unable to generate wallet address. Reason: %s", err.Error())
	}

	w.Close()
	dagAddress, err := ioutil.ReadAll(r)
	if err != nil {
		a.log.Errorf("Unable to read address from STDOUT", err)
		a.sendError("Unable to read address from STDOUT", err)
	}
	os.Stdout = rescueStdout

	// bytes, err := a.getFileContents(a.paths.AddressFile) // addr
	// if err != nil {
	// 	a.sendError("Unable to read DAG Address from filesystem. Reason: ", err)
	// 	a.log.Errorf("Unable to read DAG Address from filesystem. Reason: %s", err.Error())
	// }
	// formattedAddress := string((bytes[1 : len(bytes)-1]))

	return string(dagAddress)
}

// putTXOnNetwork will put an actual transaction on the network. This is called from the
// transactions.go file, more specifically the sendTransaction func. This in turn is triggered
// from the frontend (Transactions.vue) and the tx func. note you can either pass a priv key like
// or pass in a path to an encrypted .p12 file

// java -jar cl-wallet.jar create-transaction --keystore testkey.p12 --alias alias --storepass storepass --keypass keypass -d DAG6o9dcxo2QXCuJS8wnrR944YhFBpwc2jsh5j8f -p prev_tx -f new_tx --fee 0 --amount 1
func (a *WalletApplication) putTXOnNetwork(amount int64, fee int, address, alias string) {

	// Convert to string
	amountStr := strconv.FormatInt(amount, 10)
	feeStr := strconv.Itoa(fee)

	// newTX is the full command to sign a new transaction
	err := a.runWalletCMD("create-transaction", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass=$CL_STOREPASS", "--keypass=$CL_KEYPASS", "--account_path="+a.paths.DAGDir, "--amount="+amountStr, "--fee="+feeStr, "-d="+address, "-f="+a.paths.LastTXFile, "-p="+a.paths.PrevTXFile, "--env_args=true")
	if err != nil {
		a.sendError("Unable to send transaction. Reason: ", err)
		a.log.Errorf("Unable to send transaction. Reason: %s", err.Error())
	}
	time.Sleep(10 * time.Second) // Will sleep for 10 sec between TXs to prevent spamming.
}

// createEncryptedKeyPairToPasswordProtectedFile is called ONLY when a NEW wallet is created. This
// will create a new password protected encrypted keypair stored in $HOME/.dag/encrypted_key/priv.p12
// (a.paths.EncPrivKey)

// java -jar cl-keytool.jar --keystore testkey.p12 --alias alias --storepass storepass --keypass keypass

func (a *WalletApplication) createEncryptedKeyPairToPasswordProtectedFile(alias string) {
	err := a.runKeyToolCMD("--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass=$CL_STOREPASS", "--keypass=$CL_KEYPASS", "--env_args=true")
	if err != nil {
		a.sendError("Unable to write encrypted keys to filesystem. Reason: ", err)
		a.log.Errorf("Unable to write encrypted keys to filesystem. Reason: %s", err.Error()) // TODO: change to fatal
	}

}

// decryptKeyPair dumps the encrypted KeyStore in unencrypted .pem format
// func (a *WalletApplication) decryptKeyPair(alias, storepass, keypass string) {

// 	err := a.runWalletCMD("org.constellation.util.wallet.ExportDecryptedKeys", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass, "--priv_store_path="+a.paths.DecKeyFile, "--pub_store_path="+a.paths.PubKeyFile)
// 	if err != nil {
// 		a.sendError("Unable to write decrypted keys to filesystem. Reason:", err)
// 		a.log.Fatalf("Unable to write decrypted keys to filesystem. Reason: %s", err)
// 	}

// }
