package main

import (
	"os"

	"runtime"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	if runtime.GOOS != "windows" {
		os.Setenv("PATH", "/usr/bin:/sbin") // This is neccessary when interacting with the CLI on RedHat and other linux distros
	}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	frontend := wails.CreateApp(&wails.AppConfig{
		Width:     1280,
		Height:    780,
		Resizable: true,
		Title:     "Molly - Constellation Desktop Wallet",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})

	frontend.Bind(&WalletApplication{})
	err := frontend.Run()
	if err != nil {
		panic(err)
	}

}
