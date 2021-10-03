package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (req *CreateAccountRequest) Validate() error {
	return validate.Struct(req)
}

func NewDescribeAccountRequest() *DescribeAccountRequest {
	return &DescribeAccountRequest{}
}

func NewDescribeAccountRequestWithAccount(account string) *DescribeAccountRequest {
	return &DescribeAccountRequest{Account: account}
}

func (req *DescribeAccountRequest) Validate() error {
	if req.Account == "" {
		return errors.New("id or account is required")
	}

	return nil
}

func (req *QueryAccountRequest) Validate() error {
	return nil
}

func NewQueryAccountRequest() *QueryAccountRequest {
	return &QueryAccountRequest{
		SkipItems:      false,
	}
}

func NewCreateUserRequest() *CreateAccountRequest {
	return &CreateAccountRequest{
		Profile:     NewProfile(),
		ExpiresDays: DefaultExiresDays,
	}
}