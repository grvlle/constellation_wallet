package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID      int    `json:"id,omitempty"`
	Amount  int    `json:"amount"`
	Address string `json:"address,omitempty"`
	TS      string `json:"date,omitempty"`
	Status  bool   `json:"status,omitempty"`
}

// sendTransaction will create the tx object and populate it with data
// collected in the forms in Transactions.vue (amountSubmitted and txAddress).
func sendTransaction(amount int, address string) *Transaction {

	tx := &Transaction{
		Amount:  amount,
		Address: address, // "0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee",
		TS:      time.Now().Format("Mon Jan _2 15:04:05 2006"),
	}

	tx.ID++

	err := writeToJSON("tx.json", tx) // Temporary solution
	if err != nil {
		fmt.Println("Unable to write transaction data to tx.json.")
		tx.Status = false
	}
	tx.Status = true

	return tx
}

// func (w *Wallet) RecieveTransaction() {
// 	return rx
// }
