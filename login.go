package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

/* Database Model is located in models.go */

// CreateUser is called when creating a new wallet in frontend component Login.vue
func (a *WalletApplication) CreateWallet(keystorePath, keystorePassword, keyPassword string) bool {

	fmt.Println(keystorePath)
	if keystorePath == "" {
		keystorePath = a.wallet.KeyStorePath
	}

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

	if err := a.DB.FirstOrCreate(&Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed}, &a.wallet, 1).Error; err != nil {
		a.log.Errorf("Unable to create database object for new wallet. Reason: ", err)
		a.sendError("Unable to create database object for new wallet. Reason: ", err)
		return false
	}

	if err := a.DB.First(&a.wallet, 1).Updates(&Wallet{KeyStorePath: keystorePath, KeystorePasswordHash: keystorePasswordHashed, KeyPasswordHash: keyPasswordHashed}).Error; err != nil {
		a.log.Errorf("Unable to query database object for new wallet. Reason: ", err)
		a.sendError("Unable to query database object for new wallet. Reason: ", err)
		return false
	}

	err = a.initNewWallet()
	if err != nil {
		a.log.Errorf("Unable to initialize wallet object. Reason: ", err)
		a.sendError("Unable to initialize wallet object. Reason: ", err)
	}

	a.UserLoggedIn = false
	a.NewUser = true

	return true
}

func (a *WalletApplication) Login(keystorePath, keystorePassword, keyPassword string) bool {

	if err := a.DB.First(&a.wallet, 1).Error; err != nil {
		a.log.Errorf("Unable to query database object for new wallet. Reason: ", err)
		a.sendError("Unable to query database object for new wallet. Reason: ", err)
		return false
	}
	if keystorePath == "" {
		keystorePath = a.wallet.KeyStorePath
		a.log.Infoln("PrivateKey path: %s", keystorePath)
	}
	if a.CheckAccess(keystorePassword, a.wallet.KeystorePasswordHash) && a.CheckAccess(keyPassword, a.wallet.KeyPasswordHash) {
		a.UserLoggedIn = true
		os.Setenv("CL_STOREPASS", keystorePassword)
		os.Setenv("CL_KEYPASS", keyPassword)
	} else {
		a.UserLoggedIn = false
	}
	if a.UserLoggedIn && !a.NewUser {
		a.initExistingWallet(keystorePath)
	}

	a.NewUser = false

	return a.UserLoggedIn
}

func (a *WalletApplication) LogOut() {
	a.UserLoggedIn = true
}

func (a *WalletApplication) ImportKey() string {
	a.paths.EncPrivKeyFile = a.RT.Dialog.SelectFile()
	if a.paths.EncPrivKeyFile == "" {
		a.sendError("Please enter a valid path", nil)
	}

	a.log.Info("Path to imported key: " + a.paths.EncPrivKeyFile)
	return a.paths.EncPrivKeyFile
}

func (a *WalletApplication) SelectDirToStoreKey() string {
	a.paths.EncryptedDir = a.RT.Dialog.SelectDirectory()
	a.paths.EncPrivKeyFile = a.paths.EncryptedDir + "/key.p12"
	return a.paths.EncPrivKeyFile
}

func (a *WalletApplication) GenerateSaltedHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes[:])
	return hash, nil
}

func (a *WalletApplication) CheckAccess(password, passwordHash string) bool {
	err := a.Compare(password, passwordHash)
	if err != nil {
		a.log.Warnln("User tried to login with the wrong credentials!")
		return false
	} else {
		a.log.Infoln("Password check OK")
	}
	return true
}

func (a *WalletApplication) Compare(s, hash string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

// func (a *WalletApplication) GenerateUUID() string {
// 	n := 5
// 	b := make([]byte, n)
// 	if _, err := rand.Read(b); err != nil {
// 		panic(err)
// 	}
// 	s := fmt.Sprintf("%X", b)
// 	return s
// }
