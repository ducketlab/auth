package client

import (
	"github.com/ducketlab/mingo/grpc/gcontext"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	httpb "github.com/ducketlab/mingo/pb/http"
	"github.com/rs/xid"
	"net/http"
)

func NewHttpAuther(c *Client) *HttpAuther {
	return &HttpAuther{
		client: c,
		logger: zap.L().Named("http-interceptor"),
	}
}

type HttpAuther struct {
	logger logger.Logger
	client *Client
}

func (a *HttpAuther) Auth(r *http.Request, entry httpb.Entry) (
	authInfo interface{}, err error) {

	ctx, err := gcontext.NewGrpcInCtxFromHttpRequest(r)
	if err != nil {
		return nil, err
	}

	engine := newEntryEngine(a.client, &entry, a.log())
	engine.UseUniPath()

	tk, err := engine.ValidateIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if err := engine.ValidatePermission(tk, ctx); err != nil {
		return nil, err
	}

	if r.Header.Get(gcontext.RequestId) == "" {
		r.Header.Set(gcontext.RequestId, xid.New().String())
	}

	return tk, nil
}

func (a *HttpAuther) log() logger.Logger {
	if a == nil {
		a.logger = zap.L().Named("http-auther")
	}

	return a.logger
}

func (a *HttpAuther) SetLogger(l logger.Logger) {
	a.logger = l
}