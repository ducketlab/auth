package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/exception"
)

func (s *service) DescribeToken(ctx context.Context,
	req *token.DescribeTokenRequest) (*token.Token, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return nil, nil
}
