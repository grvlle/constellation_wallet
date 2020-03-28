package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// formTXChain retains the order of the blockchain across all accounts.
// Also calls the methods to create block objects (write json to file),
// and the method that pushes them to the network(HTTP POST).
// To retain order, formTXChain will poll the last sent TX for it's Failed state.
// if the last TX failed, it'll switch up the order to account for that not to break the chain.
// This means that all failed attempts att creating a block is also stored in the DB with
// a Failed state bool.
func (a *WalletApplication) formTXChain(amount int64, fee int64, address string, ptxObj *Transaction, ltxObj *Transaction) {

	statusLastTX := TXHistory{}
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

	fmt.Println(numberOfTX, a.WalletImported)

	// Manually control the second TX, to ensure the following order
	if numberOfTX == 1 {

		// If the first transaction failed, enforce the order.
		if statusLastTX.Failed {
			a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.EmptyTXFile)
			a.sendTransaction(a.paths.LastTXFile)
			return
		}

		// PrevTXFile has already been written and needs to be referenced.
		if a.WalletImported {
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
		fmt.Println(err)
	}
	err = json.Unmarshal(lbytes, &ltxObj)
	if err != nil {
		fmt.Println(err)
	}
	return &ptxObj, &ltxObj
}

/* Called upon wallet import */

// TXReference is used to parse the previous tx of an imported wallet.
type TXReference struct {
	Hash               string `json:"hash"`
	Amount             int    `json:"amount"`
	Receiver           string `json:"receiver"`
	Sender             string `json:"sender"`
	Fee                int    `json:"fee"`
	IsDummy            bool   `json:"isDummy"`
	LastTransactionRef struct {
		Hash    string `json:"hash"`
		Ordinal int    `json:"ordinal"`
	} `json:"lastTransactionRef"`
	SnapshotHash        string `json:"snapshotHash"`
	CheckpointBlock     string `json:"checkpointBlock"`
	TransactionOriginal struct {
		Edge struct {
			ObservationEdge struct {
				Parents []struct {
					Hash     string      `json:"hash"`
					HashType string      `json:"hashType"`
					BaseHash interface{} `json:"baseHash,omitempty"`
				} `json:"parents"`
				Data struct {
					Hash     string      `json:"hash"`
					HashType string      `json:"hashType"`
					BaseHash interface{} `json:"baseHash,omitempty"`
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
				Amount    int `json:"amount"`
				LastTxRef struct {
					Hash    string `json:"hash"`
					Ordinal int    `json:"ordinal"`
				} `json:"lastTxRef"`
				Fee  interface{} `json:"fee,omitempty"`
				Salt int64       `json:"salt"`
			} `json:"data"`
		} `json:"edge"`
		LastTxRef struct {
			Hash    string `json:"hash"`
			Ordinal int    `json:"ordinal"`
		} `json:"lastTxRef"`
		IsDummy bool `json:"isDummy"`
		IsTest  bool `json:"isTest"`
	} `json:"transactionOriginal"`
}

// rebuildTxChainState will query the blockexplorer for a transacion and write it to a.paths.PrevTXFile.
// This will allow an imported wallet to reference the last transaction sent.
func (a *WalletApplication) rebuildTxChainState(lastTXHash string) error {
	a.log.Info("Sending API call to block explorer on: " + a.Network.BlockExplorer.Handles.Transactions)

	resp, err := http.Get(a.Network.BlockExplorer.URL + a.Network.BlockExplorer.Handles.Transactions + lastTXHash)
	if err != nil {
		a.log.Errorln("Failed to send HTTP request. Reason: ", err)
		if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Delete(&a.wallet).Error; err != nil {
			a.log.Errorln("Unable to delete wallet upon failed import. Reason: ", err)
			return err
		}
		a.LoginError("Unable to collect previous TX's from blockexplorer. Please check your internet connection.")
		return err
	}
	defer resp.Body.Close()

	if resp.Body != nil {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			a.LoginError("Unable to collect previous TX's from blockexplorer. Please try again later.")
			a.log.Errorln("Unable to collect previous TX's from blockexplorer. Reason: ", err)
		}
		ok, error := a.verifyAPIResponse(bodyBytes)
		// Blockexplorer returns below string when no previous transactions are found
		if !ok && error != "Cannot find transaction" {
			a.log.Errorln("API returned the following error", error)
			// If unable to import last transaction, remove wallet from DB and logout.
			if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Delete(&a.wallet).Error; err != nil {
				a.log.Errorln("Unable to delete wallet upon failed import. Reason: ", err)
				return err
			}
			a.LoginError("The wallet import failed. Please check your internet connection and try again.")
			return err
		}

		// Parsing JSON object to TXReference ->
		lastTX := TXReference{}
		err = json.Unmarshal(bodyBytes, &lastTX)
		if err != nil {
			a.log.Errorln("Unable to fetch TX history from block explorer. Reason: ", err)
			a.sendError("Unable to fetch TX history from block explorer. Reason: ", err)
			return err
		}
		// Marshal so that we can unmarshat into tx object ->
		b, err := json.Marshal(lastTX.TransactionOriginal)
		if err != nil {
			a.log.Errorln("Unable to parse last transaction hash. Reason: ", err)
			a.sendError("Unable to fetch TX history from block explorer. Reason: ", err)
			return err
		}

		// Populating tx object ->
		tx := Transaction{}
		err = json.Unmarshal(b, &tx)
		if err != nil {
			a.log.Errorln("Unable to parse last transaction hash. Reason: ", err)
			a.sendError("Unable to fetch TX history from block explorer. Reason: ", err)
			return err
		}

		// Converting to json
		txBytes, err := json.Marshal(tx)
		if err != nil {
			a.log.Errorln("Unable to parse last transaction hash. Reason: ", err)
			a.sendError("Unable to fetch TX history from block explorer. Reason: ", err)
			return err
		}

		err = WriteToFile(a.paths.PrevTXFile, txBytes)
		if err != nil {
			return err
		}
	}
	return nil
}
