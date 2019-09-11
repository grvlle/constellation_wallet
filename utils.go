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

// watcherInit will monitor the state of all files in JSONdata
// and act accordingly upon manipulation.
func (a *App) watcherInit() error {
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
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					a.log.Infof("modified file: %s", event.Name)

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

					a.RT.Events.Emit("new_transaction", tx)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				a.log.Error(err.Error())
			}
		}
	}()

	err = watcher.Add("JSONdata")
	if err != nil {
		return err
	}
	return nil
}

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}
	return decode
}

func writeToJSON(filename string, data interface{}) error {
	JSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.MkdirAll("JSONdata", os.ModePerm)
	if err != nil {
		return err
	}
	path := filepath.Join(".", "JSONdata", filename)
	err = ioutil.WriteFile(path, JSON, 0644)
	if err != nil {
		return err
	}
	return nil
}
