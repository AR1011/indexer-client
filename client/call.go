package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AR1011/indexer-client/types"
	"github.com/valyala/fasthttp"
)

type Caller interface {
	Call(context.Context, string, *types.JRPCRequest, interface{}) error
	WithDebug()
}

type FastHTTPCaller struct {
	Client *fasthttp.Client
	debug  bool
}

var DefaultFastHTTPCaller = &FastHTTPCaller{
	Client: &fasthttp.Client{},
	debug:  true,
}

func (a *FastHTTPCaller) WithDebug() {
	a.debug = true
}

func (a FastHTTPCaller) Call(ctx context.Context, url string, jreq *types.JRPCRequest, jres interface{}) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(fasthttp.MethodPost)

	body, err := json.Marshal(jreq)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	req.SetBodyRaw(body)
	if a.debug {
		fmt.Println(string(body))
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetConnectionClose()
	if err := a.Client.DoTimeout(req, resp, time.Second*15); err != nil {
		if err == fasthttp.ErrTimeout {
			return context.DeadlineExceeded
		}
		return fmt.Errorf("fasthttp do request: %w", err)
	}

	if statusCode := resp.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		return fmt.Errorf("internal server error: %d", statusCode)
	}

	if a.debug {
		fmt.Println(string(resp.Body()))
	}

	err = json.Unmarshal(resp.Body(), jres)
	if err != nil {
		return fmt.Errorf("decode json: %w", err)
	}

	return nil
}

var _ Caller = (*FastHTTPCaller)(nil)
