package main

import (
	"os"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"runtime"
)

func main() {
	if runtime.GOOS != "windows" {
		os.Setenv("PATH", "/usr/bin:/sbin") // This is neccessary when interacting with the CLI on RedHat and other linux distros
	}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	frontend := wails.CreateApp(&wails.AppConfig{
		Width:  1530,
		Height: 815,
		Title:  "Molly - Constellation Desktop Wallet [Beta]",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	frontend.Bind(&WalletApplication{})
	frontend.Run()
}
