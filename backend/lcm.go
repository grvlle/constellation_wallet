package app

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Task int

type Signal struct {
	Status string
	Action string
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
	// Listen to TPC connections on port 1234
	listener, err := net.Listen("tcp", ":1234")
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

func (t *Task) ShutDown(signal Signal, response *Signal) error {
	fmt.Println(signal)
	return nil
}
