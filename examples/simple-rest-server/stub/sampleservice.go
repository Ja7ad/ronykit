// Code generated by RonyKIT Stub Generator (Golang); DO NOT EDIT.

package sampleservice

import (
	"context"
	"fmt"

	"github.com/clubpay/ronykit"
	"github.com/clubpay/ronykit/stub"
	"github.com/clubpay/ronykit/utils/reflector"
)

var _ fmt.Stringer

func init() {
	reflector.Register(&EchoRequest{}, "json")
	reflector.Register(&EchoResponse{}, "json")
	reflector.Register(&EmbeddedHeader{}, "json")
	reflector.Register(&ErrorMessage{}, "json")
	reflector.Register(&RedirectRequest{}, "json")
	reflector.Register(&SumRequest{}, "json")
	reflector.Register(&SumResponse{}, "json")
}

// EchoRequest is a data transfer object
type EchoRequest struct {
	RandomID int64 `json:"randomID"`
	Ok       bool  `json:"ok"`
}

// EchoResponse is a data transfer object
type EchoResponse struct {
	RandomID int64 `json:"randomID"`
	Ok       bool  `json:"ok"`
}

// EmbeddedHeader is a data transfer object
type EmbeddedHeader struct {
	SomeKey1 string `json:"someKey1"`
	SomeInt1 int64  `json:"someInt1"`
}

// ErrorMessage is a data transfer object
type ErrorMessage struct {
	Code int    `json:"code"`
	Item string `json:"item"`
}

func (x ErrorMessage) GetCode() int {
	return x.Code
}

func (x ErrorMessage) GetItem() string {
	return x.Item
}

// RedirectRequest is a data transfer object
type RedirectRequest struct {
	URL string `json:"url"`
}

// SumRequest is a data transfer object
type SumRequest struct {
	EmbeddedHeader
	Val1 int64 `json:"val1"`
	Val2 int64 `json:"val2"`
}

// SumResponse is a data transfer object
type SumResponse struct {
	EmbeddedHeader
	Val int64 `json:"val"`
}

type ISampleServiceStub interface {
	EchoGET(
		ctx context.Context, req *EchoRequest, opt ...stub.RESTOption,
	) (*EchoResponse, *stub.Error)
	EchoPOST(
		ctx context.Context, req *EchoRequest, opt ...stub.RESTOption,
	) (*EchoResponse, *stub.Error)
	Sum1(
		ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
	) (*SumResponse, *stub.Error)
	Sum2(
		ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
	) (*SumResponse, *stub.Error)
	SumRedirect(
		ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
	) (*SumResponse, *stub.Error)
}

// SampleServiceStub represents the client/stub for SampleService.
// Implements ISampleServiceStub
type SampleServiceStub struct {
	hostPort  string
	secure    bool
	verifyTLS bool

	s *stub.Stub
}

func NewSampleServiceStub(hostPort string, opts ...stub.Option) *SampleServiceStub {
	s := &SampleServiceStub{
		s: stub.New(hostPort, opts...),
	}

	return s
}

var _ ISampleServiceStub = (*SampleServiceStub)(nil)

func (s SampleServiceStub) EchoGET(
	ctx context.Context, req *EchoRequest, opt ...stub.RESTOption,
) (*EchoResponse, *stub.Error) {
	res := &EchoResponse{}
	httpCtx := s.s.REST().
		SetMethod("GET").
		SetResponseHandler(
			400,
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				res := &ErrorMessage{}
				err := stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
				if err != nil {
					return err
				}

				return stub.NewErrorWithMsg(res)
			},
		).
		DefaultResponseHandler(
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				return stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
			},
		).
		AutoRun(ctx, "/echo/:randomID", kit.JSON, req, opt...)
	defer httpCtx.Release()

	if err := httpCtx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s SampleServiceStub) EchoPOST(
	ctx context.Context, req *EchoRequest, opt ...stub.RESTOption,
) (*EchoResponse, *stub.Error) {
	res := &EchoResponse{}
	httpCtx := s.s.REST().
		SetMethod("POST").
		SetResponseHandler(
			400,
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				res := &ErrorMessage{}
				err := stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
				if err != nil {
					return err
				}

				return stub.NewErrorWithMsg(res)
			},
		).
		DefaultResponseHandler(
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				return stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
			},
		).
		AutoRun(ctx, "/echo-post", kit.JSON, req, opt...)
	defer httpCtx.Release()

	if err := httpCtx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s SampleServiceStub) Sum1(
	ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
) (*SumResponse, *stub.Error) {
	res := &SumResponse{}
	httpCtx := s.s.REST().
		SetMethod("GET").
		SetResponseHandler(
			400,
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				res := &ErrorMessage{}
				err := stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
				if err != nil {
					return err
				}

				return stub.NewErrorWithMsg(res)
			},
		).
		DefaultResponseHandler(
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				return stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
			},
		).
		AutoRun(ctx, "/sum/:val1/:val2", kit.JSON, req, opt...)
	defer httpCtx.Release()

	if err := httpCtx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s SampleServiceStub) Sum2(
	ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
) (*SumResponse, *stub.Error) {
	res := &SumResponse{}
	httpCtx := s.s.REST().
		SetMethod("POST").
		SetResponseHandler(
			400,
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				res := &ErrorMessage{}
				err := stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
				if err != nil {
					return err
				}

				return stub.NewErrorWithMsg(res)
			},
		).
		DefaultResponseHandler(
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				return stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
			},
		).
		AutoRun(ctx, "/sum", kit.JSON, req, opt...)
	defer httpCtx.Release()

	if err := httpCtx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s SampleServiceStub) SumRedirect(
	ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
) (*SumResponse, *stub.Error) {
	res := &SumResponse{}
	httpCtx := s.s.REST().
		SetMethod("GET").
		SetResponseHandler(
			400,
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				res := &ErrorMessage{}
				err := stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
				if err != nil {
					return err
				}

				return stub.NewErrorWithMsg(res)
			},
		).
		DefaultResponseHandler(
			func(ctx context.Context, r stub.RESTResponse) *stub.Error {
				return stub.WrapError(kit.UnmarshalMessage(r.GetBody(), res))
			},
		).
		AutoRun(ctx, "/sum-redirect/:val1/:val2", kit.JSON, req, opt...)
	defer httpCtx.Release()

	if err := httpCtx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

