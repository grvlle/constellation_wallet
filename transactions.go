package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func (a *WalletApplication) TriggerTXFromFE(amount float64, fee float64, address string) bool {
	a.PrepareTransaction(amount, fee, address)
	for !a.TransactionFinished {
		time.Sleep(1 * time.Second)
	}
	return a.TransactionFailed
}

// PrepareTransaction is triggered from the frontend (Transaction.vue) and will initialize a new tx.
// methods called are defined in buildchain.go
func (a *WalletApplication) PrepareTransaction(amount float64, fee float64, address string) {

	// TODO: Temp comments. Re-add once wallet goes live.
	if amount+fee > a.wallet.AvailableBalance {
		a.log.Warnln("Insufficient Balance")
		a.sendWarning("Insufficent Balance.")
		a.TransactionFailed = true
		return
	}

	if a.TransactionFinished {
		a.TransactionFinished = false

		// Asynchronously inform FE of TX state
		go func() {
			for !a.TransactionFinished {
				a.RT.Events.Emit("tx_in_transit", a.TransactionFinished)
				time.Sleep(1 * time.Second)
			}
			a.RT.Events.Emit("tx_in_transit", a.TransactionFinished)
		}()
		ptx := a.loadTXFromFile(a.paths.PrevTXFile)
		ltx := a.loadTXFromFile(a.paths.LastTXFile)

		ptxObj, ltxObj := a.convertToTXObject(ptx, ltx)

		a.formTXChain(amount, fee, address, ptxObj, ltxObj)
	}
}

func (a *WalletApplication) putTXOnNetwork(tx *Transaction) bool {
	a.log.Info("Attempting to communicate with mainnet on: " + a.Network.URL + a.Network.Handles.Transaction)
	/* TEMPORARILY COMMENTED OUT */
	bytesRepresentation, err := json.Marshal(tx)
	if err != nil {
		a.log.Errorln("Unable to parse JSON data for transaction", err)
		a.sendError("Unable to parse JSON data for transaction", err)
		return false
	}
	resp, err := http.Post(a.Network.URL+a.Network.Handles.Transaction, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		a.log.Errorln("Failed to send HTTP request. Reason: ", err)
		a.sendError("Unable to send request to mainnet. Please check your internet connection. Reason: ", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			a.log.Fatal(err)
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
		a.log.Errorln(err)
	}
	bodyString := string(bodyBytes)
	a.sendError("Unable to communicate with mainnet. Reason: "+bodyString, err)
	a.log.Errorln("Unable to put TX on the network. HTTP Code: " + string(resp.StatusCode) + " - " + bodyString)

	time.Sleep(3 * time.Second)
	return false /* TEMPORARILY SET TO TRUE. CHANGE TO FALSE */
}

func (a *WalletApplication) sendTransaction(txFile string) *TXHistory {

	txObject := a.loadTXFromFile(txFile)

	tx := &Transaction{}

	bytes := []byte(txObject)
	err := json.Unmarshal(bytes, &tx)
	if err != nil {
		a.sendError("Unable to parse the last transaction. Reason:", err)
		a.log.Errorf("Unable to parse contents of last_tx. Reason: %s", err)
		return nil
	}

	// Put TX object on network
	if a.putTXOnNetwork(tx) {
		txData := &TXHistory{
			Amount: tx.Edge.Data.Amount,
			Sender: tx.Edge.ObservationEdge.Parents[1].Hash,
			Fee:    tx.Edge.Data.Fee,
			Hash:   tx.Edge.ObservationEdge.Data.Hash,
			TS:     time.Now().Format("Mon Jan _2 15:04:05 2006"),
			Failed: false,
		}
		a.storeTX(txData)
		a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.
		a.TransactionFinished = true
		a.TransactionFailed = false
		return txData
	}
	txData := &TXHistory{
		Amount: tx.Edge.Data.Amount,
		Sender: tx.Edge.ObservationEdge.Parents[1].Hash,
		Fee:    tx.Edge.Data.Fee,
		Hash:   tx.Edge.ObservationEdge.Data.Hash,
		TS:     time.Now().Format("Mon Jan _2 15:04:05 2006"),
		Failed: true,
	}
	a.log.Errorln("TX Failed, storing with failed state.")
	a.storeTX(txData)
	a.TransactionFinished = true
	a.TransactionFailed = true
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

// loadTXFromFile takes a file, scan it and returns it in an object
func (a *WalletApplication) loadTXFromFile(txFile string) string {
	var txObjects string

	fi, err := os.Stat(txFile)
	if err != nil {
		a.log.Errorln("Unable to stat last_tx. Reason: ", err)
		a.sendError("Unable to stat last_tx. Reason: ", err)
		return ""
	}
	// get the size
	size := fi.Size()
	if size <= 0 {
		a.log.Info("TX file is empty.")
		return ""
	}

	file, err := os.Open(txFile) // acct
	if err != nil {
		a.log.Errorln("Unable to open TX file. Reason: ", err)
		a.sendError("Unable to read last tx. Aborting... Reason: ", err)
		return ""
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txObjects = scanner.Text()
	}
	defer file.Close()
	return txObjects
}
