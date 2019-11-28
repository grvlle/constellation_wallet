package main

import (
	"io"
	"os"
	"os/user"
	"path/filepath"

	"github.com/leaanthony/mewn"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
)

func main() {
	os.Setenv("PATH", "/usr/bin:/sbin") // This is neccessary when interacting with the CLI on RedHat and other linux distros

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	frontend := wails.CreateApp(&wails.AppConfig{
		Width:  1524,
		Height: 815,
		Title:  "Molly - Constellation Desktop Wallet [Beta]",
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
	log    *logrus.Logger
	Wallet *Wallet
	paths  struct {
		HomeDir        string
		DAGDir         string
		EncryptedDir   string
		DecKeyFile     string
		PubKeyFile     string
		EncPrivKeyFile string
		LastTXFile     string
		AddressFile    string
	}
}

// WailsInit initializes the Client and Server side bindings
func (a *WalletApplication) WailsInit(runtime *wails.Runtime) error {

	a.RT = runtime
	a.log = logrus.New()

	a.initDirectoryStructure()
	a.initLogger()
	a.initWallet()
	a.initTransactionHistory()

	// Initializes a struct containing all Chart Data on the dashboard
	chartData := a.ChartDataInit()

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

// Write log to STDOUT and a.paths.DAGDir/wallet.log
func (a *WalletApplication) initLogger() {
	logFile, err := os.OpenFile(a.paths.DAGDir+"/wallet.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		a.log.Fatal("Unable to create log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	a.log.SetOutput(mw)
	a.log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

// Initializes the Directory Structure and stores the paths to the WalletApplication struct.
func (a *WalletApplication) initDirectoryStructure() error {

	user, err := user.Current()
	if err != nil {
		a.sendError("Unable to retrieve filesystem paths. Reason: ", err)
		a.log.Error("Unable to retrieve filesystem paths. Reason: ", err)
	}

	a.paths.HomeDir = user.HomeDir             // Home directory of the user
	a.paths.DAGDir = a.paths.HomeDir + "/.dag" // DAG directory for configuration files and wallet specific data
	a.paths.EncryptedDir = a.paths.DAGDir + "/encrypted_key"
	a.paths.DecKeyFile = a.paths.DAGDir + "/private_decrypted.pem" // DAG wallet keys
	a.paths.PubKeyFile = a.paths.EncryptedDir + "/pub.pem"
	a.paths.EncPrivKeyFile = a.paths.EncryptedDir + "/priv.p12"
	a.paths.AddressFile = a.paths.DAGDir + "/addr" // DAG wallet keys
	a.paths.LastTXFile = a.paths.DAGDir + "/acct"  // Account information

	a.log.Info("DAG Directory: " + a.paths.DAGDir)

	err = os.MkdirAll(a.paths.DAGDir, os.ModePerm)
	if err != nil {
		return err
	}
	path := filepath.Join(a.paths.DAGDir, "txhistory.json")
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY,
		0666,
	)
	defer f.Close()

	if !fileExists(path) {
		f.WriteString("{}") // initialies empty JSON object for frontend parsing
		f.Sync()
	}

	return nil
}

// initWallet initializes the Wallet struct and the
func (a *WalletApplication) initWallet() *Wallet {
	return a.NewWallet()
}

func (a *WalletApplication) sendError(msg string, err error) {
	a.RT.Events.Emit("error_handling", msg, err.Error())
}
