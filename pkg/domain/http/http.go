package http

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {

}

func (h *handler) Registry(router router.SubRouter) {
	domainRouter := router.ResourceRouter("domain")
	domainRouter.BasePath("domains")
	domainRouter.Handle("GET", "/", h.ListDomains)
}

func (h *handler) Config() error {
	return nil
}
func init() {
	pkg.RegistryHttp("domain", api)
}



