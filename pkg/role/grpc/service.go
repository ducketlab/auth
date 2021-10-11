package grpc

import (
	"context"
	"fmt"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"github.com/ducketlab/mingo/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	Service = &service{}
)

type service struct {
	col           *mongo.Collection
	perm          *mongo.Collection

	policy policy.PolicyServiceServer

	log    logger.Logger
	role.UnimplementedRoleServiceServer
}

func (s *service) Config() error {
	if pkg.Policy == nil {
		return fmt.Errorf("dependence policy service is nil, please load first")
	}
	s.policy = pkg.Policy

	db := config.C().Mongo.GetDB()
	col := db.Collection("role")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "name", Value: bsonx.Int32(-1)},
				{Key: "domain", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	perm := db.Collection("permission")
	permIndexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = perm.Indexes().CreateMany(context.Background(), permIndexs)
	if err != nil {
		return err
	}

	s.col = col
	s.perm = perm
	s.log = zap.L().Named("role")
	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return role.HttpEntry()
}

func init() {
	pkg.RegisterService("role", Service)
}