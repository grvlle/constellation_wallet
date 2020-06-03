package api

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type RPCEndpoints int

type Signal struct {
	PID int
	Msg string
}

// InitRPCServer initializes the RPC server that listens to incoming LCM tasks
// by the RPC clients
func InitRPCServer() error {
	endpoints := new(RPCEndpoints)
	// Publish the receivers methods
	err := rpc.Register(endpoints)
	if err != nil {
		return fmt.Errorf("format of service Task isn't correct. Reason: %v", err)
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

func (rpc *RPCEndpoints) ShutDown(sig Signal, response *Signal) error {

	pid := os.Getpid()
	*response = Signal{pid, "Shutting down application"}

	// time.Sleep(3 * time.Second)
	return nil
}
