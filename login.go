package main

import (
	"golang.org/x/crypto/bcrypt"
)

/* Database Model is located in models.go */

// CreateUser is called when creating a new wallet in frontend component Login.vue
func (a *WalletApplication) CreateUser(username, password string) bool {
	var wallet Wallet
	hashed, err := a.GenerateSaltedHash(password)
	if err != nil {
		a.log.Errorf("Unable to generate password hash. Reason: ", err)
		a.sendError("Unable to generate password hash. Reason: ", err)
		return false
	}
	if err := a.DB.Create(&Wallet{Username: username, PasswordHash: hashed}).Error; err != nil {
		a.log.Errorf("Unable to create database object for new wallet. Reason: ", err)
		a.sendError("Unable to create database object for new wallet. Reason: ", err)
		return false
	}

	if err := a.DB.First(&wallet, "username = ?", username).Error; err != nil {
		a.log.Errorf("Unable to query database object for new wallet. Reason: ", err)
		a.sendError("Unable to query database object for new wallet. Reason: ", err)
		return false
	}

	//a.NewWallet()
	a.UserLoggedIn = false

	return true
}

func (a *WalletApplication) Login(username, password string) bool {
	var wallet Wallet
	a.DB.First(&wallet, "username = ?", username) // find user in database

	if !a.UserLoggedIn {
		a.UserLoggedIn = a.CheckAccess(password, wallet.PasswordHash)
		err := a.initWallet()
		if err != nil {
			a.log.Errorf("Unable to initialize wallet object. Reason: ", err)
			a.sendError("Unable to initialize wallet object. Reason: ", err)
		}
	}
	return a.UserLoggedIn
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
		a.log.Infoln("User successfully logged in!")
	}
	return true
}

func (a *WalletApplication) Compare(s, hash string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
