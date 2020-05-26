package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {

	// initialize update.log file and set log output to file
	file, err := os.OpenFile(getDefaultDagFolderPath()+"/update.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

type Update struct {
	clientRPC       *rpc.Client
	dagFolderPath   *string
	mollyBinaryPath *string
	currentVersion  *string
	newVersion      *string
	triggerUpdate   *bool
}

// Signal is used for IPC with MollyWallet
type Signal struct {
	Status string
	Action string
}

func main() {
	var update Update

	// MollyWallet provides the below data when an update is triggered
	update.dagFolderPath = flag.String("init_dag_path", getDefaultDagFolderPath(), "Enter path to dag folder")
	update.mollyBinaryPath = flag.String("init_molly_path", "", "Enter path to dag folder")
	update.newVersion = flag.String("new_version", "", "Enter the new semantic version. E.g 1.2.3")
	update.currentVersion = flag.String("current_version", "", "Enter the current semantic version. E.g 1.2.3")
	update.triggerUpdate = flag.Bool("upgrade", false, "Upgrade molly wallet binary")
	flag.Parse()

	// if trigger update, update molly
	update.Run()

	fmt.Printf("Dag Folder: %s, Current Version: %s, Molly Path: %s, New Version: %s, Update: %v\n", *update.dagFolderPath, *update.currentVersion, *update.mollyBinaryPath, *update.newVersion, *update.triggerUpdate)
}

func (u *Update) Run() {
	var err error

	// Create a TCP connection to localhost on port 1234
	u.clientRPC, err = rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	log.Infof("Successfully established RPC connection with Molly Wallet")

}

func (u *Update) TerminateAppService() {
	sig := Signal{"Finish App", "Started"}
	var reply Signal

	err := u.clientRPC.Call("Task.ShutDown", sig, &reply)
	if err != nil {
		log.Errorf("Unable to send call: %v", err)
	}
	log.Infof("Recieved response: %v", reply)
}

func backupApp() {

}

func restoreApp() {

}

func downloadAppBinary() {

}

func replaceAppBinary() {

}

func launchAppBinary() {
	// if errors trigger restoreApp
}

func getDefaultDagFolderPath() string {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Errorf("Unable to detect UserHomeDir: %v", err)
		return ""
	}
	return userDir + "/.dag"
}

// Molly wallet binary queries server for updates

// If new version, binary run update-module, passing in new version.

// Molly wallet binary shuts down.

// Update-module waits for binary to stop (check OS running processes).

// Update-module backs up old binary and other files.

// Update-module downloads and replaces binary and any other necessary files.

// Update-module relaunches app.

// If update-module ever encountered an error, restore old files and launch old app
