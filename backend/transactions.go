package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/grvlle/constellation_wallet/backend/models"
)

// Transaction contains all tx information
type Transaction struct {
	Edge struct {
		ObservationEdge struct {
			Parents []struct {
				HashReference string `json:"hashReference"`
				HashType      string `json:"hashType"`
				BaseHash      string `json:"baseHash"`
			} `json:"parents"`
			Data struct {
				HashReference string `json:"hashReference"`
				HashType      string `json:"hashType"`
				BaseHash      string `json:"baseHash"`
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
				PrevHash string `json:"prevHash"`
				Ordinal  int    `json:"ordinal"`
			} `json:"lastTxRef"`
			Fee  int64 `json:"fee,omitempty"`
			Salt int64 `json:"salt"`
		} `json:"data"`
	} `json:"edge"`
	LastTxRef struct {
		PrevHash string `json:"prevHash"`
		Ordinal  int    `json:"ordinal"`
	} `json:"lastTxRef"`
	IsDummy bool `json:"isDummy"`
	IsTest  bool `json:"isTest"`
}

/* Send a transaction */

// TriggerTXFromFE will initate a new transaction triggered from the frontend.
func (a *WalletApplication) TriggerTXFromFE(amount float64, fee float64, address string) bool {
	amountConverted := int64(amount)
	feeConverted := int64(fee)

	a.PrepareTransaction(amountConverted, feeConverted, address)
	for !a.TransactionFinished {
		time.Sleep(1 * time.Second)
	}
	return a.TransactionFailed
}

