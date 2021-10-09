package http

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/mingo/http/request"
	"github.com/ducketlab/mingo/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHttpRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := micro.NewCreateMicroRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	d, err := h.service.CreateService(
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
