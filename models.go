package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Wallet holds all wallet information.
type Wallet struct {
	ID                   uint `gorm:"AUTO_INCREMENT"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
	KeystorePasswordHash string
	KeyPasswordHash      string
	KeyStorePath         string
	WalletAlias          string      `gorm:"primary_key;unique"`
	Addresses            []Address   `sql:"-"`
	TXHistory            []TXHistory `gorm:"foreignkey:Alias"`
	ProfilePicture       string
	WalletTag            string
	Balance              int    `json:"balance"`
	AvailableBalance     int    `json:"available_balance"`
	Nonce                int    `json:"nonce"`
	TotalBalance         int    `json:"total_balance"`
	Delegated            int    `json:"delegated"`
	Deposit              int    `json:"deposit"`
	Address              string `json:"address"`
	TokenPrice           struct {
		DAG struct {
			BTC float64 `json:"BTC,omitempty"`
			USD float64 `json:"USD,omitempty"`
			EUR float64 `json:"EUR,omitempty"`
		} `json:"DAG"`
	} `json:"token_price"`
}

type TXHistory struct {
	gorm.Model
	Alias           string `json:"alias"`
	Amount          int64  `json:"amount"`
	Address         string `json:"address"`
	Fee             int    `json:"fee"`
	TransactionHash string `json:"txhash,omitempty"`
	TS              string `json:"date,omitempty"`
}

type Address string
