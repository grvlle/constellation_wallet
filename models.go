package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Wallet holds all wallet information.
type Wallet struct {
	gorm.Model
	Username         string
	PasswordHash     string
	Addresses        []Address   `sql:"-"`
	TXHistory        []TXHistory `sql:"-"`
	ProfilePicture   string
	WalletTag        string
	Balance          int    `json:"balance"`
	AvailableBalance int    `json:"available_balance"`
	Nonce            int    `json:"nonce"`
	TotalBalance     int    `json:"total_balance"`
	Delegated        int    `json:"delegated"`
	Deposit          int    `json:"deposit"`
	Address          string `json:"address"`
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

type TXHistory struct {
	gorm.Model
	FromAddress string
	ToAddress   string
	TXHash      string
	Amount      uint64
}

type Address string

func (a *WalletApplication) initDB() {
	a.DB.Model(&Wallet{}).Related(&TXHistory{})
}
