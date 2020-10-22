package app

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mcuadros/go-version"
)

// newReleaseAvailable generates a notification to FE everytime a new release on
// GitHub doesn't match a.Version.
func (a *WalletApplication) newReleaseAvailable() {
	update := new(UpdateWallet)
	update.app = a
	currentRelease := a.Version

	a.log.Infoln("Checking for new releases...")

	go func() {
		for i := 200; i > 0; i-- {
			newRelease := update.GetLatestRelease()
			if version.Compare(newRelease, currentRelease, ">") && newRelease != "" {
				a.log.Infoln("There's a newer release available")
				a.RT.Events.Emit("new_release", newRelease)
			}
			time.Sleep(time.Duration(i) * time.Second)
		}
	}()

}

func (a *WalletApplication) updateTokenBalance() error {
	balance, err := a.GetTokenBalance()
	if err != nil {
		a.sendWarning("No data recieved from the Token Balance API. Will try again soon.")
		return err
	}
	a.log.Infoln("Current Balance: ", balance)
	a.wallet.Balance, a.wallet.AvailableBalance, a.wallet.TotalBalance = balance, balance, balance
	a.RT.Events.Emit("token", a.wallet.Balance, a.wallet.AvailableBalance, a.wallet.TotalBalance)

	return nil
}

func (a *WalletApplication) javaInstalled() bool {
	var javaInstalled bool
	if a.paths.Java[len(a.paths.Java)-9:] != "javaw.exe" {
		javaInstalled = false
	} else {
		javaInstalled = true
	}
	return javaInstalled
}

func (a *WalletApplication) detectJavaPath() {

	if runtime.GOOS == "windows" {
		var jwPath string

		cmd := exec.Command("cmd", "/c", "where", "java")
		a.log.Infoln("Running command: ", cmd)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out    // Captures STDOUT
		cmd.Stderr = &stderr // Captures STDERR

		err := cmd.Run()
		if err != nil {
			errFormatted := fmt.Sprint(err) + ": " + stderr.String()
			a.log.Errorf(errFormatted)
			//a.LoginError("Unable to find Java Installation")
			//a.paths.Java = "No valid path detected"
			a.paths.Java = "javaw.exe"
			return
		}
		jPath := out.String() // May contain multiple
		if jPath == "" {
			a.LoginError("Unable to find Java Installation")
			a.paths.Java = "No valid path detected"
			return
		}
		s := strings.Split(strings.Replace(jPath, "\r\n", "\n", -1), "\n")
		jwPath = string(s[0][:len(s[0])-4]) + "w.exe" // Shifting to javaw.exe
		if s[1] != "" {
			jwPath = string(s[1][:len(s[1])-4]) + "w.exe" // Shifting to javaw.exe
			a.log.Info("Detected a secondary java path. Using that over the first one.")
		}
		a.log.Infoln("Java path selected: " + jwPath)
		a.log.Debugln(cmd)
		a.paths.Java = jwPath
	}
}

// TempFileName creates temporary file names for the transaction files
func (a *WalletApplication) TempFileName(prefix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(a.paths.TMPDir, prefix+hex.EncodeToString(randBytes))
}

// WriteCounter stores dl state of the cl binaries
type WriteCounter struct {
	Total    uint64
	LastEmit uint64
	Filename string
	a        *WalletApplication
}

// Write emits the download progress of the CL binaries to the frontend
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)

	if (wc.Total - wc.LastEmit) > uint64(800) {
		wc.a.RT.Events.Emit("downloading", wc.Filename, humanize.Bytes(wc.Total))
		wc.LastEmit = wc.Total
	}

	return n, nil
}

func (a *WalletApplication) fetchWalletJar(filename string, filepath string) error {
	url := a.WalletCLI.URL + "/v" + a.WalletCLI.Version + "/" + filename
	a.log.Info(url)

	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	counter.a = a
	counter.Filename = filename
	counter.LastEmit = uint64(0)

	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
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

func (a *WalletApplication) fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data []byte) error {

	err := ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func WriteStringToFile(filename string, text string) error {
    f, err := os.Create(filename)

    if err != nil {
        return err
    }

    defer f.Close()

    _, err2 := f.WriteString(text)

    if err2 != nil {
        return err2
    }

    return nil
}
