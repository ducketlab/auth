package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreatePolicy(ctx context.Context, req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins, err := policy.New(tk, req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u, err := ins.CheckDependence(ctx, s.user, s.role, s.namespace)
	if err != nil {
		return nil, err
	}
	ins.UserType = u.Type

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.Id, err)
	}

	return ins, nil
}

func (s *service) DescribePolicy(ctx context.Context, req *policy.DescribePolicyRequest) (
	*policy.Policy, error) {
	r, err := newDescribePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	ins := policy.NewDefaultPolicy()
	s.log.Debugf("describe policy filter: %s", r.FindFilter())
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *service) DeletePolicy(ctx context.Context, req *policy.DeletePolicyRequest) (*policy.Policy, error) {
	descReq := policy.NewDescribePolicyRequest()
	descReq.Id = req.Id
	p, err := s.DescribePolicy(ctx, descReq)
	if err != nil {
		return nil, err
	}

	r, err := newDeletePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	_, err = s.col.DeleteOne(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete policy(%s) error, %s", req.Id, err)
	}

	return p, nil
}