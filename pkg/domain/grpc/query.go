package grpc

import (
	"github.com/ducketlab/auth/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func newQueryDomainRequest(req *domain.QueryDomainRequest) *queryDomainRequest {
	return &queryDomainRequest{
		QueryDomainRequest: req,
	}
}

type queryDomainRequest struct {
	*domain.QueryDomainRequest
}

func (r *queryDomainRequest) FindFilter() bson.M {
	filter := bson.M{}
	return filter
}

type describeDomainRequest struct {
	*domain.DescribeDomainRequest
}

func newDescribeDomainRequest(req *domain.DescribeDomainRequest) (*describeDomainRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &describeDomainRequest{req}, nil
}

func (req *describeDomainRequest) FindFilter() bson.M  {
	filter := bson.M{}

	if req.Name != "" {
		filter["_id"] = req.Name
	}

	return filter
}
