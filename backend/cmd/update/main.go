package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
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
	clientRPC          *rpc.Client
	downloadURL        string
	dagFolderPath      *string
	oldMollyBinaryPath *string
	currentVersion     *string
	newVersion         *string
	triggerUpdate      *bool
}

// Signal is used for IPC with MollyWallet
type Signal struct {
	Status string
	Msg    string
}

type unzippedContents struct {
	newMollyBinaryPath string
	updateBinaryPath   string
}

func main() {
	var update Update

	update.downloadURL = "https://github.com/grvlle/constellation_wallet/releases/download"

	// MollyWallet provides the below data when an update is triggered
	update.dagFolderPath = flag.String("init_dag_path", getDefaultDagFolderPath(), "Enter the directory path to dag folder")
	update.oldMollyBinaryPath = flag.String("init_molly_path", "", "Enter the directory path where the molly binary is located")
	update.newVersion = flag.String("new_version", "", "Enter the new semantic version. E.g 1.2.3")
	update.triggerUpdate = flag.Bool("upgrade", false, "Upgrade molly wallet binary")
	flag.Parse()

	// if trigger update, update molly
	// if errors trigger RestoreBackup
	update.Run()

	fmt.Printf("Dag Folder: %s, Current Version: %s, Molly Path: %s, New Version: %s, Update: %v\n", *update.dagFolderPath, *update.currentVersion, *update.oldMollyBinaryPath, *update.newVersion, *update.triggerUpdate)
}

func (u *Update) Run() {
	var err error

	// Clean up old update artifacts
	err = u.CleanUp()
	if err != nil {
		log.Fatalf("Unable to clear previous local state: %v", err)
	}

	// Create a TCP connection to localhost on port 36866
	u.clientRPC, err = rpc.DialHTTP("tcp", "localhost:36866")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	log.Infof("Successfully established RPC connection with Molly Wallet")
	defer u.clientRPC.Close()

	zippedArchive, err := u.DownloadAppBinary()
	if err != nil {
		log.Fatalf("Unable to download v%s of Molly Wallet: %v", *u.newVersion, err)
	}

	// TODO: checksum verification

	contents, err := unzipArchive(zippedArchive, *u.dagFolderPath)
	if err != nil {
		log.Fatalf("Unable to unzip contents: %v", err)
	}

	err = u.BackupApp()
	if err != nil {
		log.Fatalf("Unable to Backup Molly Wallet: %v", err)
	}

	err = u.TerminateAppService()
	if err != nil {
		log.Fatalf("Unable to terminate Molly Wallet: %v", err)
	}

	err = u.ReplaceAppBinary(contents)
	if err != nil {
		log.Fatalf("Unable to overwrite old installation: %v", err)
	}

	err = u.LaunchAppBinary()
	if err != nil {
		log.Fatalf("Unable to start up Molly after update: %v", err)
	}

	err = u.CleanUp()
	if err != nil {
		log.Fatalf("Unable to clear previous local state: %v", err)
	}

}

// DownloadAppBinary downloads the latest Molly Wallet zip from github releases and returns the path to it
func (u *Update) DownloadAppBinary() (string, error) {

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

// TerminateAppService will send an RPC to mollywallet to terminate the application
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

func (u *Update) BackupApp() error {
	_, fileExt := getUserOS()

	// Create backup folder in ~/.dag
	err := os.Mkdir(*u.dagFolderPath+"/backup", 0755)
	if err != nil {
		return fmt.Errorf("Unable to create backup folder. Reason: %v", err)
	}

	// Copy the old Molly Wallet binary into ~/.dag/backup/
	err = copy(*u.oldMollyBinaryPath+"/mollywallet"+fileExt, *u.dagFolderPath+"/backup/mollywallet"+fileExt)
	if err != nil {
		return fmt.Errorf("Unable to backup old Molly installation. Reason: %v", err)
	}

	return nil
}

func (u *Update) ReplaceAppBinary(contents *unzippedContents) error {
	// Copy the old Molly Wallet binary into ~/.dag/backup/
	_, fileExt := getUserOS()
	err := copy(contents.newMollyBinaryPath, *u.oldMollyBinaryPath+"/mollywallet"+fileExt)
	if err != nil {
		return fmt.Errorf("Unable to overwrite old molly binary. Reason: %v", err)
	}
	return nil
}

func (u *Update) LaunchAppBinary() error {
	_, fileExt := getUserOS()
	cmd := exec.Command(*u.oldMollyBinaryPath + "/mollywallet" + fileExt)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Unable to execute run command for Molly Wallet: %v", err)
	}
	return nil
}

func (u *Update) CleanUp() error {

	if fileExists(*u.dagFolderPath + "/mollywallet.zip") {
		err := os.Remove(*u.dagFolderPath + "/mollywallet.zip")
		if err != nil {
			return err
		}
	}

	err := os.RemoveAll(*u.dagFolderPath + "/backup")
	if err != nil {
		return err
	}

	if fileExists(*u.dagFolderPath + "/new_build") {
		err := os.RemoveAll(*u.dagFolderPath + "/new_build")
		if err != nil {
			return err
		}
	}
	return nil
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

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func copy(src string, dst string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Unzips archive to dstPath, returns path to wallet binary
func unzipArchive(zippedArchive, dstPath string) (*unzippedContents, error) {

	uz := unzip.New(zippedArchive, dstPath+"/"+"new_build/")
	err := uz.Extract()
	if err != nil {
		return nil, err
	}
	_, fileExt := getUserOS()

	contents := &unzippedContents{
		newMollyBinaryPath: dstPath + "/" + "new_build/mollywallet" + fileExt,
		updateBinaryPath:   dstPath + "/" + "new_build/update" + fileExt,
	}

	return contents, err
}

// getUserOS returns the users OS as well as the file extension of executables for said OS
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
