package app

import (
	"encoding/json"

	"github.com/grvlle/constellation_wallet/backend/models"
)

// formTXChain retains the order of the blockchain across all accounts.
// Also calls the methods to create block objects (write json to file),
// and the method that pushes them to the network(HTTP POST).
// To retain order, formTXChain will poll the last sent TX for it's Failed state.
// if the last TX failed, it'll switch up the order to account for that not to break the chain.
// This means that all failed attempts at creating a block is also stored in the DB with
// a Failed state.
func (a *WalletApplication) formTXChain(amount int64, fee int64, address string, ptxObj *Transaction, ltxObj *Transaction) {

	statusLastTX := models.TXHistory{}
	if err := a.DB.Last(&statusLastTX).Error; err != nil {
		a.log.Warnln("No previous TX detected for this wallet. Reason: ", err)
	}

	if statusLastTX.Failed {
		a.log.Warnln("Previous Transaction has a failed state. Adapting...", statusLastTX.Failed)
	}
	// Queries the number of previous transactions for this wallet.
	numberOfTX := a.DB.Model(&a.wallet).Association("TXHistory").Count()

	// First TX does not contain a TXref
	if numberOfTX == 0 {
		a.log.Infoln("Detected that this is the first TX sent from this key.")
		a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.EmptyTXFile)
		a.sendTransaction(a.paths.LastTXFile)
		return
	}

	a.log.Infoln("Number of previous TX's detected: ", numberOfTX)

	// Manually control the second TX, to ensure the following order
	if numberOfTX == 1 {

		// If the first transaction failed, enforce the order.
		if statusLastTX.Failed {
			a.log.Warnln("Detected that the previous transaction failed.")
			a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.EmptyTXFile)
			a.sendTransaction(a.paths.LastTXFile)
			return
		}

		// Check for edge case where PrevTXFile has already been written and needs to be referenced.
		// This occurs when a wallet with 1 previous tx has been imported.
		prevTXFileContents := a.loadTXFromFile(a.paths.PrevTXFile)
		if a.WalletImported && prevTXFileContents != "" {
			a.log.Warnln("One previous transaction has been imported. Using that as reference.")
			a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.PrevTXFile)
			a.sendTransaction(a.paths.LastTXFile)
			return
		}

		a.produceTXObject(amount, fee, address, a.paths.PrevTXFile, a.paths.LastTXFile)
		a.sendTransaction(a.paths.PrevTXFile)
		return
	}

	// Returns the TX object that has the highest ordinal (the highest determines if it's to be referenced or reference the other tx)
	newTX := a.determineBlockOrder(ptxObj, ltxObj)

	// If the last TX is in failed state, we reset the order.
	if newTX == a.paths.PrevTXFile && statusLastTX.Failed {
		a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.PrevTXFile)
		a.sendTransaction(a.paths.LastTXFile)
		return
	}

	if newTX != a.paths.PrevTXFile && !statusLastTX.Failed {
		a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.PrevTXFile)
		a.sendTransaction(a.paths.LastTXFile)
		return
	}
	a.produceTXObject(amount, fee, address, a.paths.PrevTXFile, a.paths.LastTXFile)
	a.sendTransaction(a.paths.PrevTXFile)
}

func (a *WalletApplication) determineBlockOrder(ptxObj, ltxObj *Transaction) string {
	// The higher ordinal will always be the TX carrying the TX Ref.
	a.log.Info("Last TX Ordinal: ", ltxObj.LastTxRef.Ordinal, " Previous TX Ordinal: ", ptxObj.LastTxRef.Ordinal)
	if ltxObj.LastTxRef.Ordinal > ptxObj.LastTxRef.Ordinal {
		return a.paths.PrevTXFile
	}
	return a.paths.LastTXFile

}

// convertToTXObject takes the Path to the prev_tx and last_tx files and returns a
// pointer to two workable objects.
func (a *WalletApplication) convertToTXObject(ptx, ltx string) (*Transaction, *Transaction) {
	var ptxObj Transaction
	var ltxObj Transaction

	rbytes := []byte(ptx)
	lbytes := []byte(ltx)

	err := json.Unmarshal(rbytes, &ptxObj)
	if err != nil {
		a.log.Warnln("TX Object: ", string(rbytes), err)
	}
	err = json.Unmarshal(lbytes, &ltxObj)
	if err != nil {
		a.log.Warnln("TX Object: ", string(rbytes), err)
	}
	return &ptxObj, &ltxObj
}

/* Called upon wallet import */

// TXReference is used to parse the previous tx of an imported wallet.
type TXReference struct {
	Hash               string `json:"hash"`
	Amount             int64  `json:"amount"`
	Receiver           string `json:"receiver"`
	Sender             string `json:"sender"`
	Fee                int    `json:"fee"`
	IsDummy            bool   `json:"isDummy"`
	LastTransactionRef struct {
		PrevHash string `json:"prevHash"`
		Ordinal  int    `json:"ordinal"`
	} `json:"lastTransactionRef"`
	SnapshotHash        string `json:"snapshotHash"`
	CheckpointBlock     string `json:"checkpointBlock"`
	TransactionOriginal struct {
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
				Fee  interface{} `json:"fee"`
				Salt int64       `json:"salt"`
			} `json:"data"`
		} `json:"edge"`
		LastTxRef struct {
			PrevHash string `json:"prevHash"`
			Ordinal  int    `json:"ordinal"`
		} `json:"lastTxRef"`
		IsDummy bool `json:"isDummy"`
		IsTest  bool `json:"isTest"`
	} `json:"transactionOriginal"`
}
