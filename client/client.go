package client

import (
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"google.golang.org/grpc"
)

var (
	client *Client
)

type Client struct {
	config *Config
	connection *grpc.ClientConn
	logger logger.Logger
}

func C() *Client  {
	return client
}

func SetGlobal(cli *Client) {
	client = cli
}

func NewClient(conf *Config) (*Client, error) {
	err := zap.DevelopmentSetup()

	if err != nil {
		return nil, err
	}

	log := zap.L()

	conn, err := grpc.Dial(conf.address, grpc.WithInsecure(), grpc.WithPerRPCCredentials(conf.Authentication))
	if err != nil {
		return nil, err
	}

	return &Client{
		config: conf,
		connection: conn,
		logger:  log,
	}, nil
}

func (c *Client) GetClientId() string {
	return c.config.clientId
}

func (c *Client) Domain() domain.DomainServiceClient {
	return domain.NewDomainServiceClient(c.connection)
}

func (c *Client) Token() token.TokenServiceClient {
	return token.NewTokenServiceClient(c.connection)
}

func (c *Client) Namespace() namespace.NamespaceServiceClient {
	return namespace.NewNamespaceServiceClient(c.connection)
}

func (c *Client) User() user.UserServiceClient {
	return user.NewUserServiceClient(c.connection)
}

func (c *Client) Micro() micro.MicroServiceClient {
	return micro.NewMicroServiceClient(c.connection)
}

func (c *Client) Endpoint() endpoint.EndpointServiceClient {
	return endpoint.NewEndpointServiceClient(c.connection)
}

func (c *Client) Policy() policy.PolicyServiceClient {
	return policy.NewPolicyServiceClient(c.connection)
}

func (c *Client) Role() role.RoleServiceClient {
	return role.NewRoleServiceClient(c.connection)
}

func (c *Client) Permission() permission.PermissionServiceClient {
	return permission.NewPermissionServiceClient(c.connection)
}