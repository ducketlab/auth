package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service role.RoleServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("role")

	r.BasePath("roles")
	r.Handle("POST", "/", h.CreateRole)
	r.Handle("GET", "/:id", h.DescribeRole)
	r.Handle("GET", "/:id/permissions", h.ListRolePermission)
	r.Handle("POST", "/:id/permissions", h.AddPermissionToRole)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Role()
	return nil
}

func init() {
	pkg.RegistryV1Http("role", api)
}
