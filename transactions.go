package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Transaction contains all tx information
type Transaction struct {
	Edge struct {
		ObservationEdge struct {
			Parents []struct {
				Hash     string `json:"hash,omitempty"`
				HashType string `json:"hashType,omitempty"`
			} `json:"parents,omitempty"`
			Data struct {
				Hash     string `json:"hash,omitempty"`
				HashType string `json:"hashType,omitempty"`
			} `json:"data,omitempty"`
		} `json:"observationEdge,omitempty"`
		SignedObservationEdge struct {
			SignatureBatch struct {
				Hash       string `json:"hash,omitempty"`
				Signatures []struct {
					Signature string `json:"signature,omitempty"`
					ID        struct {
						Hex string `json:"hex,omitempty"`
					} `json:"id,omitempty"`
				} `json:"signatures,omitempty"`
			}
		} `json:"signedObservationEdge,omitempty"`
		Data struct {
			Amount int64 `json:"amount,omitempty"`
			Salt   int64 `json:"salt,omitempty"`
			Fee    int   `json:"fee,omitempty"`
		} `json:"data,omitempty"`
		PreviousHash string `json:"previousHash,omitempty"`
		Count        int    `json:"count,omitempty"`
		IsDummy      bool   `json:"isDummy,omitempty"`
		Signature    []int  `json:"signature,omitempty"`
	} `json:"edge,omitempty"`
}

func sendTransaction(amount int64, address string) *Transaction {
	fee := 10
	tx := &Transaction{}

	tx.Edge.Data.Amount = amount
	tx.Edge.ObservationEdge.Data.Hash = address
	tx.Edge.Data.Fee = fee

	amountStr := strconv.FormatInt(amount, 10)
	feeStr := strconv.FormatInt(amount, 10)

	// newTX is the full command to sign a new transaction
	newTX := "java -cp constellation-assembly-1.0.12.jar org.constellation.SignNewTx " + amountStr + ` "` + address + `" ` + feeStr
	parts := strings.Fields(newTX)
	head := parts[0]
	parts = parts[1:len(parts)]

	os.Setenv("PATH", "/usr/bin:/sbin") // This is neccessary when interacting with the CLI on RedHat and other linux distros
	cmd := exec.Command(head, parts...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out    // Captures STDOUT
	cmd.Stderr = &stderr // Captures STDERR

	err := cmd.Run()
	if err != nil {
		err := fmt.Sprint(err) + ": " + stderr.String()
		fmt.Println("Unable to send transaction. Reason:", err)
		return tx
	}
	fmt.Println(out.String())

	return tx
}

// sendTransaction will create the tx object and populate it with data
// collected in the forms in Transactions.vue (amountSubmitted and txAddress).
// func sendTransaction(amount int, address string) *Transaction {

// 	tx := &Transaction{
// 		Amount:  amount,
// 		Address: address, // "0x161D1B0bca85e29dF546AFba1360eEc6Ab4aA7Ee",
// 		TS:      time.Now().Format("Mon Jan _2 15:04:05 2006"),
// 	}

// 	tx.ID++

// 	err := writeToJSON("tx.json", tx) // Temporary solution
// 	if err != nil {
// 		fmt.Println("Unable to write transaction data to tx.json.")
// 		tx.Status = false
// 	}
// 	tx.Status = true

// 	return tx
// }

// func (w *Wallet) RecieveTransaction() {
// 	return rx
// }
