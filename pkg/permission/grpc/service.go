package grpc

import (
	"errors"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"github.com/ducketlab/mingo/pb/http"
)

var (
	Service = &service{}
)

type service struct {
	log      logger.Logger
	policy   policy.PolicyServiceServer
	role     role.RoleServiceServer
	endpoint endpoint.EndpointServiceServer

	permission.UnimplementedPermissionServiceServer
}

func (s *service) Config() error {
	if pkg.Policy == nil {
		return errors.New("dependence policy service is nil")
	}
	s.policy = pkg.Policy

	if pkg.Role == nil {
		return errors.New("dependence role service is nil")
	}
	s.role = pkg.Role

	if pkg.Endpoint == nil {
		return errors.New("dependence endpoint service is nil")
	}
	s.endpoint = pkg.Endpoint
	s.log = zap.L().Named("permission")
	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return permission.HttpEntry()
}

func init() {
	pkg.RegisterService("permission", Service)
}