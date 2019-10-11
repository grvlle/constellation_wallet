package main

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// monitorFileState will monitor the state of all files in .dag
// and act accordingly upon manipulation.
func (a *WalletApplication) monitorFileState() error {
	a.log.Info("Starting Watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				} // If a JSONdata/*.json file is written to.
				if event.Op&fsnotify.Write&fsnotify.Create == fsnotify.Write|fsnotify.Create {
					a.log.Infof("modified file: %s", event.Name)
					switch fileModified := event.Name; {

					case fileModified == a.paths.LastTXFile:
						a.log.Debug("Last TX File has been modified")
						// tx := &Transaction{}

						// file, err := os.Open(a.paths.LastTXFile)
						// if err != nil {
						// 	a.log.Errorf("Unable to read tx data. Reason: %s", err)
						// }
						// defer file.Close()

						// scanner := bufio.NewScanner(file)
						// scanner.Split(bufio.ScanLines)
						// var txObjects []string

						// for scanner.Scan() {
						// 	txObjects = append(txObjects, scanner.Text())
						// }

						// file.Close()

						// a.RT.Events.Emit("clear_transactions", true)

						// for _, eachTX := range txObjects {
						// 	// fmt.Println(eachTX)
						// 	bytes := []byte(eachTX)
						// 	err = json.Unmarshal(bytes, &tx)
						// 	if err != nil {
						// 		a.log.Warnf("Unable to parse contents of acct. Reason: %s", err)
						// 	}
						// 	txData := &txInformation{
						// 		ID:              tx.Edge.Count,
						// 		Amount:          tx.Edge.Data.Amount,
						// 		Address:         tx.Edge.ObservationEdge.Parents[0].Hash,
						// 		Fee:             tx.Edge.Data.Fee,
						// 		TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
						// 		TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
						// 	}

						// 	a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.
						// }

					case fileModified == a.paths.KeyFile:
						a.log.Debug("Key File has been modified")
						a.RT.Events.Emit("wallet_keys", a.Wallet.PrivateKey.Key, a.Wallet.PublicKey.Key)

					case fileModified == "JSONdata/chart_data.json":
						a.log.Info("Chart Data file modified")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					a.log.Error(err.Error())
				}
			}
		}
	}()

	err = watcher.Add(a.paths.DAGDir)
	if err != nil {
		return err
	}
	return nil
}

func (a *WalletApplication) collectOSPath() error {
	user, err := user.Current()
	if err != nil {
		a.log.Errorf("Unable to retrieve fs paths. Reason: ", err)
	}

	a.paths.HomeDir = user.HomeDir
	a.paths.DAGDir = a.paths.HomeDir + "/.dag"
	a.paths.KeyFile = a.paths.DAGDir + "/key"
	a.paths.LastTXFile = a.paths.DAGDir + "/acct"

	a.log.Info("The following paths will be used:\n Home Directory: " + a.paths.HomeDir + "\n DAG Directory: " + a.paths.DAGDir + "\n KeyFile: " + a.paths.KeyFile + "\n Transactions File: " + a.paths.LastTXFile + "\n")

	return nil
}

// writeToJSON is a helper function that will remove a requested file(filename),
// and recreate it with new data(data). This is to avoid ticking off the
// monitorFileState function with double write events.
func writeToJSON(filename string, data interface{}) error {
	user, err := user.Current()
	if err != nil {
		return err
	}
	JSON, err := json.Marshal(data)
	path := filepath.Join(user.HomeDir+"/.dag", filename)
	os.Remove(path)

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY,
		0666,
	)
	defer f.Close()

	f.Write(JSON)
	f.Sync()

	if err != nil {
		return err
	}
	return nil
}

// This function is called by WailsInit and will initialize the dir structure.
func (a *WalletApplication) setupDirectoryStructure() error {
	err := os.MkdirAll(a.paths.DAGDir, os.ModePerm)
	if err != nil {
		return err
	}
	return err
}
