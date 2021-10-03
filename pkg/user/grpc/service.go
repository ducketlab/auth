package grpc

import (
	"context"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
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
	user.UnimplementedUserServiceServer
}

func (s *service) Config() error {
	db := config.C().Mongo.GetDB()
	uc := db.Collection("user")

	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "department_id", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := uc.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		return err
	}

	s.col = uc
	s.log = zap.L().Named("user")
	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return user.HttpEntry()
}

func init() {
	pkg.RegisterService("user", Service)
}