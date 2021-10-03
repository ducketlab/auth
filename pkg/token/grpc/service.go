package grpc

import (
	"context"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/token/issuer"
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
	col    *mongo.Collection
	log    logger.Logger
	issuer issuer.Issuer
	token.UnimplementedTokenServiceServer
}

func (s *service) Config() error {

	tokenIssuer, err := issuer.NewTokenIssuer()
	if err != nil {
		return err
	}

	s.issuer = tokenIssuer

	db := config.C().Mongo.GetDB()
	col := db.Collection("token")

	indexes := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "refresh_token", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = col.Indexes().CreateMany(context.Background(), indexes)

	if err != nil {
		return err
	}

	s.col = col
	s.log = zap.L().Named("token")

	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return token.HttpEntry()
}

func init() {
	pkg.RegisterService("token", Service)
}
