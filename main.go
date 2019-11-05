package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	frontend := wails.CreateApp(&wails.AppConfig{
		Width:  1524,
		Height: 815,
		Title:  "Constellation Desktop Wallet",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	frontend.Bind(&WalletApplication{})
	frontend.Run()
}

// WalletApplication holds all application specific objects
// such as the Client/Server event bus and logger
type WalletApplication struct {
	RT     *wails.Runtime
	log    *wails.CustomLogger
	Wallet *Wallet
	paths  struct {
		HomeDir    string
		DAGDir     string
		KeyFile    string
		LastTXFile string
	}
}

// WailsInit initializes the Client and Server side bindings
func (a *WalletApplication) WailsInit(runtime *wails.Runtime) error {

	a.RT = runtime
	a.log = runtime.Log.New("Constellation Wallet")

	a.collectOSPath()
	a.setupDirectoryStructure()
	a.walletInit()
	a.initTransactionHistory()

	// Initializes a struct containing all Chart Data on the dashboard
	chartData := ChartDataInit()

	// Below methods are continously updating the client side modules.
	a.nodeStats(chartData)
	a.txStats(chartData)
	a.networkStats(chartData)
	a.blockAmount()
	a.tokenAmount()
	a.pricePoller()
	a.passKeysToFrontend()

	// Monitors the .dag folder for file manipulation
	a.monitorFileState()

	return nil
}

// WalletInit initializes the Wallet struct and the
func (a *WalletApplication) walletInit() *Wallet {
	return a.NewWallet()
}

func (a *WalletApplication) sendError(msg string, err error) {
	e := err.Error()
	a.RT.Events.Emit("error_handling", msg, e)
}
