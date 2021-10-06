package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service namespace.NamespaceServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("namespace")
	r.BasePath("namespaces")
	r.Handle("POST", "/", h.Create)
	r.Handle("GET", "/:id", h.Get)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Namespace()
	return nil
}

func init() {
	pkg.RegistryV1Http("namespace", api)
}
