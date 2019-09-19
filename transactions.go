package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

type txInformation struct {
	ID              int    `json:"id,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	Address         string `json:"address,omitempty"`
	Fee             int    `json:"fee,omitempty"`
	TransactionHash string `json:"txhash,omitempty"`
	TS              string `json:"date,omitempty"`
}

// PrepareTransaction is triggered from the frontend and will initialize a new tx
func (a *WalletApplication) PrepareTransaction(amount int64, address string) *Transaction {
	fee := 10
	tx := &Transaction{}

	tx.Edge.Data.Amount = amount
	tx.Edge.ObservationEdge.Data.Hash = address
	tx.Edge.Data.Fee = fee

	a.sendTransaction(amount, fee, address)

	return tx
}

func (a *WalletApplication) sendTransaction(amount int64, fee int, address string) {

	amountStr := strconv.FormatInt(amount, 10)
	feeStr := strconv.Itoa(fee)

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
		a.log.Errorf("Unable to send transaction. Reason:", err)
	}
	fmt.Println(out.String())
	a.updateLastTransactions()
}

func (a *WalletApplication) updateLastTransactions() {
	tx := &Transaction{}

	file, err := os.Open(a.paths.DAGDir + "/acct")
	if err != nil {
		a.log.Errorf("Unable to read tx data. Reason: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txObjects []string

	for scanner.Scan() {
		txObjects = append(txObjects, scanner.Text())
	}

	file.Close()

	for _, eachTX := range txObjects {
		fmt.Println(eachTX)
		bytes := []byte(eachTX)
		err = json.Unmarshal(bytes, &tx)
		if err != nil {
			a.log.Warnf("Unable to parse contents of acct. Reason: %s", err)
		}
		txData := &txInformation{
			ID:              tx.Edge.Count,
			Amount:          tx.Edge.Data.Amount,
			Address:         tx.Edge.ObservationEdge.Parents[0].Hash,
			Fee:             tx.Edge.Data.Fee,
			TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
			TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
		}

		a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.
	}
}
