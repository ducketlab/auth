package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/mingo/exception"
)

func (s *service) CreateDomain(ctx context.Context,
	req *domain.CreateDomainRequest) (*domain.Domain, error) {

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)

	if err != nil {
		return nil, err
	}

	d, err := domain.New(tk.Account, req)

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return d, nil
}