// PrepareTransaction is triggered from the frontend (Transaction.vue) and will initialize a new tx.
// methods called are defined in buildchain.go
func (a *WalletApplication) PrepareTransaction(amount int64, fee int64, address string) {

	balance, err := a.GetTokenBalance()
	if err != nil {
		a.log.Errorln("Error when querying wallet balance. Reason: ", err)
		a.sendWarning("Unable to poll balance for wallet. Please try again later.")
		a.TransactionFailed = true
		return
	}

	if amount+fee > int64(balance) {
		a.log.Warnf("Trying to send: %d", amount+fee)
		a.log.Warnf("Insufficient Balance: %d", int64(balance))
		a.sendWarning("Insufficent Balance.")
		a.TransactionFailed = true
		return
	}

	if a.TransactionFinished {
		a.TransactionFinished = false

		// Asynchronously inform FE of TX state in wallet.
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

func (a *WalletApplication) putTXOnNetwork(tx *Transaction) (bool, string) {
	url := a.Network.URL + "/transactions"

	a.log.Info("Attempting to put transaction in network by sending to: " + url)

	bytesRepresentation, err := json.Marshal(tx)
	if err != nil {
		a.log.Errorln("Unable to parse JSON data for transaction", err)
		a.sendError("Unable to parse JSON data for transaction", err)
		return false, ""
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		a.log.Errorln("Failed to send HTTP request. Reason: ", err)
		a.sendError("Unable to send request to mainnet. Please check your internet connection. Reason: ", err)
		return false, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			a.log.Errorln(err)
		}
		bodyString := string(bodyBytes)
		a.sendError("Unable to communicate with network. Reason: "+bodyString, err)
		a.log.Errorln(fmt.Sprintf("Unable to put TX on the network. HTTP Code: %d - %s", resp.StatusCode, bodyString))

		return false, ""
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		a.log.Errorln("Failed to read the response body")
		a.log.Errorln("Failed to read the response body. Reason: ", err)
	}

	var result struct {
		Hash string `json:"hash"`
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil || len(result.Hash) == 0 {
		a.log.Errorln("Failed to read the response body")
		a.log.Errorln("Failed to read the response body. Reason: ", err)
	}

	hash := result.Hash

	a.log.Info("Transaction Hash: ", hash)
	a.TxPending(hash)
	a.log.Infoln("Transaction has been successfully sent to the network.")
	a.sendSuccess("Transaction successfully sent!")

	return true, hash
}

/* Note: Called from frontend to post a generated TX to the network */
func (a *WalletApplication) SendTransaction2(txJson string) bool {

	a.postTransaction(txJson)

	return !a.TransactionFailed
}

func (a *WalletApplication) sendTransaction(txFile string) *models.TXHistory {

	txObject := a.loadTXFromFile(txFile)

	return a.postTransaction(txObject)
}

func (a *WalletApplication) postTransaction(txObject string) *models.TXHistory {

	tx := &Transaction{}

	bytes := []byte(txObject)
	err := json.Unmarshal(bytes, &tx)
	if err != nil {
		a.sendError("Unable to parse the last transaction. Reason:", err)
		a.log.Errorf("Unable to parse contents of last_tx. Reason: %s", err)
		return nil
	}

	// Put TX object on network time.Now().Unix(), time.Now().Format("Jan _2 15:04:05")
	TXSuccessfullyPutOnNetwork, hash := a.putTXOnNetwork(tx)

	if TXSuccessfullyPutOnNetwork {
		txData := &models.TXHistory{
			Amount:      tx.Edge.Data.Amount,
			Source:      tx.Edge.ObservationEdge.Parents[0].HashReference,
			Destination: tx.Edge.ObservationEdge.Parents[1].HashReference,
			Fee:         tx.Edge.Data.Fee,
			Hash:        hash,
			Timestamp:   time.Now().Format(time.RFC3339),
			Status:      "Pending",
			Failed:      false,
		}
		a.storeTX(txData)
		a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.
		a.TransactionFinished = true
		a.TransactionFailed = false
		return txData
	}

	txData := &models.TXHistory{
		Amount:      tx.Edge.Data.Amount,
		Source:      tx.Edge.ObservationEdge.Parents[0].HashReference,
		Destination: tx.Edge.ObservationEdge.Parents[1].HashReference,
		Fee:         tx.Edge.Data.Fee,
		Hash:        hash,
		Timestamp:   time.Now().Format(time.RFC3339),
		Status:      "Error",
		Failed:      true,
	}

	a.log.Errorln("TX Failed, storing with failed state.")
	a.storeTX(txData)
	a.TransactionFinished = true
	a.TransactionFailed = true

	return txData
}

func (a *WalletApplication) storeTX(txData *models.TXHistory) {

	if txData == nil {
		return
	}
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Association("TXHistory").Append(txData).Error; err != nil {
		a.log.Errorln("Unable to update the DB record with the new TX. Reason: ", err)
		a.sendError("Unable to update the DB record with the new TX. Reason: ", err)
	}
	a.log.Infoln("Successfully stored tx in DB")
}

// loadTXFromFile takes a file, scans it and returns it in an object
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

/* Query TX */

// TxProcessed will query the last transaction. If no answer is returned, it means it's processed and the
// method will return true.
func (a *WalletApplication) TxProcessed(TXHash string) bool {
	todo := "TODO"

	a.log.Info("Communicating with mainnet on: " + todo)

	resp, err := http.Get(todo)
	if err != nil {
		a.log.Errorln("Failed to send HTTP request. Reason: ", err)
		if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Delete(&a.wallet).Error; err != nil {
			a.log.Errorln("Unable to delete wallet upon failed import. Reason: ", err)
			return false
		}
		a.log.Errorln("Unable to verify transaction status. Please check your internet connection.")
		return false
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		return false
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return false
	}

	if result["cbBaseHash"] != nil {
		a.log.Infoln("CheckPoint Hash :", result["cbBaseHash"])
		return true
	}

	// null response means it's snapshotted
	return string(bodyBytes) == "null"

}

type txStatus struct {
	Complete string
	Pending  string
	Error    string
}

// TxPending takes a TX Hash and updates the frontend with the current status (Pending/Error/Complete)
func (a *WalletApplication) TxPending(TXHash string) {

	status := &txStatus{
		Complete: "Complete",
		Pending:  "Pending",
		Error:    "Error",
	}

	consensus := 0

	select {
	case <-a.killSignal:
		return
	default:
		go func() bool {
			for retryCounter := 0; retryCounter < 30; retryCounter++ {
				processed := a.TxProcessed(TXHash)
				if !processed {
					a.log.Warnf("Transaction %v pending", TXHash)
					a.RT.Events.Emit("tx_pending", status.Pending)
					time.Sleep(time.Duration(retryCounter) * time.Second) // Increase polling interval

					if retryCounter == 29 {
						// Register failed transaction
						a.sendWarning("Unable to get verification of processed transaction from the network. Please try again later.")
						a.log.Errorf("Unable to get status from the network on transaction: %s", TXHash)
						a.RT.Events.Emit("tx_pending", status.Error)
						if err := a.DB.Table("tx_histories").Where("hash = ?", TXHash).Updates(map[string]interface{}{"status": status.Error, "failed": true}).Error; err != nil {
							a.log.Errorln("Unable to query database object for the imported wallet. Reason: ", err)
							a.LoginError("Unable to query database object for the imported wallet.")
							return false
						}
						a.RT.Events.Emit("update_tx_history", []models.TXHistory{}) // Clear TX history
						a.initTXFromDB()
						return false
					}

					consensus = 0 // Reset consensus
				}
				if processed && consensus != 3 {
					consensus++
					a.log.Infof("TX status check has reached consensus %v/3", consensus)
					time.Sleep(1 * time.Second)
				}
				if processed && consensus == 3 { // Need five consecetive confirmations that TX has been processed.
					break
				}

			}
			a.log.Infof("Transaction %v has been successfully processed", TXHash)
			a.sendSuccess("Transaction " + TXHash[:30] + "... has been successfully processed")
			if err := a.DB.Table("tx_histories").Where("hash = ?", TXHash).UpdateColumn("status", status.Complete).Error; err != nil {
				a.log.Errorln("Unable to query database object for the imported wallet. Reason: ", err)
				a.LoginError("Unable to query database object for the imported wallet.")
				return false
			}
			a.RT.Events.Emit("tx_pending", status.Complete)
			a.RT.Events.Emit("update_tx_history", []models.TXHistory{}) // Clear TX history
			a.initTXFromDB()
			return true

		}()
	}
}
