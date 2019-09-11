package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version        = byte(0x00) // Number zero
)

// Wallet holds all wallet information.
type Wallet struct {
	Balance          int    `json:"balance"`
	AvailableBalance int    `json:"available_balance"`
	Nonce            int    `json:"nonce"`
	TotalBalance     int    `json:"total_balance"`
	Delegated        int    `json:"delegated"`
	Deposit          int    `json:"deposit"`
	Address          []byte `json:"address"`
	TokenPrice       struct {
		DAG struct {
			BTC float64 `json:"BTC,omitempty"`
			USD float64 `json:"USD,omitempty"`
			EUR float64 `json:"EUR,omitempty"`
		} `json:"DAG"`
	} `json:"token_price"`
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet initates a new wallet object
func (a *App) NewWallet() *Wallet {
	private, public := NewKeyPair()
	a.wallet = &Wallet{
		Balance:          1024155,
		AvailableBalance: 1012233,
		Nonce:            420,
		TotalBalance:     1012420,
		Delegated:        42,
		Deposit:          0,
		Address:          []byte{0x00},
		PrivateKey:       private,
		PublicKey:        public,
	}
	return a.wallet
}

// GetAddress returns the address in human readable.
func (w Wallet) GetAddress() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	w.Address = address

	err := writeToJSON("wallet.json", w) // Temporary solution
	if err != nil {
		fmt.Println("Unable to write transaction data to wallet.json.")
	}

	return address
}

// NewKeyPair is used to generate a new pub/priv key using ECDSA. This
// function is called when a NewWallet() is created.
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}

// PublicKeyHash produces a ripemd160 hash from the sha256 encrypted
// public key.
func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}
	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

// Checksum will perform a double hash on the payload and
// return a checksum
func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}
