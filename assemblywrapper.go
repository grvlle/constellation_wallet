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

func (a *WalletApplication) runCMD(scalaFunc string, scalaArgs ...string) error {

	main := "java"
	cmds := []string{"-cp", "bcprov-jdk15on-1.62.jar:constellation-assembly-1.0.12.jar", scalaFunc}
	args := append(cmds, scalaArgs...)
	cmd := exec.Command(main, args...)

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
	err := a.runCMD("org.constellation.util.wallet.GenerateAddress", "--pub_key_str="+a.Wallet.PublicKey.Key, "--store_path="+a.paths.DAGDir)
	if err != nil {
		a.sendError("Unable to generate wallet address. Reason:", err)
		a.log.Errorf("Unable to generate wallet address. Reason: %s", err)
	}

	bytes, err := a.getFileContents(a.paths.AddressFile) // addr
	if err != nil {
		a.sendError("Unable to read DAG Address from filesystem. Reason:", err)
		a.log.Errorf("Unable to read DAG Address from filesystem. Reason: %s", err)
	}

	return string(bytes)
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
	err := a.runCMD("org.constellation.util.wallet.CreateNewTransaction", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass, "--account_path="+a.paths.DAGDir, "--amount="+amountStr, "--fee="+feeStr, "--destination="+address, "--store_path="+a.paths.LastTXFile, "--priv_key_str="+a.Wallet.PrivateKey.Key, "--pub_key_str="+a.Wallet.PublicKey.Key)
	if err != nil {
		a.sendError("Unable to send transaction. Reason:", err)
		a.log.Errorf("Unable to send transaction. Reason: %s", err)
	}
	time.Sleep(10) // Will sleep for 10 sec between TXs to prevent spamming.
}

// createEncryptedKeyPairToPasswordProtectedFile is called ONLY when a NEW wallet is created. This
// will create a new password protected encrypted keypair stored in $HOME/.dag/encrypted_key/priv.p12
// (a.paths.EncPrivKey)
func (a *WalletApplication) createEncryptedKeyPairToPasswordProtectedFile(alias, storepass, keypass string) {

	err := a.runCMD("org.constellation.keytool.KeyTool", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass)
	if err != nil {
		a.sendError("Unable to write encrypted keys to filesystem. Reason:", err)
		a.log.Fatalf("Unable to write encrypted keys to filesystem. Reason: %s", err)
	}

}

// decryptKeyPair dumps the encrypted KeyStore in unencrypted .pem format
func (a *WalletApplication) decryptKeyPair(alias, storepass, keypass string) {

	err := a.runCMD("org.constellation.util.wallet.ExportDecryptedKeysTest", "--keystore="+a.paths.EncPrivKeyFile, "--alias="+alias, "--storepass="+storepass, "--keypass="+keypass, "--keypass="+keypass, "--priv_store_path="+"decrypted_keystore", "--pub_store_path="+"decrypted_keystore.pub")
	if err != nil {
		a.sendError("Unable to write encrypted keys to filesystem. Reason:", err)
		a.log.Fatalf("Unable to write encrypted keys to filesystem. Reason: %s", err)
	}

}
