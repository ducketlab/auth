package grpc

import (
	"context"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/namespace"
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
	col *mongo.Collection
	log logger.Logger
	namespace.UnimplementedNamespaceServiceServer
}

func (s *service) Config() error {
	db := config.C().Mongo.GetDB()
	ac := db.Collection("namespace")

	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "domain", Value: bsonx.Int32(-1)},
				{Key: "name", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := ac.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		return err
	}

	s.col = ac
	s.log = zap.L().Named("namespace")

	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return namespace.HttpEntry()
}

func init() {
	pkg.RegisterService("namespace", Service)
}