package grpc

import (
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func newDescribeQuery(req *micro.DescribeMicroRequest) (*describeRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeRequest{req}, nil
}

type describeRequest struct {
	*micro.DescribeMicroRequest
}

func (r *describeRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Name != "" {
		filter["name"] = r.Name
	}
	if r.Id != "" {
		filter["_id"] = r.Id
	}
	if r.ClientId != "" {
		filter["client_id"] = r.ClientId
	}

	return filter
}
