package api

import (
	"fmt"
	"net/http"

	"github.com/clubpay/ronykit"
	"github.com/clubpay/ronykit/desc"
	"github.com/clubpay/ronykit/exmples/simple-rest-server/dto"
	"github.com/clubpay/ronykit/std/gateway/fasthttp"
	"github.com/clubpay/ronykit/std/gateway/fastws"
	"github.com/goccy/go-reflect"
)

type Sample struct{}

func NewSample() *Sample {
	s := &Sample{}

	return s
}

func (x *Sample) Desc() *desc.Service {
	return desc.NewService("SampleService").
		AddContract(
			desc.NewContract().
				SetInput(&dto.EchoRequest{}).
				AddSelector(fasthttp.Selector{
					Method:    fasthttp.MethodGet,
					Predicate: "echo",
					Path:      "/echo/:randomID",
				}).
				AddSelector(fastws.Selector{
					Predicate: "echoRequest",
				}).
				AddModifier(func(envelope *ronykit.Envelope) {
					envelope.SetHdr("X-Custom-Header", "justForTestingModifier")
				}).
				SetCoordinator(Forwarder).
				SetHandler(EchoHandler),
		).
		AddContract(
			desc.NewContract().
				SetInput(&dto.SumRequest{}).
				AddSelector(fasthttp.Selector{
					Method: fasthttp.MethodGet,
					Path:   "/sum/:val1/:val2",
				}).
				AddSelector(fasthttp.Selector{
					Method: fasthttp.MethodPost,
					Path:   "/sum",
				}).
				SetHandler(SumHandler),
		).
		AddContract(
			desc.NewContract().
				SetInput(&dto.SumRequest{}).
				AddSelector(fasthttp.Selector{
					Method: fasthttp.MethodGet,
					Path:   "/sum-redirect/:val1/:val2",
				}).
				AddSelector(fasthttp.Selector{
					Method: fasthttp.MethodPost,
					Path:   "/sum-redirect",
				}).
				SetHandler(SumRedirectHandler),
		).
		AddContract(
			desc.NewContract().
				SetInput(&dto.RedirectRequest{}).
				AddSelector(fasthttp.Selector{
					Method: fasthttp.MethodGet,
					Path:   "/redirect",
				}).
				SetHandler(Redirect),
		)
}

func Forwarder(ctx *ronykit.LimitedContext) (ronykit.ClusterMember, error) {
	// c := ctx.Cluster()

	// switch m := ctx.In().GetMsg().(type) {
	// case *dto.EchoRequest:
	// 	id := utils.Int64ToStr(m.RandomID % utils.StrToInt64(c.Me().ServerID()))
	// 	member, err := c.MemberByID(ctx.Context(), id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	return member, nil
	// }
	//
	// return c.Me(), nil

	return nil, nil
}

func EchoHandler(ctx *ronykit.Context) {
	req, ok := ctx.In().GetMsg().(*dto.EchoRequest)
	if !ok {
		ctx.Out().
			SetMsg(
				fasthttp.Err("E01", fmt.Sprintf("Request was not echoRequest: %s", reflect.TypeOf(ctx.In().GetMsg()))),
			).Send()

		return
	}

	ctx.Out().
		SetHdr("Content-Type", "application/json").
		SetMsg(
			&dto.EchoResponse{
				RandomID: req.RandomID,
				Ok:       req.Ok,
			},
		).Send()

	return
}

func SumHandler(ctx *ronykit.Context) {
	req, ok := ctx.In().GetMsg().(*dto.SumRequest)
	if !ok {
		ctx.Out().
			SetMsg(fasthttp.Err("E01", "Request was not echoRequest")).
			Send()

		return
	}

	ctx.Out().
		SetHdr("Content-Type", "application/json").
		SetMsg(
			&dto.SumResponse{
				Val: req.Val1 + req.Val2,
			},
		).Send()

	return
}

func SumRedirectHandler(ctx *ronykit.Context) {
	req, ok := ctx.In().GetMsg().(*dto.SumRequest)
	if !ok {
		ctx.Out().
			SetMsg(fasthttp.Err("E01", "Request was not echoRequest")).
			Send()

		return
	}

	rc, ok := ctx.Conn().(ronykit.RESTConn)
	if !ok {
		ctx.Out().
			SetMsg(fasthttp.Err("E01", "Only supports REST requests")).
			Send()

		return
	}

	switch rc.GetMethod() {
	case fasthttp.MethodGet:
		rc.Redirect(
			http.StatusTemporaryRedirect,
			fmt.Sprintf("http://%s/sum/%d/%d", rc.GetHost(), req.Val1, req.Val2),
		)
	case fasthttp.MethodPost:
		rc.Redirect(
			http.StatusTemporaryRedirect,
			fmt.Sprintf("http://%s/sum", rc.GetHost()),
		)
	default:
		ctx.Out().
			SetMsg(fasthttp.Err("E01", "Unsupported method")).
			Send()

		return
	}

	return
}

func Redirect(ctx *ronykit.Context) {
	req := ctx.In().GetMsg().(*dto.RedirectRequest) //nolint:forcetypeassert

	rc := ctx.Conn().(ronykit.RESTConn) //nolint:forcetypeassert
	rc.Redirect(http.StatusTemporaryRedirect, req.URL)
}
