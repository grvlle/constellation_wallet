package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

// WRT stands for Wails Runtime that is the Client/Server event bus
type WRT struct {
	RT *wails.Runtime
}

// WailsInit initializes the Client and Server side bindings
func (w *Wallet) WailsInit(runtime *wails.Runtime) error {
	WailsRuntimeObject := &WRT{}
	WailsRuntimeObject.RT = runtime

	runtime.Window.SetTitle("Constellation Desktop Wallet")

	w.BlockAmount(runtime)
	w.TokenAmount(runtime)
	w.PricePoller(runtime)

	return nil
}

func main() {

	wallet := NewWallet()
	wallet.GetAddress()

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
