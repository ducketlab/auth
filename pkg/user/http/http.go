package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service user.UserServiceClient
	domain  domain.DomainServiceClient
}

func (h *handler) Registry(router router.SubRouter) {

}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.User()
	h.domain = client.Domain()
	return nil
}

func init() {
	pkg.RegistryV1Http("user", api)
}