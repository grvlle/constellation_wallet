package app

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Task int

type Signal struct {
	Status string
	Msg    string
}

// initRPCServer initialized the RPC server that listens to incoming LCM tasks
// by the RPC clients
func initRPCServer() error {
	task := new(Task)
	// Publish the receivers methods
	err := rpc.Register(task)
	if err != nil {
		return fmt.Errorf("Format of service Task isn't correct. Reason: %v", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 36866
	listener, err := net.Listen("tcp", ":36866")
	if err != nil {
		return fmt.Errorf("Listen error: %v", err)
	}
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		return fmt.Errorf("Error serving: %v", err)
	}
	return nil
}

func (t *Task) ShutDown(sig Signal, response *Signal) error {
	fmt.Println(sig.Msg)
	// *response = Signal{"OK", "Shutting down application"}
	os.Exit(0)
	return nil
}
