package http

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/http/context"
	"github.com/ducketlab/mingo/http/request"
	"github.com/ducketlab/mingo/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	cid, cs := pkg.GetClientCredentialsFromHttpRequest(r)
	if cid == "" && cs == "" {
		response.Failed(w, exception.NewBadRequest("service client credentials in header missed"))
		return
	}

	ctx, err := pkg.NewGrpcInCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ctx.SetClientCredentials(cid, cs)
	req := endpoint.NewDefaultRegistryRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	_, err = pkg.Endpoint.Registry(
		ctx.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
}


func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := context.GetContext(r)

	id := rctx.Ps.ByName("id")
	req := endpoint.NewDescribeEndpointRequestWithId(id)

	var header, trailer metadata.MD
	d, err := h.endpoint.DescribeEndpoint(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, d)
	return
}