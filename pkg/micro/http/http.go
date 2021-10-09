package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service micro.MicroServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("service")
	r.BasePath("services")

	r.Handle("POST", "/", h.CreateService)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Micro()
	return nil
}

func init() {
	pkg.RegistryV1Http("service", api)
}
