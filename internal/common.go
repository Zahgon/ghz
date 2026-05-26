package internal

import (
	"google.golang.org/grpc"

	"github.com/bojand/ghz/internal/gtime"
	"github.com/bojand/ghz/internal/helloworld"
	"github.com/bojand/ghz/internal/sleep"
	"github.com/bojand/ghz/internal/wrapped"
)

// TestPort is the port.
var TestPort string

// TestLocalhost is the localhost.
var TestLocalhost string

// StartServer starts the server.
//
// For testing only.
func StartServer(secure bool) (*helloworld.Greeter, *grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// StartSleepServer starts the sleep test server
func StartSleepServer(secure bool) (*sleep.SleepService, *grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// StartWrappedServer starts the wrapped test server
func StartWrappedServer(secure bool) (*wrapped.WrappedService, *grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// StartTimeServer starts the wrapped test server
func StartTimeServer(secure bool) (*gtime.TimeService, *grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
