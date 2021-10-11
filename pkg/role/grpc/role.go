package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.Role, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	r, err := role.New(tk, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), r); err != nil {
		return nil, exception.NewInternalServerError("inserted role(%s) document error, %s",
			r.Name, err)
	}

	permReq := role.NewAddPermissionToRoleRequest()
	permReq.Permissions = req.Permissions
	permReq.RoleId = r.Id
	ps, err := s.AddPermissionToRole(ctx, permReq)
	if err != nil {
		return nil, err
	}
	r.Permissions = ps.Items
	return r, nil
}

func (s *service) DescribeRole(ctx context.Context, req *role.DescribeRoleRequest) (*role.Role, error) {
	query, err := newDescribeRoleRequest(req)
	if err != nil {
		return nil, err
	}

	ins := role.NewDefaultRole()
	if err := s.col.FindOne(context.TODO(), query.FindFilter(), query.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("role %s not found", req)
		}

		return nil, exception.NewInternalServerError("find role %s error, %s", req, err)
	}

	if req.WithPermissions {
		queryPerm := role.NewQueryPermissionRequest()
		queryPerm.RoleId = ins.Id
		ps, err := s.QueryPermission(ctx, queryPerm)
		if err != nil {
			return nil, err
		}
		ins.Permissions = ps.Items
	}

	return ins, nil
}

