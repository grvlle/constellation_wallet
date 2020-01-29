package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
			Amount    float64 `json:"amount"`
			LastTxRef struct {
				Hash    string `json:"hash"`
				Ordinal int    `json:"ordinal"`
			} `json:"lastTxRef"`
			Fee  float64 `json:"fee"`
			Salt int64   `json:"salt"`
		} `json:"data"`
	} `json:"edge"`
	LastTxRef struct {
		Hash    string `json:"hash"`
		Ordinal int    `json:"ordinal"`
	} `json:"lastTxRef"`
	IsDummy bool `json:"isDummy"`
	IsTest  bool `json:"isTest"`
}

func (a *WalletApplication) networkHeartbeat() {
	//TODO
}

func (a *WalletApplication) putTXOnNetwork(tx *Transaction) {
	a.log.Info("Attempting to communicate with mainnet on: http://" + a.Network.URL + a.Network.Handles.Transaction)
	bytesRepresentation, err := json.Marshal(tx)
	if err != nil {
		a.log.Errorln("Unable to parse JSON data for transaction", err)
		a.sendError("Unable to parse JSON data for transaction", err)
	}
	resp, err := http.Post("http://"+a.Network.URL+a.Network.Handles.Transaction, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		a.log.Infoln("Transaction has been successfully sent to the network.")
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		a.log.Info(bodyString)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	a.log.Info(resp.StatusCode)
	a.log.Info(bodyString)
	return

}

// sendTransaction is called from the front end code. This function will call the wallet.jar
// to put the actual transaction on chain. It'll then call updateLastTransaction in order
// to display transaction history to the user.
func (a *WalletApplication) sendTransaction(amount float64, fee float64, address string) *TXHistory {
	txObject, err := a.loadTX()
	if err != nil {
		return nil
	}
	tx := &Transaction{}

	bytes := []byte(txObject)
	err = json.Unmarshal(bytes, &tx)
	if err != nil {
		a.sendError("Unable to parse the last transaction. Reason:", err)
		a.log.Errorf("Unable to parse contents of last_tx. Reason: %s", err)
		return nil
	}

	// Put TX object on network
	a.putTXOnNetwork(tx)

	txData := &TXHistory{
		Amount:          tx.Edge.Data.Amount,
		Address:         tx.Edge.ObservationEdge.Parents[1].Hash,
		Fee:             tx.Edge.Data.Fee,
		TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
		TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
	}
	a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.

	a.storeTX(txData)

	return txData

	//a.updateLastTransactions()
}

func (a *WalletApplication) storeTX(txData *TXHistory) {

	if txData == nil {
		return
	}
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Association("TXHistory").Append(txData).Error; err != nil {
		a.log.Errorf("Unable to update the DB record with the new TX. Reason: ", err)
		a.sendError("Unable to update the DB record with the new TX. Reason: ", err)
	}
	a.log.Infoln("Successfully stored tx in DB")
}

// loadTX will scan the lines and append them to txObjects which is later returned to
// initTransactionHistory
func (a *WalletApplication) loadTX() (string, error) {
	var txObjects string

	fi, err := os.Stat(a.paths.LastTXFile)
	if err != nil {
		a.log.Errorln("Unable to stat last_tx. Reason: ", err)
		a.sendError("Unable to stat last_tx. Reason: ", err)
		return "", err
	}
	// get the size
	size := fi.Size()
	if size <= 0 {
		a.log.Errorln("last_tx is empty. Reason: ", err)
		a.sendError("Unable to send transaction. Please report this issue. Your funds are safe. Reason: ", err)
		return "", err
	}

	file, err := os.Open(a.paths.LastTXFile) // acct
	if err != nil {
		a.log.Errorln("Unable to open last_tx. Reason: ", err)
		a.sendError("Unable to read last tx. Aborting... Reason: ", err)
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txObjects = scanner.Text()
	}
	defer file.Close()
	return txObjects, nil
}

// PrepareTransaction is triggered from the frontend and will initialize a new tx
func (a *WalletApplication) PrepareTransaction(amount float64, fee float64, address string) *Transaction {

	if amount+fee > a.wallet.AvailableBalance {
		a.log.Warnln("Insufficient Balance")
		a.sendWarning("Insufficent Balance.")
		return nil
	}

	tx := &Transaction{}

	tx.Edge.Data.Amount = amount
	tx.Edge.ObservationEdge.Data.Hash = address
	tx.Edge.Data.Fee = fee
	a.produceTXObject(amount, fee, address)
	a.sendTransaction(amount, fee, address)

	return tx
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
