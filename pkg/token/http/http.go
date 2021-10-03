package http

import (
	"errors"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/http/router"
)

var (
	api = &handler{}
)

type handler struct {
	service token.TokenServiceClient
}

func (h *handler) Registry(router router.SubRouter) {

	r := router.ResourceRouter("token")
	r.BasePath("/oauth2/tokens")
	r.Handle("POST", "/", h.IssueToken)
}

func (h *handler) Config() error {
	c := client.C()
	if c == nil {
		return errors.New("grpc c not initial")
	}

	h.service = c.Token()
	return nil
}

func init()  {
	pkg.RegistryV1Http("token", api)
}