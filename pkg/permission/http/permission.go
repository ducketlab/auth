package http

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/mingo/http/context"
	"github.com/ducketlab/mingo/http/request"
	"github.com/ducketlab/mingo/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func (h *handler) CheckPermission(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := permission.NewCheckPermissionRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.NamespaceId = rctx.Ps.ByName("id")

	var header, trailer metadata.MD
	d, err := h.service.CheckPermission(
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
