package main

import (
	"encoding/json"
	"os"
	"time"
)

type Transaction struct {
	ID      int       `json:"id,omitempty"`
	Amount  int       `json:"amount,omitempty"`
	Address string    `json:"address,omitempty"`
	TS      time.Time `json:"date,omitempty"`
}

func sendTransaction(amount int, address string) *Transaction {
	tx := &Transaction{
		Amount:  amount,
		Address: address, //"0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee",
		TS:      time.Now(),
	}
	tx.ID++
	transactionConfirmed(tx)
	return tx
}

func transactionConfirmed(tx *Transaction) {
	txJSON, err := json.Marshal(tx)

	f, err := os.OpenFile("tx.json", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	f.Write(txJSON)
	defer f.Close()
}

// func (w *Wallet) RecieveTransaction() {
// 	return rx
// }
