package kit_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/clubpay/ronykit/kit"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRonykit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ronykit Suite")
}

type testConn struct {
	id       uint64
	clientIP string
	stream   bool
	kv       map[string]string
	buf      *bytes.Buffer
}

var _ kit.Conn = (*testConn)(nil)

func newTestConn(id uint64, clientIP string, stream bool) *testConn {
	return &testConn{
		id:       id,
		clientIP: clientIP,
		stream:   stream,
		kv:       map[string]string{},
		buf:      &bytes.Buffer{},
	}
}

func (t testConn) ConnID() uint64 {
	return t.id
}

func (t testConn) ClientIP() string {
	return t.clientIP
}

func (t testConn) Write(data []byte) (int, error) {
	return t.buf.Write(data)
}

func (t testConn) Read() ([]byte, error) {
	return io.ReadAll(t.buf)
}

func (t testConn) ReadString() (string, error) {
	d, err := t.Read()
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (t testConn) Stream() bool {
	return t.stream
}

func (t testConn) Walk(f func(key string, val string) bool) {
	for k, v := range t.kv {
		if !f(k, v) {
			return
		}
	}
}

func (t testConn) Get(key string) string {
	return t.kv[key]
}

func (t testConn) Set(key string, val string) {
	t.kv[key] = val
}

func (t testConn) Keys() []string {
	keys := make([]string, 0, len(t.kv))
	for k := range t.kv {
		keys = append(keys, k)
	}

	return keys
}

// testRESTSelector implements kit.RESTRouteSelector for testing purposes.
type testRESTSelector struct {
	enc    kit.Encoding
	method string
	path   string
}

func (t testRESTSelector) Query(q string) interface{} {
	return nil
}

func (t testRESTSelector) GetEncoding() kit.Encoding {
	return t.enc
}

func (t testRESTSelector) GetMethod() string {
	return t.method
}

func (t testRESTSelector) GetPath() string {
	return t.path
}

// testRPCSelector implements kit.RPCSelector for testing purposes.
type testRPCSelector struct {
	enc       kit.Encoding
	predicate string
}

func (t testRPCSelector) Query(q string) interface{} {
	return nil
}

func (t testRPCSelector) GetEncoding() kit.Encoding {
	return t.enc
}

func (t testRPCSelector) GetPredicate() string {
	return t.predicate
}
