package grpc

import (
	"fmt"
	"github.com/ducketlab/auth/pkg/token"
	"go.mongodb.org/mongo-driver/bson"
)

func newDescribeTokenRequest(req *token.DescribeTokenRequest) *describeTokenRequest {
	return &describeTokenRequest{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}
}

type describeTokenRequest struct {
	AccessToken  string
	RefreshToken string
}

func (req *describeTokenRequest) String() string {
	return fmt.Sprintf("access_token: %s, refresh_token: %s",
		req.AccessToken, req.RefreshToken)
}

func (req *describeTokenRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.AccessToken != "" {
		filter["_id"] = req.AccessToken
	}
	if req.RefreshToken != "" {
		filter["refresh_token"] = req.RefreshToken
	}

	return filter
}


