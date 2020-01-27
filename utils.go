package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

func (a *WalletApplication) detectJavaPath() {

	if runtime.GOOS == "windows" {

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
			a.LoginError("Unable to find Java Installation")
		}
		jPath := out.String() // Path to java.exe
		if jPath == "" {
			a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
			a.log.Errorln("Unable to detect your Java Path. Please make sure that Java is installed.")
		}
		jwPath := string(jPath[:len(jPath)-6]) + "w.exe" // Shifting to javaw.exe
		a.log.Infoln("Java path detected: " + jwPath)
		a.log.Debugln(cmd)
		a.paths.Java = jwPath
	}
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
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
