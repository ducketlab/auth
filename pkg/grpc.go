package pkg

import (
	"fmt"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/pb/http"
	"google.golang.org/grpc"
)

var (
	// Token service
	Token token.TokenServiceServer
	// Domain service
	Domain domain.DomainServiceServer
	// User service
	User user.UserServiceServer
	// Namespace service
	Namespace namespace.NamespaceServiceServer
	// Micro service
	Micro micro.MicroServiceServer
	// Endpoint service
	Endpoint endpoint.EndpointServiceServer
	// Role service
	Role role.RoleServiceServer
	// Policy service
	Policy policy.PolicyServiceServer
	// Permission service
	Permission permission.PermissionServiceServer
)

var (
	services      []Service
	successLoaded []string
	entrySet      = http.NewEntrySet()
)

type Service interface {
	Config() error
	HttpEntry() *http.EntrySet
}

func HttpEntry() *http.EntrySet {
	return entrySet
}

func GetPathEntry(path string) *http.Entry {
	es := HttpEntry()
	for i := range es.Items {
		if es.Items[i].Path == path {
			return es.Items[i]
		}
	}

	return nil
}

func LoadedService() []string {
	return successLoaded
}

func addService(name string, service Service) {
	services = append(services, service)
	successLoaded = append(successLoaded, name)
}

func InitGrpcApi(server *grpc.Server)  {
	domain.RegisterDomainServiceServer(server, Domain)
	user.RegisterUserServiceServer(server, User)
	token.RegisterTokenServiceServer(server, Token)
	namespace.RegisterNamespaceServiceServer(server, Namespace)
	micro.RegisterMicroServiceServer(server, Micro)
	endpoint.RegisterEndpointServiceServer(server, Endpoint)
	role.RegisterRoleServiceServer(server, Role)
	policy.RegisterPolicyServiceServer(server, Policy)
	permission.RegisterPermissionServiceServer(server, Permission)
}

func RegisterService(name string, svr Service) {
	switch value := svr.(type) {
	case domain.DomainServiceServer:
		if Domain != nil {
			registryError(name)
		}
		Domain = value
		addService(name, svr)
	case user.UserServiceServer:
		if User != nil {
			registryError(name)
		}
		User = value
		addService(name, svr)
	case token.TokenServiceServer:
		if Token != nil {
			registryError(name)
		}
		Token = value
		addService(name, svr)
	case namespace.NamespaceServiceServer:
		if Namespace != nil {
			registryError(name)
		}
		Namespace = value
		addService(name, svr)
	case micro.MicroServiceServer:
		if Micro != nil {
			registryError(name)
		}
		Micro = value
		addService(name, svr)
	case endpoint.EndpointServiceServer:
		if Endpoint != nil {
			registryError(name)
		}
		Endpoint = value
		addService(name, svr)
	case role.RoleServiceServer:
		if Role != nil {
			registryError(name)
		}
		Role = value
		addService(name, svr)
	case policy.PolicyServiceServer:
		if Policy != nil {
			registryError(name)
		}
		Policy = value
		addService(name, svr)
	case permission.PermissionServiceServer:
		if Permission != nil {
			registryError(name)
		}
		Permission = value
		addService(name, svr)
	default:
		panic(fmt.Sprintf("unknown service type %s", name))
	}
}

func registryError(name string) {
	panic("service " + name + "has registered")
}

func InitService() error {
	for _, s := range services {
		if err := s.Config(); err != nil {
			return err
		}
		entrySet.Merge(s.HttpEntry())
	}
	return nil
}
