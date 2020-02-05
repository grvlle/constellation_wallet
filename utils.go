package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/rand"
  "net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

  "github.com/dustin/go-humanize"
)

type WriteCounter struct {
  Total     uint64
  LastEmit  uint64
  Filename  string
  a         *WalletApplication
}

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
				for c := 0; c <= 20; c++ {
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

// Convert byte slice to float64
func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
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
  url := a.WalletCLI.URL + "-v" + a.WalletCLI.Version + "/" + filename
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

  if err = os.Rename(filepath + ".tmp", filepath); err != nil {
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
