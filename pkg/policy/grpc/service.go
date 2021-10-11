package grpc

import (
	"context"
	"fmt"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"github.com/ducketlab/mingo/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	Service = &service{}
)

type service struct {
	log           logger.Logger
	col           *mongo.Collection

	namespace namespace.NamespaceServiceServer
	user      user.UserServiceServer
	role      role.RoleServiceServer

	policy.UnimplementedPolicyServiceServer
}

func (s *service) Config() error {
	if pkg.Namespace == nil {
		return fmt.Errorf("dependence namespace service is nil, please load first")
	}
	s.namespace = pkg.Namespace

	if pkg.User == nil {
		return fmt.Errorf("dependence user service is nil, please load first")
	}
	s.user = pkg.User

	if pkg.Role == nil {
		return fmt.Errorf("dependence role service is nil, please load first")
	}
	s.role = pkg.Role

	db := config.C().Mongo.GetDB()
	col := db.Collection("policy")

	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := col.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		return err
	}

	s.col = col
	s.log = zap.L().Named("policy")
	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return policy.HttpEntry()
}

func init() {
	pkg.RegisterService("policy", Service)
}