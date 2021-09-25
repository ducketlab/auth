package grpc

import (
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/pb/http"
)

var (
	Service = service{}
)

type service struct {
	token.UnimplementedTokenServiceServer
}

func (s service) Config() error {
	return nil
}

func (s service) HttpEntry() *http.EntrySet {
	return nil
}

func init()  {
	pkg.RegisterService("token", Service)
}


