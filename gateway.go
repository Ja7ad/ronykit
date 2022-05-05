package ronykit

// Gateway is main component of the EdgeServer. Without Gateway, the EdgeServer is not functional. You can use
// some standard bundles in std/bundle path. However, if you need special handling of communication
// between your server and the clients you are free to implement your own Gateway.
type Gateway interface {
	// Start starts the bundle to accept connections.
	Start()
	// Shutdown shuts down the the bundle gracefully.
	Shutdown()
	// Subscribe will be called by the EdgeServer the gateway. These delegate functions
	// must be called by the Gateway implementation.
	//
	// NOTE: This func will be called only once and before calling Start function.
	Subscribe(d GatewayDelegate)
	// Dispatch receives the messages from external clients and returns DispatchFunc
	// The user of the Gateway does not need to implement this. If you are using some standard bundles
	// like std/bundle/fasthttp or std/bundle/fastws then all the implementation is taken care of.
	Dispatch(ctx *Context, in []byte, execFunc ExecuteFunc) error
	Register(svc Service)
}

type (
	WriteFunc   func(conn Conn, e *Envelope) error
	ExecuteFunc func(ctx *Context, wf WriteFunc, handlers ...HandlerFunc)
)

// GatewayDelegate is the delegate that connects the Gateway to the rest of the system.
type GatewayDelegate interface {
	// OnOpen must be called whenever a new connection is established.
	OnOpen(c Conn)
	// OnClose must be called whenever the connection is gone.
	OnClose(connID uint64)
	// OnMessage must be called whenever a new message arrives.
	OnMessage(c Conn, msg []byte)
}

// Encoding defines the encoding of the messages which will be sent/received. Gateway implementor needs
// to call correct method based on the encoding value.
type Encoding int32

const (
	Undefined Encoding = 0
	JSON               = 1 << iota
	Proto
	Binary
	Text
	CustomDefined
)

func (e Encoding) Support(o Encoding) bool {
	return e&o == o
}
