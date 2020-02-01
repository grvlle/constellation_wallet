package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func (a *WalletApplication) formTXChain(amount float64, fee float64, address string, ptxObj *Transaction, ltxObj *Transaction) {

	statusLastTX := TXHistory{}
	if err := a.DB.Last(&statusLastTX).Error; err != nil {
		a.log.Warnln("No previous TX detected for this wallet. Reason: ", err)
	}

	a.log.Infoln(statusLastTX.Failed)

	// Queries the number of previous transactions for this wallet.
	numberOfTX := a.DB.Model(&a.wallet).Association("TXHistory").Count()

	// First TX does not contain a TXref
	if numberOfTX == 0 {
		a.log.Infoln("Detected that this is the first TX sent from this key.")
		a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.EmptyTXFile)
		a.sendTransaction(a.paths.LastTXFile)
		return
	}

	// Manually control the second TX, to ensure the following order
	if numberOfTX == 1 {
		// If the first transaction failed, enforce the order
		if statusLastTX.Failed {
			a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.EmptyTXFile)
			a.sendTransaction(a.paths.LastTXFile)
		}
		a.produceTXObject(amount, fee, address, a.paths.PrevTXFile, a.paths.LastTXFile)
		a.sendTransaction(a.paths.PrevTXFile)
		return
	}

	newTX := a.determineBlockOrder(ptxObj, ltxObj)

	// If the last TX is in failed state, we reset the order.
	if newTX == a.paths.PrevTXFile && statusLastTX.Failed {
		a.produceTXObject(amount, fee, address, a.paths.PrevTXFile, a.paths.LastTXFile)
		a.sendTransaction(a.paths.PrevTXFile)
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
	a.log.Info("ltx: ", ltxObj.LastTxRef.Ordinal, "ptx: ", ptxObj.LastTxRef.Ordinal)
	if ltxObj.LastTxRef.Ordinal > ptxObj.LastTxRef.Ordinal {
		return a.paths.PrevTXFile
	}
	return a.paths.LastTXFile

}

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

// collectTXHistory is called by initTransactionHistory to read and parse the LastTXFile.
// It will scan the lines and WalletApplicationend them to txObjects which is later returned to
// initTransactionHistory
func loadTXFromFile(txFile string) string {
	var txObjects string
	file, err := os.Open(txFile) // acct
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
