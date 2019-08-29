package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

type FrontEnd struct {
	runtime *wails.Runtime
	logger  *wails.CustomLogger
}

func (f *FrontEnd) WailsInit(runtime *wails.Runtime) error {
	f.logger = runtime.Log.New("FrontEnd")
	f.logger.Info("I'm here")
	f.runtime = runtime
	f.BlockAmount()
	f.TokenAmount()

	return nil
}

func (f *FrontEnd) BlockAmount() error {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(300)
			f.runtime.Events.Emit("blocks", randomNumber)
			time.Sleep(2 * time.Second)
		}
	}()
	return nil
}

func (f *FrontEnd) TokenAmount() error {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(3000)
			f.runtime.Events.Emit("error", "$", randomNumber)
			time.Sleep(2 * time.Second)
		}
	}()
	return nil
}

func basic() string {
	fmt.Println("Hej!!")
	return "Tjo!"
}

func main() {

	fe := &FrontEnd{}

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
	app.Bind(fe)
	app.Run()
}
