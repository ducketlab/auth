package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/exception"
)

func (s *service) CheckPermission(ctx context.Context, req *permission.CheckPermissionRequest) (
	*role.Permission, error) {

	if req.EndpointId == "" {
		req.EndpointId = endpoint.GenHashId(req.ServiceId, req.Path)
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	roleReq := permission.NewQueryRoleRequest(req.NamespaceId)
	roleReq.WithPermission = true
	roleSet, err := s.QueryRole(ctx, roleReq)
	if err != nil {
		return nil, err
	}

	ep, err := s.endpoint.DescribeEndpoint(ctx, endpoint.NewDescribeEndpointRequestWithId(req.EndpointId))
	if err != nil {
		return nil, err
	}
	s.log.Debugf("check roles %s has permission access endpoint [%s]", roleSet.RoleNames(), ep.Entry)

	if !ep.Entry.PermissionEnable {
		return role.NewDefaulPermission(), nil
	}

	p, ok, err := roleSet.HasPermission(ep)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, exception.NewPermissionDeny(
			"in namespace %s, role %s has no permission access endpoint: %s",
			req.NamespaceId,
			roleSet.RoleNames(),
			ep.Entry.Path,
		)
	}

	return p, nil
}
