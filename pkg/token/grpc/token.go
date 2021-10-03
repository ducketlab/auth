package grpc

import (
	"context"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) DescribeToken(ctx context.Context,
	req *token.DescribeTokenRequest) (*token.Token, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.describeToken(newDescribeTokenRequest(req))
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	return tk, nil
}

func (s *service) IssueToken(ctx context.Context,
	req *token.IssueTokenRequest) (*token.Token, error) {

	tk, err := s.issuer.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}

	tk.WithRemoteIp(req.GetRemoteIp())
	tk.WithUserAgent(req.GetUserAgent())

	if err := s.saveToken(tk); err != nil {
		return nil, err
	}

	return tk, nil
}

func (s *service) describeToken(req *describeTokenRequest) (*token.Token, error) {
	tk := new(token.Token)

	if err := s.col.FindOne(context.TODO(), req.FindFilter()).Decode(tk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", req)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", req, err)
	}

	return tk, nil
}
