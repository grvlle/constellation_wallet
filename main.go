package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	wallet := &Wallet{}
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1524,
		Height: 815,
		Title:  "Constellation Wallet",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(wallet)
	app.Run()
}
