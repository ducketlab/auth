package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/version"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateDomain(_ context.Context,
	req *domain.CreateDomainRequest) (*domain.Domain, error) {

	d, err := domain.New(version.ServiceName, req)

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := s.col.InsertOne(context.TODO(), d); err != nil {
		return nil, exception.NewInternalServerError("inserted a domain document error, %s", err)
	}

	return d, nil
}

func (s *service) DescribeDomain(_ context.Context,
	req *domain.DescribeDomainRequest) (*domain.Domain, error) {

	r, err := newDescribeDomainRequest(req)

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	d := domain.NewDefault()

	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
				return nil, exception.NewNotFound("domain %s not found", req)
		}

		return nil, exception.NewInternalServerError("find domain %s error, %s", req.Name, err)
	}

	return d, nil
}

func (s *service) QueryDomain(_ context.Context, req *domain.QueryDomainRequest) (*domain.Set, error) {
	r := newQueryDomainRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter())

	if err != nil {
		return nil, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	domainSet := domain.NewDomainSet()

	for resp.Next(context.TODO()) {
		d := new(domain.Domain)
		if err := resp.Decode(d); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		domainSet.Add(d)
	}

	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	domainSet.Total = count

	return domainSet, nil
}

func (s *service) DeleteDomain(ctx context.Context,
	req *domain.DeleteDomainRequest) (*domain.Domain, error) {

	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.Name})

	if err != nil {
		return nil, exception.NewInternalServerError("delete domain(%s) error, %s", req.Name, err)
	}

	if result.DeletedCount == 0 {
		return nil, exception.NewNotFound("domain %s not found", req.Name)
	}

	return nil, nil
}