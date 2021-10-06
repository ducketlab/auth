package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateNamespace(ctx context.Context, req *namespace.CreateNamespaceRequest) (
	*namespace.Namespace, error) {

	ins, err := newNamespace(req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError(
			"inserted namespace(%s) document error, %s",
			ins.Name, err)
	}

	return ins, nil
}

func (s *service) DescribeNamespace(ctx context.Context,
	req *namespace.DescribeNamespaceRequest) (*namespace.Namespace, error) {

	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := namespace.NewDefaultNamespace()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("namespace %s not found", req)
		}

		return nil, exception.NewInternalServerError("find namespace %s error, %s", req.Id, err)
	}

	return ins, nil
}

