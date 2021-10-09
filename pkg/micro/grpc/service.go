package grpc

import (
	"context"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/micro"
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
	log      logger.Logger

	micro.UnimplementedMicroServiceServer
}


func (s *service) Config() error {
	db := config.C().Mongo.GetDB()
	sc := db.Collection("micro")

	indexes := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "name", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "client_id", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := sc.Indexes().CreateMany(context.Background(), indexes)

	if err != nil {
		return err
	}
	s.col = sc

	s.log = zap.L().Named("micro")

	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return micro.HttpEntry()
}

func init() {
	pkg.RegisterService("micro", Service)
}