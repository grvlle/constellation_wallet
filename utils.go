package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/mr-tron/base58"
)

type txInformation struct {
	ID              int    `json:"id,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	Address         string `json:"address,omitempty"`
	Fee             int    `json:"fee,omitempty"`
	TransactionHash string `json:"txhash,omitempty"`
	TS              string `json:"date,omitempty"`
}

// monitorFileState will monitor the state of all files in JSONdata
// and act accordingly upon manipulation.
func (a *App) monitorFileState() error {
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
				if event.Op&fsnotify.Write == fsnotify.Write {
					a.log.Infof("modified file: %s", event.Name)
					switch fileModified := event.Name; {

					case fileModified == a.paths.DAGDir+"/acct":
						tx := &Transaction{}

						file, err := os.Open(a.paths.DAGDir + "/acct")
						if err != nil {
							a.log.Errorf("Unable to read tx data. Reason: %s", err)
						}
						defer file.Close()

						scanner := bufio.NewScanner(file)
						scanner.Split(bufio.ScanLines)
						var txObjects []string

						for scanner.Scan() {
							txObjects = append(txObjects, scanner.Text())
						}

						file.Close()

						for _, eachTX := range txObjects {
							fmt.Println(eachTX)
							bytes := []byte(eachTX)
							// if err != nil {
							// 	a.log.Errorf("Unable to read contents of acct. Reason: %s", err)
							// }
							err = json.Unmarshal(bytes, &tx)
							if err != nil {
								a.log.Errorf("Unable to parse contents of acct. Reason: %s", err)
							}
							txData := &txInformation{
								ID:              tx.Edge.Count,
								Amount:          tx.Edge.Data.Amount,
								Address:         tx.Edge.ObservationEdge.Parents[0].Hash,
								Fee:             tx.Edge.Data.Fee,
								TransactionHash: tx.Edge.ObservationEdge.Data.Hash,
								TS:              time.Now().Format("Mon Jan _2 15:04:05 2006"),
							}

							a.RT.Events.Emit("new_transaction", txData) // Pass the tx to the frontend as a new transaction.
						}
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

func (a *App) collectOSPath() error {
	user, err := user.Current()
	if err != nil {
		a.log.Errorf("Unable to retrieve fs paths. Reason: ", err)
	}

	a.paths.HomeDir = user.HomeDir
	a.paths.DAGDir = a.paths.HomeDir + "/.dag"
	a.paths.KeyFile = a.paths.DAGDir + "/key"
	a.paths.LastTXFile = a.paths.DAGDir + "/acct"

	return nil
}

// writeToJSON is a helper function that will remove a requested file(filename),
// and recreate it with new data(data). This is to avoid ticking off the
// monitorFileState function with double write events.
func writeToJSON(filename string, data interface{}) error {

	JSON, err := json.Marshal(data)

	path := filepath.Join(".", "JSONdata", filename)
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
func setupDirectoryStructure() error {
	err := os.MkdirAll("JSONdata", os.ModePerm)
	if err != nil {
		return err
	}

	return err
}

// Base58Encode is used to hash out the keys
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

// Base58Decode decoder
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}
	return decode
}
