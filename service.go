package ronykit

// Service defines a set of RPC handlers which usually they are related to one service.
// Name must be unique per each Gateway.
type Service interface {
	// Name of the service which must be unique per EdgeServer.
	Name() string
	// Contracts returns a list of APIs which this service provides.
	Contracts() []Contract
}

// ServiceWrapper lets you add customizations to your service. A specific case of it is serviceInterceptor
// which can add Pre- and Post- handlers to all the Contracts of the Service.
type ServiceWrapper interface {
	Wrap(Service) Service
}

type ServiceWrapperFunc func(Service) Service

func (sw ServiceWrapperFunc) Wrap(svc Service) Service {
	return sw(svc)
}

// WrapService wraps a service, this is useful for adding middlewares to the service.
// Some middlewares like OpenTelemetry, Logger, ... could be added to the service using
// this function.
func WrapService(svc Service, wrappers ...ServiceWrapper) Service {
	for _, w := range wrappers {
		svc = w.Wrap(svc)
	}

	return svc
}

// ServiceGenerator generates a service. desc.Service is the implementor of this.
type ServiceGenerator interface {
	Generate() Service
}

// Contract defines the set of Handlers based on the Query. Query is different per bundles,
// hence, this is the implementor's task to make sure return correct value based on 'q'.
// In other words, Contract 'r' must return valid response for 'q's required by Gateway 'b' in
// order to be usable by Gateway 'b' otherwise it panics.
type Contract interface {
	RouteSelector() RouteSelector
	MemberSelector() MemberSelector
	Encoding() Encoding
	Input() Message
	Handlers() []HandlerFunc
	Modifiers() []Modifier
}

// ContractWrapper is like an interceptor which can add Pre- and Post- handlers to all
// the Contracts of the Contract.
type ContractWrapper interface {
	Wrap(Contract) Contract
}

// ContractWrapperFunc implements ContractWrapper interface.
type ContractWrapperFunc func(Contract) Contract

func (sw ContractWrapperFunc) Wrap(svc Contract) Contract {
	return sw(svc)
}

// WrapContract wraps a contract, this is useful for adding middlewares to the contract.
// Some middlewares like OpenTelemetry, Logger, ... could be added to the contract using
// this function.
func WrapContract(c Contract, wrappers ...ContractWrapper) Contract {
	for _, w := range wrappers {
		c = w.Wrap(c)
	}

	return c
}

// RouteSelector holds information about how this Contract is going to be selected. Each
// Gateway may need different information to route the request to the right Contract.
type RouteSelector interface {
	Query(q string) interface{}
}

// RESTRouteSelector defines the RouteSelector which could be used in REST operations
// Gateway implementation which handle REST requests could check the selector if it supports REST.
type RESTRouteSelector interface {
	RouteSelector
	GetMethod() string
	GetPath() string
}

// RPCRouteSelector defines the RouteSelector which could be used in RPC operations.
// Gateway implementation which handle RPC requests could check the selector if it supports RPC
type RPCRouteSelector interface {
	RouteSelector
	GetPredicate() string
}
