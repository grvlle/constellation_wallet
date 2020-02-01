package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Wallet holds all wallet information.
type Wallet struct {
	ID                   uint `gorm:"AUTO_INCREMENT"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
	Path                 Path `gorm:"foreignkey:Alias"`
	KeystorePasswordHash string
	KeyPasswordHash      string
	KeyStorePath         string
	WalletAlias          string      `gorm:"primary_key;unique"`
	Addresses            []Address   `sql:"-"`
	TXHistory            []TXHistory `gorm:"foreignkey:Alias"`
	ProfilePicture       string
	WalletTag            string
	Balance              float64 `json:"balance"`
	AvailableBalance     float64 `json:"available_balance"`
	Nonce                float64 `json:"nonce"`
	TotalBalance         float64 `json:"total_balance"`
	Delegated            float64 `json:"delegated"`
	Deposit              float64 `json:"deposit"`
	Address              string  `json:"address"`
	TokenPrice           struct {
		DAG struct {
			BTC float64 `json:"BTC,omitempty"`
			USD float64 `json:"USD,omitempty"`
			EUR float64 `json:"EUR,omitempty"`
		} `json:"DAG"`
	} `json:"token_price"`
}

type Path struct {
	ID          uint   `json:"id"`
	Alias       string `json:"alias"`
	LastTXFile  string
	PrevTXFile  string
	EmptyTXFile string
}

type TXHistory struct {
	ID              uint    `json:"id"`
	Alias           string  `json:"alias"`
	Amount          float64 `json:"amount"`
	Address         string  `json:"address"`
	Fee             float64 `json:"fee"`
	TransactionHash string  `json:"txhash"`
	TS              string  `json:"date"`
	Failed          bool
}

type Address string
