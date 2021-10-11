package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/exception"
)

func insertDocs(perms []*role.Permission) []interface{} {
	var docs []interface{}
	for i := range perms {
		docs = append(docs, perms[i])
	}
	return docs
}

func (s *service) AddPermissionToRole(ctx context.Context, req *role.AddPermissionToRoleRequest) (*role.PermissionSet, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add permission error, %s", err)
	}

	ins, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithId(req.RoleId))
	if err != nil {
		return nil, err
	}

	perms := role.NewPermission(ins.Id, tk.Account, req.Permissions)
	if _, err := s.perm.InsertMany(context.TODO(), insertDocs(perms)); err != nil {
		return nil, exception.NewInternalServerError("inserted permission(%s) document error, %s",
			perms, err)
	}

	set := role.NewPermissionSet()
	set.Items = perms
	return set, nil
}

func (s *service) QueryPermission(ctx context.Context, req *role.QueryPermissionRequest) (
	*role.PermissionSet, error) {

	query, err := newQueryPermissionRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query permission filter: %s", query.FindFilter())
	resp, err := s.perm.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find permission error, error is %s", err)
	}

	set := role.NewPermissionSet()
	if !req.SkipItmes {
		for resp.Next(context.TODO()) {
			ins := role.NewDefaultPermission()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode permission error, error is %s", err)
			}
			set.Add(ins)
		}
	}

	count, err := s.perm.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get permission count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}
