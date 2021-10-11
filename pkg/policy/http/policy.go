package http

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/mingo/http/context"
	"github.com/ducketlab/mingo/http/request"
	"github.com/ducketlab/mingo/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := policy.NewCreatePolicyRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Type = policy.PolicyType_CUSTOM

	var header, trailer metadata.MD
	d, err := h.service.CreatePolicy(
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
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := policy.NewDescribePolicyRequest()
	req.Id = rctx.Ps.ByName("id")

	var header, trailer metadata.MD
	d, err := h.service.DescribePolicy(
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
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := policy.NewDeletePolicyRequestWithId(rctx.Ps.ByName("id"))
	req.Domain = tk.Domain

	var header, trailer metadata.MD
	_, err = h.service.DeletePolicy(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, "ok")
}