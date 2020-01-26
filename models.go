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
	TXHistory            []TXHistory `sql:"-"`
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
	FromAddress string
	ToAddress   string
	TXHash      string
	Amount      uint64
}

type Address string

func (a *WalletApplication) CreateNewDBRecord(record interface{}) error {
	if !a.DB.NewRecord(record) {
		a.log.Warn("The value's primary key is not blank")
	}
	if err := a.DB.Create(record).Error; err != nil {
		a.log.Warn("Unable to create new Database record")
		return err
	}
	a.log.Info("A new Database Record were successfully added.")
	return nil
}
