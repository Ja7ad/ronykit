package ronykit

type envelopeCarrier struct {
	Hdr          map[string]string `json:"hdr"`
	Msg          []byte            `json:"msg"`
	ContractID   string            `json:"cid"`
	ServiceName  string            `json:"svc"`
	ContextIndex int               `json:"idx"`
	IsREST       bool              `json:"isRest"`
}

func envelopeCarrierFromContext(ctx *Context) envelopeCarrier {
	ec := envelopeCarrier{
		Hdr:          map[string]string{},
		Msg:          nil,
		ContractID:   ctx.ContractID(),
		ServiceName:  ctx.ServiceName(),
		ContextIndex: ctx.handlerIndex,
	}
	_, ec.IsREST = ctx.Conn().(RESTConn)
	ctx.In().
		WalkHdr(func(key, val string) bool {
			ec.Hdr[key] = val

			return true
		})

	return ec
}
