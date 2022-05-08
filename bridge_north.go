package ronykit

// Gateway is main component of the EdgeServer. Without Gateway, the EdgeServer is not functional. You can use
// some standard bundles in std/bundle path. However, if you need special handling of communication
// between your server and the clients you are free to implement your own Gateway.
type Gateway interface {
	Bundle
	Dispatcher
	// Subscribe will be called by the EdgeServer. These delegate functions
	// must be called by the Gateway implementation. In other words, Gateway communicates
	// with EdgeServer through the GatewayDelegate methods.
	//
	// NOTE: This func will be called only once and before calling Start function.
	Subscribe(d GatewayDelegate)
}

// GatewayDelegate is the delegate that connects the Gateway to the rest of the system.
type GatewayDelegate interface {
	// OnOpen must be called whenever a new connection is established.
	OnOpen(c Conn)
	// OnClose must be called whenever the connection is gone.
	OnClose(connID uint64)
	// OnMessage must be called whenever a new message arrives.
	OnMessage(c Conn, msg []byte)
}

type northBridge struct {
	ctxPool
	l  Logger
	b  Gateway
	eh ErrHandler
}

var _ GatewayDelegate = (*northBridge)(nil)

func (n *northBridge) OnOpen(c Conn) {
	// Maybe later we can do something
}

func (n *northBridge) OnClose(connID uint64) {
	// Maybe later we can do something
}

func (n *northBridge) OnMessage(c Conn, msg []byte) {
	ctx := n.acquireCtx(c)

	err := n.b.Dispatch(ctx, msg, n.execFunc)
	if err != nil {
		n.eh(ctx, err)
	}

	n.releaseCtx(ctx)

	return
}

func (n *northBridge) execFunc(ctx *Context, arg ExecuteArg) {
	ctx.wf = arg.WriteFunc
	ctx.handlers = append(ctx.handlers, arg.HandlerFuncChain...)
	ctx.Next()

	return
}