type MockOption func(*SampleServiceStubMock)

func MockEchoGET(
	f func(ctx context.Context, req *EchoRequest, opt ...stub.RESTOption) (*EchoResponse, *stub.Error),
) MockOption {
	return func(sm *SampleServiceStubMock) {
		sm.echoget = f
	}
}

func MockEchoPOST(
	f func(ctx context.Context, req *EchoRequest, opt ...stub.RESTOption) (*EchoResponse, *stub.Error),
) MockOption {
	return func(sm *SampleServiceStubMock) {
		sm.echopost = f
	}
}

func MockSum1(
	f func(ctx context.Context, req *SumRequest, opt ...stub.RESTOption) (*SumResponse, *stub.Error),
) MockOption {
	return func(sm *SampleServiceStubMock) {
		sm.sum1 = f
	}
}

func MockSum2(
	f func(ctx context.Context, req *SumRequest, opt ...stub.RESTOption) (*SumResponse, *stub.Error),
) MockOption {
	return func(sm *SampleServiceStubMock) {
		sm.sum2 = f
	}
}

func MockSumRedirect(
	f func(ctx context.Context, req *SumRequest, opt ...stub.RESTOption) (*SumResponse, *stub.Error),
) MockOption {
	return func(sm *SampleServiceStubMock) {
		sm.sumredirect = f
	}
}

// SampleServiceStubMock represents the mocked for client/stub for SampleService.
// Implements ISampleServiceStub
type SampleServiceStubMock struct {
	echoget     func(ctx context.Context, req *EchoRequest, opt ...stub.RESTOption) (*EchoResponse, *stub.Error)
	echopost    func(ctx context.Context, req *EchoRequest, opt ...stub.RESTOption) (*EchoResponse, *stub.Error)
	sum1        func(ctx context.Context, req *SumRequest, opt ...stub.RESTOption) (*SumResponse, *stub.Error)
	sum2        func(ctx context.Context, req *SumRequest, opt ...stub.RESTOption) (*SumResponse, *stub.Error)
	sumredirect func(ctx context.Context, req *SumRequest, opt ...stub.RESTOption) (*SumResponse, *stub.Error)
}

func NewSampleServiceStubMock(opts ...MockOption) *SampleServiceStubMock {
	s := &SampleServiceStubMock{}
	for _, o := range opts {
		o(s)
	}

	return s
}

var _ ISampleServiceStub = (*SampleServiceStubMock)(nil)

func (s SampleServiceStubMock) EchoGET(
	ctx context.Context, req *EchoRequest, opt ...stub.RESTOption,
) (*EchoResponse, *stub.Error) {
	if s.echoget == nil {
		return nil, stub.WrapError(fmt.Errorf("method not mocked"))
	}

	return s.echoget(ctx, req, opt...)
}

func (s SampleServiceStubMock) EchoPOST(
	ctx context.Context, req *EchoRequest, opt ...stub.RESTOption,
) (*EchoResponse, *stub.Error) {
	if s.echopost == nil {
		return nil, stub.WrapError(fmt.Errorf("method not mocked"))
	}

	return s.echopost(ctx, req, opt...)
}

func (s SampleServiceStubMock) Sum1(
	ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
) (*SumResponse, *stub.Error) {
	if s.sum1 == nil {
		return nil, stub.WrapError(fmt.Errorf("method not mocked"))
	}

	return s.sum1(ctx, req, opt...)
}

func (s SampleServiceStubMock) Sum2(
	ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
) (*SumResponse, *stub.Error) {
	if s.sum2 == nil {
		return nil, stub.WrapError(fmt.Errorf("method not mocked"))
	}

	return s.sum2(ctx, req, opt...)
}

func (s SampleServiceStubMock) SumRedirect(
	ctx context.Context, req *SumRequest, opt ...stub.RESTOption,
) (*SumResponse, *stub.Error) {
	if s.sumredirect == nil {
		return nil, stub.WrapError(fmt.Errorf("method not mocked"))
	}

	return s.sumredirect(ctx, req, opt...)
}
