package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type UpdateWallet struct {
	currentVersion  string
	newVersion      string
	mollyBinaryPath string
	dagFolderPath   string
	app             *WalletApplication
}

func (u *UpdateWallet) Run() {
	var err error

	u.currentVersion = u.app.Version
	u.newVersion = u.GetLatestRelease()
	u.mollyBinaryPath, err = os.Executable()
	if err != nil {
		u.app.log.Errorln("Unable to collect the path of the molly wallet binary. Reason: ", err)
	}
	u.dagFolderPath = u.app.paths.DAGDir

	err = u.TriggerUpdate()
	if err != nil {
		u.app.log.Errorln("Unable to Update Molly Wallet. Reason: ", err)
		u.app.sendError("Unable to Update Molly Wallet. Reason: ", err)
	}

}

func (u *UpdateWallet) TriggerUpdate() error {

	_, fileExt := getUserOS()

	main := u.dagFolderPath + "/update" + fileExt
	args := []string{"-init_dag_path=" + u.dagFolderPath, "-init_molly_path=" + u.mollyBinaryPath, "-new_version=" + u.newVersion, "-upgrade=" + "true"}

	cmd := exec.Command(main, args...)
	u.app.log.Infoln("Running command: ", cmd)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr // Captures STDERR

	err := cmd.Run()
	if err != nil {
		errFormatted := fmt.Sprint(err) + ": " + stderr.String()
		return errors.New(errFormatted)
	}
	return nil
}

func (u *UpdateWallet) GetLatestRelease() string {

	const (
		url = "https://api.github.com/repos/StardustCollective/molly_wallet/releases/latest"
	)

	resp, err := http.Get(url)
	if err != nil {
		u.app.log.Errorln("Failed to send HTTP request. Reason: ", err)
		return ""
	}
	if resp == nil {
		u.app.log.Errorln("Unable to get the latest release. Empty response from Github API, Reason: ", err)
		u.app.sendWarning("Unable to get the latest release. Please check your internet connectivity and restart the wallet application.")
		return ""
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		u.app.log.Warn("Unable to get the latest release. Reason: ", err)
		return ""
	}
	if bodyBytes == nil {
        u.app.log.Warn("Unable to get the latest release. Reason: Empty response from server")
        return ""
    }

	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return ""
	}

	release := result["tag_name"]
	bytes := []byte(release.(string))
	version := string(bytes[1:6])
	return version

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
