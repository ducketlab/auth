package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service permission.PermissionServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("permission")
	r.BasePath("namespaces")
	r.Handle("POST", "/:id/permissions", h.CheckPermission)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Permission()
	return nil
}

func init() {
	pkg.RegistryV1Http("permission", api)
}
