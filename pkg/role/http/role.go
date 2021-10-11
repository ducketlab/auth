package http

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/http/context"
	"github.com/ducketlab/mingo/http/request"
	"github.com/ducketlab/mingo/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func (h *handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := role.NewCreateRoleRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	d, err := h.service.CreateRole(
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

func (h *handler) DescribeRole(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	pid := rctx.Ps.ByName("id")
	qs := r.URL.Query()

	req := role.NewDescribeRoleRequestWithId(pid)
	req.WithPermissions = qs.Get("with_permissions") == "true"

	var header, trailer metadata.MD
	ins, err := h.service.DescribeRole(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) ListRolePermission(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := role.NewQueryPermissionRequestFromHTTP(r)
	rctx := context.GetContext(r)
	req.RoleId = rctx.Ps.ByName("id")

	var header, trailer metadata.MD
	d, err := h.service.QueryPermission(
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

func (h *handler) AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	req := role.NewAddPermissionToRoleRequest()
	req.RoleId = rctx.Ps.ByName("id")

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	d, err := h.service.AddPermissionToRole(
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