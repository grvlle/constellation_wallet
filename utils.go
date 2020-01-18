package main

import (
	"encoding/json"
	"io/ioutil"
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

					case fileModified == a.paths.DecKeyFile:
						a.log.Debug("Key File has been modified")
						a.RT.Events.Emit("wallet_keys", a.Wallet.PrivateKey, a.Wallet.PublicKey)

					case fileModified == "JSONdata/chart_data.json":
						a.log.Info("Chart Data file modified")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					a.sendError("", err)
					a.log.Error(err.Error())
				}
			}
		}
	}()

	err = watcher.Add(a.paths.DAGDir)
	if err != nil {
		a.sendError("Failed to start watcher. Reason: ", err)
		return err
	}
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

func (a *WalletApplication) directoryCreator(directories ...string) error {
	for _, d := range directories {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *WalletApplication) getFileContents(filePath string) ([]byte, error) {
	path := filepath.Join(filePath)
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return fileContents, nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func reverseElement(elements []*txInformation) []*txInformation {
	reversed := []*txInformation{}
	for i := range elements {
		n := elements[len(elements)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}
