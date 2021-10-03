package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateAccount(ctx context.Context,
	req *user.CreateAccountRequest) (*user.User, error) {

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)

	if err != nil {
		return nil, err
	}

	u, err := user.New(req)

	if tk != nil {
		u.Domain = tk.Domain
	}

	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) DescribeAccount(ctx context.Context,
	req *user.DescribeAccountRequest) (*user.User, error) {

	r, err := newDescribeRequest(req)
	if err != nil {
		return nil, err
	}

	ins := user.NewDefaultUser()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req, err)
	}

	return ins, nil
}

func (s *service) QueryAccount(ctx context.Context, req *user.QueryAccountRequest) (*user.Set, error) {
	r, err := newQueryUserRequest(req)
	if err != nil {
		return nil, err
	}
	return s.queryAccount(ctx, r)
}