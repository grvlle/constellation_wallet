package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	fmt.Println("Hej!!")
	return "Tjo!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1424,
		Height: 768,
		Title:  "Constellation Wallet",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
