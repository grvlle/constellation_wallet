package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
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
		if len(jPath) <= 0 {
			go func() {
				for c := 0; c >= 10; c++ {
					a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
					time.Sleep(1 * time.Second)
				}
			}()
			a.log.Errorln("Unable to detect your Java Path. Please make sure that Java is installed.")
			return
		}
		jwPath := string(jPath[:len(jPath)-6]) + "w.exe" // Shifting to javaw.exe
		a.log.Infoln("Java path detected: " + jwPath)
		a.log.Debugln(cmd)
		a.paths.Java = jwPath
	}
}

func (a *WalletApplication) TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(a.paths.TMPDir, prefix+hex.EncodeToString(randBytes)+suffix)
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

func (a *WalletApplication) directoryCreator(directories ...string) error {
	for _, d := range directories {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
