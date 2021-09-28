package grpc

import (
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/mingo/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Service = &service{}
)

type service struct {
	col *mongo.Collection
	domain.UnimplementedDomainServiceServer
}

func (s *service) Config() error {
	db := config.C().Mongo.GetDB()

	dc := db.Collection("domain")

	s.col = dc

	return nil
}

func (s *service) HttpEntry() *http.EntrySet {
	return domain.HttpEntry()
}

func init() {
	pkg.RegisterService("domain", Service)
}
