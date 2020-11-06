package models

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
	Contact              []Contact   `gorm:"foreignkey:Alias"`
	TXHistory            []TXHistory `gorm:"foreignkey:Alias"`
	ProfilePicture       string
	WalletTag            string
	DarkMode             bool
	Currency             string
	TermsOfService       bool
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
