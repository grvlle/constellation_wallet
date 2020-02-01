package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func (a *WalletApplication) sendTX(amount float64, fee float64, address string, ptxObj *Transaction, ltxObj *Transaction) {

	newTX := a.determineBlockOrder(ptxObj, ltxObj)
	a.log.Infoln(newTX)

	if a.SecondTX {
		a.produceTXObject(amount, fee, address, a.paths.PrevTXFile, a.paths.LastTXFile)
		a.sendTransaction(amount, fee, address, a.paths.LastTXFile)
		a.FirstTX = false
		a.SecondTX = false
		return
	}

	if a.FirstTX {
		a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.EmptyTXFile)
		a.sendTransaction(amount, fee, address, a.paths.LastTXFile)
		a.FirstTX = false
		a.SecondTX = true
		return
	}

	if newTX != a.paths.PrevTXFile {
		a.produceTXObject(amount, fee, address, a.paths.LastTXFile, a.paths.PrevTXFile)
		a.sendTransaction(amount, fee, address, a.paths.LastTXFile)
		return
	}
	a.produceTXObject(amount, fee, address, a.paths.PrevTXFile, a.paths.LastTXFile)
	a.sendTransaction(amount, fee, address, a.paths.PrevTXFile)
}

func (a *WalletApplication) determineBlockOrder(ptxObj, ltxObj *Transaction) string {

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
