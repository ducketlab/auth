package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	endpoint endpoint.EndpointServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("endpoint")

	r.BasePath("endpoints")
	r.Handle("POST", "/", h.Create)
	r.Handle("GET", "/:id", h.Get)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.endpoint = client.Endpoint()
	return nil
}

func init() {
	pkg.RegistryV1Http("endpoint", api)
}