package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

/* AssemblyWrapper contains the code that is interacting with the wallet assembly provided
by the Constellation Engineering team. The code is interfacing with the wallet assembly using
the CLI */

func (a *WalletApplication) runWalletCMD(scalaFunc string, scalaArgs ...string) error {

	main := "java"
	cmds := []string{"-cp", "cl-wallet.jar", scalaFunc}
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
func (a *WalletApplication) createAddressFromPublicKey() string {
	a.log.Infoln("Creating DAG Address from Public Key...")
	err := a.runWalletCMD("org.constellation.util.wallet.GenerateAddress", "--pub_key_str="+"MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEwL78JcMnzMHM+bHm2NjlcF6PghRAvZU//Nwn/6O9Ckh6QBApecq3ybAFaOWPRyADy6lIKRRvGw+YxL714+lO1Q==", "--store_path="+a.paths.AddressFile)
	if err != nil {
		a.sendError("Unable to generate wallet address. Reason:", err)
		a.log.Errorf("Unable to generate wallet address. Reason: %s", err.Error())
	}

	bytes, err := a.getFileContents(a.paths.AddressFile) // addr
	if err != nil {
		a.sendError("Unable to read DAG Address from filesystem. Reason: ", err)
		a.log.Errorf("Unable to read DAG Address from filesystem. Reason: %s", err.Error())
	}
	formattedAddress := string((bytes[1 : len(bytes)-1]))

	return formattedAddress
}

// putTXOnNetwork will put an actual transaction on the network. This is called from the
// transactions.go file, more specifically the sendTransaction func. This in turn is triggered
// from the frontend (Transactions.vue) and the tx func. note you can either pass a priv key like
// or pass in a path to an encrypted .p12 file

// java -cp constellation-assembly-1.0.12.jar org.constellation.util.wallet.CreateNewTransaction --keystore=<path to kp.p12> --alias=alias --storepass=storepass --keypass=keypass --account_path=<tx-file.txt> --amount=137.035999084 --fee=0.007297 --destination=receiverAddress --store_path=src/test/resources/new-tx.txt --priv_key_str=<base64 hash of privkey> --pub_key_str=<base64 hash of pub>
func (a *WalletApplication) putTXOnNetwork(amount int64, fee int, address, alias, storepass, keypass string) {

	// Convert to string
	amountStr := strconv.FormatInt(amount, 10)
	feeStr := strconv.Itoa(fee)

	// newTX is the full command to sign a new transaction
	err := a.runWalletCMD("org.constellation.util.wallet.CreateNewTransaction", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass, "--account_path="+a.paths.DAGDir, "--amount="+amountStr, "--fee="+feeStr, "--destination="+address, "--store_path="+a.paths.LastTXFile, "--priv_key_str="+a.Wallet.PrivateKey, "--pub_key_str="+"MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEwL78JcMnzMHM+bHm2NjlcF6PghRAvZU//Nwn/6O9Ckh6QBApecq3ybAFaOWPRyADy6lIKRRvGw+YxL714+lO1Q==")
	if err != nil {
		a.sendError("Unable to send transaction. Reason: ", err)
		a.log.Errorf("Unable to send transaction. Reason: %s", err.Error())
	}
	time.Sleep(10 * time.Second) // Will sleep for 10 sec between TXs to prevent spamming.
}

// createEncryptedKeyPairToPasswordProtectedFile is called ONLY when a NEW wallet is created. This
// will create a new password protected encrypted keypair stored in $HOME/.dag/encrypted_key/priv.p12
// (a.paths.EncPrivKey)
func (a *WalletApplication) createEncryptedKeyPairToPasswordProtectedFile(alias, storepass, keypass string) {
	err := a.runKeyToolCMD("--keystore="+a.paths.EncryptedDir+"/key", "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass)
	if err != nil {
		a.sendError("Unable to write encrypted keys to filesystem. Reason: ", err)
		a.log.Errorf("Unable to write encrypted keys to filesystem. Reason: %s", err.Error()) // TODO: change to fatal
	}

}

// decryptKeyPair dumps the encrypted KeyStore in unencrypted .pem format
func (a *WalletApplication) decryptKeyPair(alias, storepass, keypass string) {

	err := a.runWalletCMD("org.constellation.util.wallet.ExportDecryptedKeys", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass, "--priv_store_path="+a.paths.DecKeyFile, "--pub_store_path="+a.paths.PubKeyFile)
	if err != nil {
		a.sendError("Unable to write decrypted keys to filesystem. Reason:", err)
		a.log.Fatalf("Unable to write decrypted keys to filesystem. Reason: %s", err)
	}

}
