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
