package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/types/ftime"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateService(ctx context.Context, req *micro.CreateMicroRequest) (
	*micro.Micro, error) {

	ins, err := micro.New(req)
	if err != nil {
		return nil, err
	}

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins.Creater = tk.Account
	ins.Domain = tk.Domain

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
	}
	return ins, nil
}

func (s *service) DescribeService(ctx context.Context, req *micro.DescribeMicroRequest) (
	*micro.Micro, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := new(micro.Micro)
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find service %s error, %s", req, err)
	}
	return ins, nil
}

func (s *service) ValidateClientCredential(ctx context.Context, req *micro.ValidateClientCredentialRequest) (
	*micro.Micro, error) {
	descReq := micro.NewDescribeServiceRequestWithClientId(req.ClientId)
	ins, err := s.DescribeService(ctx, descReq)
	if err != nil {
		return nil, err
	}
	if err := ins.ValidateClientCredential(req.ClientSecret); err != nil {
		return nil, err
	}
	ins.Desensitize()
	return ins, nil
}

func (s *service) RefreshServiceClientSecret(ctx context.Context, req *micro.DescribeMicroRequest) (
	*micro.Micro, error) {

	ins, err := s.DescribeService(ctx, req)
	if err != nil {
		return nil, err
	}
	if !ins.ClientEnabled {
		return nil, exception.NewBadRequest("client is not enabled")
	}

	ins.ClientSecret = token.MakeBearer(32)
	ins.ClientRefreshAt = ftime.Now().Timestamp()
	if err := s.update(ins); err != nil {
		return nil, err
	}

	return ins, nil
}