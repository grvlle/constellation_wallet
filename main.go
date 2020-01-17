package main

import (
	"io"
	"os"
	"os/user"

	"github.com/jinzhu/gorm"
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
	DB     *gorm.DB
	paths  struct {
		HomeDir        string
		DAGDir         string
		EncryptedDir   string
		DecKeyFile     string
		PubKeyFile     string
		EncPrivKeyFile string
		LastTXFile     string
		AddressFile    string
		ImageDir       string
	}
	UserLoggedIn bool
}

// WailsInit initializes the Client and Server side bindings
func (a *WalletApplication) WailsInit(runtime *wails.Runtime) error {
	var err error

	a.UserLoggedIn = false
	a.RT = runtime
	a.log = logrus.New()
	a.DB, err = gorm.Open("sqlite3", "/home/vito/.dag/store.db")
	if err != nil {
		a.log.Panicf("failed to connect database", err)
	}
	// Migrate the schema
	a.DB.AutoMigrate(&Wallet{}, &TXHistory{})

	a.initDirectoryStructure()
	a.initLogger()
	// Monitors the .dag folder for file manipulation
	err = a.monitorFileState()
	if err != nil {
		return err
	}
	return nil
}

// WailsShutdown is called when the application is closed
func (a *WalletApplication) WailsShutdown() {
	a.DB.Close()
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
func (a *WalletApplication) initDirectoryStructure() {

	user, err := user.Current()
	if err != nil {
		a.sendError("Unable to retrieve filesystem paths. Reason: ", err)
		a.log.Error("Unable to retrieve filesystem paths. Reason: ", err)
	}

	a.paths.HomeDir = user.HomeDir             // Home directory of the user
	a.paths.DAGDir = a.paths.HomeDir + "/.dag" // DAG directory for configuration files and wallet specific data
	a.paths.EncryptedDir = a.paths.DAGDir + "/keys"
	a.paths.DecKeyFile = a.paths.EncryptedDir + "/private_decrypted" // DAG wallet keys
	a.paths.PubKeyFile = a.paths.EncryptedDir + "/decrypted_keystore.pub"
	a.paths.EncPrivKeyFile = a.paths.EncryptedDir + "/key.p12"
	a.paths.AddressFile = a.paths.DAGDir + "/addr"  // DAG wallet keys
	a.paths.ImageDir = "./frontend/src/assets/img/" // Image Folder

	a.log.Info("DAG Directory: ", a.paths.DAGDir)

	err = a.directoryCreator(a.paths.DAGDir, a.paths.EncryptedDir)
	if err != nil {
		a.sendError("Unable to set up directory structure. Make sure you run the wallet with the right priviledges. Reason: ", err)
		a.log.Errorf("Unable to set up directory structure. Make sure you run the wallet with the right priviledges. Reason: ", err)
	}
}

// initWallet initializes the Wallet data post-login.
func (a *WalletApplication) initWallet() error {

	a.Wallet = &Wallet{
		Balance:          1024155,
		AvailableBalance: 1012233,
		Nonce:            420,
		TotalBalance:     1012420,
		Delegated:        42,
		Deposit:          0,
		Address:          "",
	}

	a.Wallet.PrivateKey.Key, a.Wallet.PublicKey.Key = a.getKeys()
	a.Wallet.Address = a.createAddressFromPublicKey()

	//a.initTransactionHistory()

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

	return nil
}

func (a *WalletApplication) sendError(msg string, err error) {
	errStr := string(err.Error())
	a.RT.Events.Emit("error_handling", msg, errStr)
}
