package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/mr-tron/base58"
)

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

					case fileModified == "JSONdata/tx.json":
						tx := &Transaction{}

						file, err := os.Open("JSONdata/tx.json")
						if err != nil {
							a.log.Errorf("Unable to open JSONdata/tx.json. Reason: %s", err)
						}
						defer file.Close()

						bytes, err := ioutil.ReadAll(file)
						if err != nil {
							a.log.Errorf("Unable to read contents of JSONdata/tx.json. Reason: %s", err)
						}
						json.Unmarshal(bytes, &tx)
						a.RT.Events.Emit("new_transaction", tx) // Pass the tx to the frontend as a new transaction.

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

	err = watcher.Add("JSONdata")
	if err != nil {
		return err
	}
	return nil
}

// writeToJSON is a helper function that will remove a requested file(filename),
// and recreate it with new data(data). This is to avoid ticking off the
// monitorFileState function with double write events.
func writeToJSON(filename string, data interface{}) error {

	JSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

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
