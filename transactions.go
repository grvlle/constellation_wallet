package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Transaction contains all tx information
type Transaction struct {
	Edge struct {
		ObservationEdge struct {
			Parents []struct {
				Hash     string `json:"hash"`
				HashType string `json:"hashType"`
			} `json:"parents"`
			Data struct {
				Hash     string `json:"hash"`
				HashType string `json:"hashType"`
			} `json:"data"`
		} `json:"observationEdge"`
		SignedObservationEdge struct {
			SignatureBatch struct {
				Hash       string `json:"hash"`
				Signatures []struct {
					Signature string `json:"signature"`
					ID        struct {
						Hex string `json:"hex"`
					} `json:"id"`
				} `json:"signatures"`
			} `json:"signatureBatch"`
		} `json:"signedObservationEdge"`
		Data struct {
			Amount    int64 `json:"amount"`
			LastTxRef struct {
				Hash    string `json:"hash"`
				Ordinal int    `json:"ordinal"`
			} `json:"lastTxRef"`
			Fee  int   `json:"fee"`
			Salt int64 `json:"salt"`
		} `json:"data"`
	} `json:"edge"`
	LastTxRef struct {
		Hash    string `json:"hash"`
		Ordinal int    `json:"ordinal"`
	} `json:"lastTxRef"`
	IsDummy bool `json:"isDummy"`
	IsTest  bool `json:"isTest"`
}

// PrepareTransaction is triggered from the frontend and will initialize a new tx
func (a *WalletApplication) PrepareTransaction(amount int64, fee int, address string) *Transaction {

	tx := &Transaction{}

	tx.Edge.Data.Amount = amount
	tx.Edge.ObservationEdge.Data.Hash = address
	tx.Edge.Data.Fee = fee

	a.sendTransaction(amount, fee, address)

	return tx
}

// sendTransaction is called from the front end code. This function will call the wallet.jar
// to put the actual transaction on chain. It'll then call updateLastTransaction in order
// to display transaction history to the user.
func (a *WalletApplication) sendTransaction(amount int64, fee int, address string) {
	a.putTXOnNetwork(amount, fee, address)
	a.storeTX()
	//a.updateLastTransactions()
}

func (a *WalletApplication) initTXs() {
	transactions := &a.wallet.TXHistory
	a.DB.Model(&a.wallet).Where("alias = ?", a.wallet.WalletAlias).Association("TXHistory").Find(&transactions)

	a.log.Warnln(a.wallet.TXHistory)
	for i := range a.wallet.TXHistory {
		a.RT.Events.Emit("new_transaction", &a.wallet.TXHistory[i]) // Pass the tx to the frontend as a new transaction.
	}

}

func (a *WalletApplication) storeTX() {
	txData := a.parseTX()
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Association("TXHistory").Append(txData).Error; err != nil {
		a.log.Errorf("Unable to update the DB record with the new TX. Reason: ", err)
		a.sendError("Unable to update the DB record with the new TX. Reason: ", err)
	}
	a.log.Infoln("Successfully stored tx in DB")
}

func (a *WalletApplication) parseTX() *TXHistory {
	txObject := a.loadTX()
	tx := &Transaction{}

	bytes := []byte(txObject)
	err := json.Unmarshal(bytes, &tx)
	if err != nil {
		a.sendError("Unable to parse the last transaction. Reason:", err)
		a.log.Warnf("Unable to parse contents of last_tx. Reason: %s", err)
	}
	txData := &TXHistory{
		Amount:          tx.Edge.Data.Amount,
		Address:         tx.Edge.ObservationEdge.Parents[1].Hash,
		Fee:             tx.Edge.Data.Fee,
		TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
		TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
	}
	a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.

	return txData
}

// collectTXHistory is called by initTransactionHistory to read and parse the LastTXFile.
// It will scan the lines and append them to txObjects which is later returned to
// initTransactionHistory
func (a *WalletApplication) loadTX() string {
	var txObjects string
	file, err := os.Open(a.paths.LastTXFile) // acct
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txObjects = scanner.Text()
	}
	defer file.Close()
	return txObjects
}

// updateLastTransaction will read the contents of txhistory.json into memory and pass it onto
// the frontend to display transaction history of the user.
// func (a *WalletApplication) updateLastTransactions() {
// 	tx := &Transaction{}
// 	txObjects := a.collectTXHistory()

// 	lastTXindex := len(txObjects)

// 	bytes := []byte(txObjects[lastTXindex-1])
// 	err := json.Unmarshal(bytes, &tx)
// 	if err != nil {
// 		a.sendError("Unable to parse contents of acct. Reason:", err)
// 		a.log.Warnf("Unable to parse contents of acct. Reason: %s", err)
// 	}
// 	txData := &txInformation{
// 		ID:              tx.Edge.Count,
// 		Amount:          tx.Edge.Data.Amount,
// 		Address:         tx.Edge.ObservationEdge.Parents[1].Hash,
// 		Fee:             tx.Edge.Data.Fee,
// 		TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
// 		TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
// 	}

// 	a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.
// }

// // initTransactions history is called from wailsInit() in main.go. It will parse the tx file
// // generated by the wallet.jar file, append to the data, reverse the objects and write it to
// // txhistory.json
// func (a *WalletApplication) initTransactionHistory() {
// 	tx := &Transaction{}

// 	txObjects := a.collectTXHistory() // Parses the LastTXFile and returns the json objects
// 	var txObjectsPopulated []*txInformation
// 	var txObjectsReversed []*txInformation

// 	for _, txObject := range txObjects {
// 		bytes := []byte(txObject)
// 		err := json.Unmarshal(bytes, &tx)
// 		if err != nil {
// 			a.sendError("Unable to parse contents of acct. Reason:", err)
// 			a.log.Warnf("Unable to parse contents of acct. Reason: %s", err)
// 		}
// 		txData := &txInformation{
// 			ID:              tx.Edge.Count,
// 			Amount:          tx.Edge.Data.Amount,
// 			Address:         tx.Edge.ObservationEdge.Parents[1].Hash,
// 			Fee:             tx.Edge.Data.Fee,
// 			TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
// 			TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
// 		}
// 		txObjectsPopulated = append(txObjectsPopulated, txData)
// 	}

// 	txObjectsReversed = reverseElement(txObjectsPopulated) // We want to display the last tx at the top

// 	// Need to emit twice for it to stick. Bug with the Wails lib.
// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			a.RT.Events.Emit("update_tx_history", txObjectsReversed)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	// TODO: Add txhistory to database
// 	err := writeToJSON("txhistory", txObjectsReversed)
// 	if err != nil {
// 		a.sendError("Unable to write txhistory to fs. Reason:", err)
// 		a.log.Errorf("Unable to read txhistory to fs. Reason: %s", err)
// 	}

// }

// // collectTXHistory is called by initTransactionHistory to read and parse the LastTXFile.
// // It will scan the lines and append them to txObjects which is later returned to
// // initTransactionHistory
// func (a *WalletApplication) collectTXHistory() []string {
// 	var txObjects []string
// 	file, err := os.Open(a.paths.LastTXFile) // acct
// 	if err != nil {
// 		a.sendError("Unable to read tx data. Reason:", err)
// 		a.log.Errorf("Unable to read tx data. Reason: %s", err)
// 	}

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)

// 	for scanner.Scan() {
// 		txObjects = append(txObjects, scanner.Text())
// 	}
// 	defer file.Close()
// 	return txObjects
// }
