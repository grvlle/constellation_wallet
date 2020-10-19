package app

import (
	"encoding/json"
	"io"
	"os"
	"os/user"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"

	"github.com/grvlle/constellation_wallet/backend/api"
	"github.com/grvlle/constellation_wallet/backend/models"
)

// WalletApplication holds all application specific objects
// such as the Client/Server event bus and logger
type WalletApplication struct {
	Version    string
	RT         *wails.Runtime
	log        *logrus.Logger
	wallet     models.Wallet
	DB         *gorm.DB
	killSignal chan struct{}
	Network    struct {
		URL     string
		Handles struct {
			Send        string // Takes TX Object, returns TX Hash (200)
			Transaction string // Takes TX Object, returns TX Hash (200)
			Balance     string // Polls the wallets available balance
		}
		BlockExplorer struct {
			URL     string
			Handles struct {
				Transactions string // Takes TX Hash, returns TX info
				Checkpoints  string // Takes Checkpoint block hash, returns checkpoint block info
				Snapshots    string // Takes SnapshotHash, returns info
				CollectTX    string // Takes DAG address, returns tx objects

			}
		}
	}
	paths struct {
		HomeDir      string
		DAGDir       string
		TMPDir       string
		EncryptedDir string
		EmptyTXFile  string
		PrevTXFile   string
		LastTXFile   string
		AddressFile  string
		ImageDir     string
		Java         string
	}
	KeyStoreAccess      bool
	TransactionFinished bool
	TransactionFailed   bool
	UserLoggedIn        bool
	NewUser             bool
	WalletImported      bool
	FirstTX             bool
	SecondTX            bool
	WidgetRunning       struct {
		PassKeysToFrontend bool
		DashboardWidgets   bool
	}
	WalletCLI struct {
		URL     string
		Version string
	}
}

// Constants of the application
const (
	ServiceLogin = "molly-wallet-login"
	ServiceSeed = "molly-wallet-seed"
	ServicePKey = "molly-wallet-pkey"

	MainnetBlockExplorerURL = "https://block-explorer.constellationnetwork.io"
	TestnetBlockExplorerURL = "https://api-be.exchanges.constellationnetwork.io"
	MainnetLoadBalancerURL = "http://lb.constellationnetwork.io:9000"
	TestnetLoadBalancerURL = "http://lb.exchanges.constellationnetwork.io:9000"
)

// WailsShutdown is called when the application is closed
func (a *WalletApplication) WailsShutdown() {
	a.wallet = models.Wallet{}
	close(a.killSignal) // Kills the Go Routines
	a.DB.Close()
}

// WailsInit initializes the Client and Server side bindings
func (a *WalletApplication) WailsInit(runtime *wails.Runtime) error {
	var err error

	a.log = logrus.New()
	err = a.initDirectoryStructure()
	if err != nil {
		a.log.Errorln("Unable to set up directory structure. Reason: ", err)
	}

	a.initLogger()

	err = api.InitRPCServer()
	if err != nil {
		a.log.Panicf("Unable to initialize RPC Server. Reason: %v", err)
	}
	a.log.Infoln("RPC Server initialized.")

	a.UserLoggedIn = false
	a.NewUser = false
	a.TransactionFinished = true
	a.RT = runtime
	a.killSignal = make(chan struct{}) // Used to kill go routines and hand back system resources
	a.wallet.Currency = "USD"          // Set default currency
	a.WalletCLI.URL = "https://github.com/Constellation-Labs/constellation/releases/download"
	a.WalletCLI.Version = "2.13.9"
	a.Version = "1.4.1"

	a.DB, err = gorm.Open("sqlite3", a.paths.DAGDir+"/store.db")
	if err != nil {
		a.log.Panicln("failed to connect database", err)
	}
	// Migrate the schema
	a.DB.AutoMigrate(&models.Wallet{}, &models.TXHistory{}, &models.Path{}, &models.Contact{})
	a.detectJavaPath()
	a.initMainnetConnection()
	a.newReleaseAvailable()

	return nil
}

// initLogger writes logs to STDOUT and a.paths.DAGDir/wallet.log
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
		return err
	}

	a.paths.HomeDir = user.HomeDir             // Home directory of the user
	a.paths.DAGDir = a.paths.HomeDir + "/.dag" // DAG directory for configuration files and wallet specific data
	a.paths.TMPDir = a.paths.DAGDir + "/tmp"
	a.paths.LastTXFile = a.paths.TMPDir + "/last_tx"
	a.paths.PrevTXFile = a.paths.TMPDir + "/prev_tx"
	a.paths.EmptyTXFile = a.paths.TMPDir + "/genesis_tx"
	a.paths.ImageDir = "./frontend/src/assets/img/" // Image Folder

	a.log.Info("DAG Directory: ", a.paths.DAGDir)

	err = a.directoryCreator(a.paths.DAGDir, a.paths.TMPDir)
	if err != nil {
		return err
	}

	return nil
}

// initMainnetConnection populates the WalletApplication struct with mainnet data
func (a *WalletApplication) initMainnetConnection() {
	a.Network.URL = MainnetLoadBalancerURL // Temp

	a.Network.Handles.Send = "/send"
	a.Network.Handles.Transaction = "/transaction"
	a.Network.Handles.Balance = "/address/"

	a.Network.BlockExplorer.URL = MainnetBlockExplorerURL
	a.Network.BlockExplorer.Handles.Transactions = "/transactions/"
	a.Network.BlockExplorer.Handles.Checkpoints = "/checkpoints/"
	a.Network.BlockExplorer.Handles.Snapshots = "/snapshots/"
	// a.Network.BlockExplorer.Handles.CollectTX = "/transactions?sender="
}

// APIError reported by the blockexplerer/loadbalancer are reported in the following format
// {"error": "Cannot find transactions for sender"}
type APIError struct {
	Error string
}

// verifyAPIResponse takes API response converted to a byte array and checks if the API returned
// an error. If it did, it'll return the error message.
func (a *WalletApplication) verifyAPIResponse(r []byte) (bool, string) {
	APIErr := APIError{}
	if string(r[3:8]) == "error" {
		err := json.Unmarshal(r, &APIErr)
		if err != nil {
			a.log.Errorln("Unable to parse API error. Reason: ", err)
		}
		return false, APIErr.Error
	}
	return true, ""
}

func (a *WalletApplication) sendSuccess(msg string) {

	if len(msg) > 200 {
		msg = string(msg[:200]) // Restrict error size for frontend
		a.RT.Events.Emit("success", msg)
		return
	}
	a.RT.Events.Emit("success", msg)
}

func (a *WalletApplication) sendWarning(msg string) {

	if len(msg) > 200 {
		msg = string(msg[:200]) // Restrict error size for frontend
		a.RT.Events.Emit("warning", msg)
		return
	}
	a.RT.Events.Emit("warning", msg)
}

func (a *WalletApplication) sendError(msg string, err error) {

	var errStr string

	if err != nil {
		b := []byte(err.Error())
		if len(b) > 80 {
			errStr = string(b[:80]) // Restrict error size for frontend
		} else if len(b) < 80 {
			errStr = string(b)
		} else {
			errStr = ""
		}

		a.RT.Events.Emit("error_handling", msg, errStr+" ...")
	} else {
		a.RT.Events.Emit("error_handling", msg+" ...")
	}

}
