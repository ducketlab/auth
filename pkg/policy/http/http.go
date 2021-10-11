package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service policy.PolicyServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("policy")

	r.BasePath("policies")
	r.Handle("POST", "/", h.Create)
	r.Handle("GET", "/:id", h.Get)
	r.Handle("DELETE", "/:id", h.Delete)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Policy()
	return nil
}

func init() {
	pkg.RegistryV1Http("policy", api)
}
