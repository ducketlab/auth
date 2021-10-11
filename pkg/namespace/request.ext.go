package namespace

import (
	"errors"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/go-playground/validator/v10"
)

var (
	validater = validator.New()
)

func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{}
}

func (req *CreateNamespaceRequest) UpdateOwner(tk *token.Token) {
	req.CreateBy = tk.Account
	req.Domain = tk.Domain
}

func NewDescribeNamespaceRequest() *DescribeNamespaceRequest {
	return &DescribeNamespaceRequest{}
}

func (req *CreateNamespaceRequest) Validate() error {
	return validater.Struct(req)
}

func (req *DescribeNamespaceRequest) Validate() error {
	if req.Id == "" {
		return errors.New("id  is required")
	}

	return nil
}

func NewNewDescribeNamespaceRequestWithId(id string) *DescribeNamespaceRequest {
	req := NewDescribeNamespaceRequest()
	req.Id = id
	return req
}
