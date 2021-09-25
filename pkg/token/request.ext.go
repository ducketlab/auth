package token

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func NewDescribeTokenRequest() *DescribeTokenRequest {
	return &DescribeTokenRequest{}
}

func NewDescribeTokenRequestWithAccessToken(at string) *DescribeTokenRequest {
	req := NewDescribeTokenRequest()
	req.AccessToken = at
	return req
}

func (m *DescribeTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	if m.AccessToken == "" && m.RefreshToken == "" {
		return errors.New(
			"describe token request validate error, access_token and refresh_token required one")
	}

	return nil
}