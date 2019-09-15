package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	a := &App{}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1524,
		Height: 815,
		Title:  "App Wallet",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(a)
	app.Bind(a.wallet)
	app.Bind(sendTransaction)
	app.Run()
}

// App stands for Wails Runtime that is the Client/Server event bus
type App struct {
	RT    *wails.Runtime
	log   *wails.CustomLogger
	paths struct {
		HomeDir    string
		DAGDir     string
		KeyFile    string
		LastTXFile string
	}
	wallet *Wallet
}

// WailsInit initializes the Client and Server side bindings
func (a *App) WailsInit(runtime *wails.Runtime) error {
	a.RT = runtime
	a.log = runtime.Log.New("Constellation")
	setupDirectoryStructure()
	a.WalletInit()

	runtime.Window.SetTitle("Constellation Desktop Wallet")

	chartData := ChartDataInit()
	a.nodeStats(a.RT, chartData)

	a.BlockAmount()
	a.TokenAmount()
	a.PricePoller()

	a.collectOSPath()
	a.monitorFileState()

	return nil
}

// WalletInit initializes the Wallet struct and the
func (a *App) WalletInit() *Wallet {
	w := a.NewWallet()
	w.GetAddress()
	return w
}
