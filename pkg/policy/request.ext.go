package policy

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	MaxUserPolicy = 2048
)

var (
	validate = validator.New()
)

func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func NewDescribePolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{}
}

func (req *DescribePolicyRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}

func NewDeletePolicyRequestWithId(id string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.Id = id
	return req
}

func NewDeletePolicyRequestWithNamespaceId(namespaceID string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.NamespaceId = namespaceID
	return req
}

func NewDeletePolicyRequestWithAccount(account string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.Account = account
	return req
}

func NewDeletePolicyRequestWithRoleId(roleID string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.RoleId = roleID
	return req
}

func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{}
}

func (req *DeletePolicyRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}
