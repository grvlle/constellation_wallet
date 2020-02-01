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
	Alias string `json:"alias"`
	Edge  struct {
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

func (a *WalletApplication) putTXOnNetwork(tx *Transaction) bool {
	a.log.Info("Attempting to communicate with mainnet on: http://" + a.Network.URL + a.Network.Handles.Transaction)
	bytesRepresentation, err := json.Marshal(tx)
	if err != nil {
		a.log.Errorln("Unable to parse JSON data for transaction", err)
		a.sendError("Unable to parse JSON data for transaction", err)
		return false
	}
	resp, err := http.Post("http://"+a.Network.URL+a.Network.Handles.Transaction, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		a.log.Errorf("Failed to send HTTP request. Reason: ", err)
		a.sendError("Unable to send request to mainnet. Please check your internet connection. Reason: ", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		if len(bodyBytes) == 64 {
			a.log.Info(bodyString)
			a.log.Infoln("Transaction has been successfully sent to the network.")
			a.sendSuccess("Transaction successfully sent!")
			return true
		}
		a.log.Warn(bodyString)
		a.sendWarning("Unable to put transaction on the network. Reason: " + bodyString)
		return false
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	a.sendError("Unable to communicate with mainnet. Reason: "+bodyString, err)
	a.log.Errorln("Unable to put TX on the network. HTTP Code: " + string(resp.StatusCode) + " - " + bodyString)
	return false
}

func (a *WalletApplication) sendTransaction(amount float64, fee float64, address string, txFile string) *TXHistory {

	txObject, err := a.loadTX(txFile)
	if err != nil {
		a.log.Errorln("Unable to read tx file.")
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
	if a.putTXOnNetwork(tx) {
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
	}

	a.log.Errorln("TX Failed, skipping database.")
	return nil

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
func (a *WalletApplication) loadTX(txFile string) (string, error) {
	var txObjects string

	fi, err := os.Stat(txFile)
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

	file, err := os.Open(txFile) // acct
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

// PrepareTransaction is triggered from the frontend (Transaction.vue) and will initialize a new tx.
// methods called are defined in buildchain.go
func (a *WalletApplication) PrepareTransaction(amount float64, fee float64, address string) {

	// TODO: Temp comments. Re-add once wallet goes live.
	// if amount+fee > a.wallet.AvailableBalance {
	// 	a.log.Warnln("Insufficient Balance")
	// 	a.sendWarning("Insufficent Balance.")
	// 	return nil
	// }

	ptx := loadTXFromFile(a.paths.PrevTXFile)
	ltx := loadTXFromFile(a.paths.LastTXFile)

	ptxObj, ltxObj := a.convertToTXObject(ptx, ltx)

	a.sendTX(amount, fee, address, ptxObj, ltxObj)
	time.Sleep(5 * time.Second)
}
