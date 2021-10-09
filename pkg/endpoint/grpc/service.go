package grpc

import (
	"context"
	"fmt"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/micro"
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
	col           *mongo.Collection
	micro         micro.MicroServiceServer
	log           logger.Logger

	endpoint.UnimplementedEndpointServiceServer
}

func (s *service) Config() error {
	db := config.C().Mongo.GetDB()
	col := db.Collection("endpoint")

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

	if pkg.Micro == nil {
		return fmt.Errorf("dependence micro service is nil, please load first")
	}
	s.micro = pkg.Micro
	s.log = zap.L().Named("endpoint")
	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return endpoint.HttpEntry()
}

func init() {
	pkg.RegisterService("endpoint", Service)
}
