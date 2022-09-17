package main

import (
	"math/rand"
	"os"
	"time"

	"runtime"

	"github.com/wailsapp/wails"

	_ "embed"

	app "github.com/grvlle/constellation_wallet/backend"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	if runtime.GOOS != "windows" {
		os.Setenv("PATH", "/usr/bin:/sbin") // This is neccessary when interacting with the CLI on RedHat and other linux distros
	}

	frontend := wails.CreateApp(&wails.AppConfig{
		Width:     1280,
		Height:    780,
		Resizable: true,
		Title:     "Molly Wallet",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})

	frontend.Bind(&app.WalletApplication{})
	err := frontend.Run()
	if err != nil {
		panic(err)
	}

}
