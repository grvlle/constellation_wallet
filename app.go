package main

import (
	"io"
	"os"
	"os/user"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
)

// WalletApplication holds all application specific objects
// such as the Client/Server event bus and logger
type WalletApplication struct {
	RT     *wails.Runtime
	log    *logrus.Logger
	wallet Wallet
	DB     *gorm.DB
	paths  struct {
		HomeDir        string
		DAGDir         string
		EncryptedDir   string
		EncPrivKeyFile string
		PrevTXFile     string
		LastTXFile     string
		AddressFile    string
		ImageDir       string
	}
	KeyStoreAccess bool
	UserLoggedIn   bool
	NewUser        bool
	WidgetRunning  struct {
		PassKeysToFrontend bool
		DashboardWidgets   bool
	}
}

// WailsShutdown is called when the application is closed
func (a *WalletApplication) WailsShutdown() {
	a.DB.Close()
}

// WailsInit initializes the Client and Server side bindings
func (a *WalletApplication) WailsInit(runtime *wails.Runtime) error {
	var err error

	a.log = logrus.New()
	err = a.initDirectoryStructure()
	if err != nil {
		a.log.Errorf("Unable to set up directory structure. Reason: ", err)
	}

	a.initLogger()

	a.UserLoggedIn = false
	a.NewUser = false
	a.RT = runtime

	a.DB, err = gorm.Open("sqlite3", a.paths.DAGDir+"/store.db")
	if err != nil {
		a.log.Panicf("failed to connect database", err)
	}
	// Migrate the schema
	a.DB.AutoMigrate(&Wallet{}, &TXHistory{})

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
	a.paths.EncryptedDir = a.paths.DAGDir + "/keys"
	a.paths.EncPrivKeyFile = a.paths.EncryptedDir + "/key.p12"
	a.paths.LastTXFile = a.paths.DAGDir + "/last_tx"
	a.paths.PrevTXFile = a.paths.DAGDir + "/prev_tx"
	a.paths.ImageDir = "./frontend/src/assets/img/" // Image Folder

	a.log.Info("DAG Directory: ", a.paths.DAGDir)

	err = a.directoryCreator(a.paths.DAGDir, a.paths.EncryptedDir)
	if err != nil {
		return err
	}

	files := []string{a.paths.LastTXFile, a.paths.PrevTXFile}

	for _, f := range files {
		file, err := os.Create(f)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
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
	}
}
