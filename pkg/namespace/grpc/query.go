package grpc

import (
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/types/ftime"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func newNamespace(req *namespace.CreateNamespaceRequest) (*namespace.Namespace, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &namespace.Namespace{
		Id:          xid.New().String(),
		Domain:      req.Domain,
		CreateBy:    req.CreateBy,
		CreateAt:    ftime.Now().Timestamp(),
		UpdateAt:    ftime.Now().Timestamp(),
		Name:        req.Name,
		Picture:     req.Picture,
		Owner:       req.Owner,
		Description: req.Description,
	}

	return ins, nil
}

func newDescribeQuery(req *namespace.DescribeNamespaceRequest) (*describeNamespaceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeNamespaceRequest{req}, nil
}

type describeNamespaceRequest struct {
	*namespace.DescribeNamespaceRequest
}

func (r *describeNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Id != "" {
		filter["_id"] = r.Id
	}

	return filter
}
