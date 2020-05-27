package app

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
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

	errs := make(chan error)

	// Start accept incoming HTTP connections
	go func() {
		err = http.Serve(listener, nil)
		if err != nil {
			errs <- fmt.Errorf("Error serving: %v", err)
			return
		}
	}()

	select {
	case err := <-errs:
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

func (t *Task) ShutDown(sig Signal, response *Signal) error {
	fmt.Println(sig.Msg)
	*response = Signal{"OK", "Shutting down application"}
	time.Sleep(time.Second * 10)
	// TODO: A more graceful shutdown (WailsShutdown)
	os.Exit(0)
	return nil
}
