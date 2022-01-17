package ronykit

import (
	"sync"

	"github.com/ronaksoft/ronykit/utils"
)

type Envelope struct {
	kvl utils.SpinLock
	kv  map[string]string
	m   Message
	p   *sync.Pool
}

func (e *Envelope) SetHdr(key, value string) *Envelope {
	e.kvl.Lock()
	e.kv[key] = value
	e.kvl.Unlock()

	return e
}

func (e *Envelope) SetHdrMap(kv map[string]string) *Envelope {
	e.kvl.Lock()
	for k, v := range kv {
		e.kv[k] = v
	}
	e.kvl.Unlock()

	return e
}

func (e *Envelope) GetHdr(key string) string {
	e.kvl.Lock()
	v := e.kv[key]
	e.kvl.Unlock()

	return v
}

func (e *Envelope) WalkHdr(f func(key string, val string) bool) {
	e.kvl.Lock()
	for k, v := range e.kv {
		if !f(k, v) {
			break
		}
	}
	e.kvl.Unlock()
}

func (e *Envelope) SetMsg(msg Message) *Envelope {
	e.m = msg

	return e
}

func (e *Envelope) GetMsg() Message {
	if e.m == nil {
		return nil
	}

	return e.m
}

func (e *Envelope) Release() {
	for k := range e.kv {
		delete(e.kv, k)
	}
	e.m = nil

	e.p.Put(e)
}

var envelopePool = &sync.Pool{}

func NewEnvelope() *Envelope {
	e, ok := envelopePool.Get().(*Envelope)
	if !ok {
		e = &Envelope{
			kv: map[string]string{},
			p:  envelopePool,
		}
	}

	return e
}
