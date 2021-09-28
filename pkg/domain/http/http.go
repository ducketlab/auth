package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service domain.DomainServiceClient
}

func (h *handler) Registry(router router.SubRouter) {
	domainRouter := router.ResourceRouter("domain")
	domainRouter.BasePath("domains")
	domainRouter.Handle("POST", "/", h.CreateDomain)
	domainRouter.Handle("GET", "/", h.ListDomains)
}

func (h *handler) Config() error {

	cli := client.C()
	if cli == nil {
		return errors.New("grpc client not initial")
	}

	h.service = cli.Domain()
	return nil
}

func init() {
	pkg.RegistryV1Http("domain", api)
}



