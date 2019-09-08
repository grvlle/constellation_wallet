package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Transaction struct {
	ID      int    `json:"id,omitempty"`
	Amount  int    `json:"amount"`
	Address string `json:"address,omitempty"`
	TS      string `json:"date,omitempty"`
	Status  error  `json:"status,omitempty"`
}

// sendTransaction will create the tx object and populate it with data
// collected in the forms in Transactions.vue (amountSubmitted and txAddress).
func sendTransaction(amount int, address string) *Transaction {

	tx := &Transaction{
		Amount:  amount,
		Address: address, //"0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee",
		TS:      time.Now().Format("Mon Jan _2 15:04:05 2006"),
	}

	tx.ID++
	err := transactionConfirmed(tx)
	if err != nil {
		fmt.Println("Unable to write transaction data to tx.json.")
		tx.Status = err
	}

	return tx
}

// transactionConfirmed will create a json file and populate the
// transaction in json format. This is a temp function.
func transactionConfirmed(tx *Transaction) error {
	txJSON, err := json.Marshal(tx)
	f, err := os.Create("tx.json")
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(txJSON)

	return nil
}

// func (w *Wallet) RecieveTransaction() {
// 	return rx
// }
