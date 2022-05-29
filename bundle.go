package ronykit

import "context"

// Bundle is the pluggable component of the EdgeServer.
type Bundle interface {
	// Start starts the gateway to accept connections.
	Start(ctx context.Context) error
	// Shutdown shuts down the gateway gracefully.
	Shutdown(ctx context.Context) error
	// Register registers the route in the Bundle. This is how Bundle get information
	// about the services and their contracts.
	Register(serviceName, contractID string, sel RouteSelector, input Message)
}

type (
	ExecuteArg struct {
		WriteFunc
		ServiceName string
		ContractID  string
		Route       string
	}
	WriteFunc   func(conn Conn, e Envelope) error
	ExecuteFunc func(arg ExecuteArg)
)

var NoExecuteArg = ExecuteArg{}
