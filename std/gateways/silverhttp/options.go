package silverhttp

import "github.com/clubpay/ronykit/kit"

type Option func(b *bundle)

func WithServerName(name string) Option {
	return func(b *bundle) {
		b.srvName = name
	}
}

func WithBufferSize(size int64) Option {
	return func(b *bundle) {
		b.srv.MaxBodySize = size
	}
}

func WithLogger(l kit.Logger) Option {
	return func(b *bundle) {
		b.l = l
	}
}

func Listen(addr string) Option {
	return func(b *bundle) {
		b.listen = addr
	}
}

func WithCORS(cfg CORSConfig) Option {
	return func(b *bundle) {
		b.cors = newCORS(cfg)
	}
}

func WithPredicateKey(key string) Option {
	return func(b *bundle) {
		b.predicateKey = key
	}
}

func WithWebsocketEndpoint(endpoint string) Option {
	return func(b *bundle) {
		b.wsEndpoint = endpoint
	}
}

func WithCustomRPC(in kit.IncomingRPCFactory, out kit.OutgoingRPCFactory) Option {
	return func(b *bundle) {
		b.rpcInFactory = in
		b.rpcOutFactory = out
	}
}
