package grpc

import (
	"fmt"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func newDescribeEndpointRequest(req *endpoint.DescribeEndpointRequest) (*describeEndpointRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeEndpointRequest{req}, nil
}

type describeEndpointRequest struct {
	*endpoint.DescribeEndpointRequest
}

func (req *describeEndpointRequest) String() string {
	return fmt.Sprintf("endpoint: %s", req.Id)
}

func (req *describeEndpointRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Id != "" {
		filter["_id"] = req.Id
	}

	return filter
}


