package grpc

import (
	"github.com/ducketlab/auth/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func newQueryDomainRequest(req *domain.QueryDomainRequest) *request {
	return &request{
		QueryDomainRequest: req,
	}
}

type request struct {
	*domain.QueryDomainRequest
}

func (r *request) FindFilter() bson.M {
	filter := bson.M{}
	return filter
}
