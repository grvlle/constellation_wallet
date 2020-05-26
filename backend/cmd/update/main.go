package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"os"
	"runtime"

	"github.com/artdarek/go-unzip"
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
	log.SetLevel(log.InfoLevel)
}

type Update struct {
	clientRPC       *rpc.Client
	downloadURL     string
	dagFolderPath   *string
	mollyBinaryPath *string
	currentVersion  *string
	newVersion      *string
	triggerUpdate   *bool
}

// Signal is used for IPC with MollyWallet
type Signal struct {
	Status string
	Msg    string
}

func main() {
	var update Update

	update.downloadURL = "https://github.com/grvlle/constellation_wallet/releases/download"

	// MollyWallet provides the below data when an update is triggered
	update.dagFolderPath = flag.String("init_dag_path", getDefaultDagFolderPath(), "Enter path to dag folder")
	update.mollyBinaryPath = flag.String("init_molly_path", "", "Enter path to dag folder")
	update.newVersion = flag.String("new_version", "", "Enter the new semantic version. E.g 1.2.3")
	update.currentVersion = flag.String("current_version", "", "Enter the current semantic version. E.g 1.2.3")
	update.triggerUpdate = flag.Bool("upgrade", false, "Upgrade molly wallet binary")
	flag.Parse()

	// if trigger update, update molly
	// if errors trigger RestoreBackup
	update.Run()

	fmt.Printf("Dag Folder: %s, Current Version: %s, Molly Path: %s, New Version: %s, Update: %v\n", *update.dagFolderPath, *update.currentVersion, *update.mollyBinaryPath, *update.newVersion, *update.triggerUpdate)
}

func (u *Update) Run() {
	var err error

	// Create a TCP connection to localhost on port 36866
	u.clientRPC, err = rpc.DialHTTP("tcp", "localhost:36866")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	log.Infof("Successfully established RPC connection with Molly Wallet")

	zippedArchive, err := u.DownloadAppBinary()
	if err != nil {
		log.Errorf("Unable to download v%s of Molly Wallet: %v", *u.newVersion, err)
	}

	unzipArchive(zippedArchive, *u.dagFolderPath)

	err = u.TerminateAppService()
	if err != nil {
		log.Errorf("Unable to terminate Molly Wallet: %v", err)
	}

}

func (u *Update) DownloadAppBinary() (string, error) {

	fmt.Println("DOWNLOAD")

	filename := "mollywallet.zip"
	osBuild, _ := getUserOS() // returns linux, windows, darwin or unsupported as well as the file extension (e.g ".exe")

	if osBuild == "unsupported" {
		return "", fmt.Errorf("The OS is not supported.")
	}

	url := u.downloadURL + "/v" + *u.newVersion + "-" + osBuild + "/" + filename
	// e.g https://github.com/grvlle/constellation_wallet/releases/download/v1.1.9-linux/mollywallet.zip
	log.Infof("Constructed the following URL: %s", url)

	out, err := os.Create(*u.dagFolderPath + "/" + filename + ".tmp")
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return "", err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		out.Close()
		return "", err
	}

	out.Close()

	if err = os.Rename(*u.dagFolderPath+"/"+filename+".tmp", *u.dagFolderPath+"/"+filename); err != nil {
		return "", err
	}

	return *u.dagFolderPath + "/" + filename, nil
}

func (u *Update) TerminateAppService() error {
	sig := Signal{"OK", "Terminate Molly Wallet. Begining Update..."}
	var response Signal

	err := u.clientRPC.Call("Task.ShutDown", sig, &response)
	if err != nil {
		return err
	}

	if response.Status != "OK" {
		return fmt.Errorf(response.Msg)
	}
	return nil
}

func backupApp() {

}

func replaceAppBinary() {

}

func launchAppBinary() {

}

func (u *Update) RestoreBackup() {

}

func getDefaultDagFolderPath() string {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Errorf("Unable to detect UserHomeDir: %v", err)
		return ""
	}
	return userDir + "/.dag"
}

// Unzips archive to dstPath, returns path to wallet binary
func unzipArchive(zippedArchive, dstPath string) (string, error) {
	uz := unzip.New(zippedArchive, dstPath+"/"+"new_build/")
	err := uz.Extract()
	if err != nil {
		return "", err
	}

	_, fileExt := getUserOS()

	return dstPath + "/" + "new_build/mollywallet" + fileExt, err
}

// returns the users OS as well as the file extension
func getUserOS() (string, string) {
	var osBuild string
	var fileExt string

	switch os := runtime.GOOS; os {
	case "darwin":
		osBuild = "darwin"
		fileExt = ""
	case "linux":
		osBuild = "linux"
		fileExt = ""
	case "windows":
		osBuild = "windows"
		fileExt = ".exe"
	default:
		osBuild = "unsupported"
		fileExt = ""
	}

	return osBuild, fileExt
}

// Molly wallet binary queries server for updates

// If new version, binary run update-module, passing in new version.

// Update-module downloads the new version to tmp

// Molly wallet binary shuts down when download is complete.

// Update-module waits for binary to stop (check OS running processes).

// Update-module backs up old binary and other files.

// Replaces binary and any other necessary files.

// Update-module relaunches app.

// If update-module ever encountered an error, restore old files and launch old app
