package grpc

import (
	"fmt"
	"github.com/ducketlab/auth/pkg/policy"
	"github.com/ducketlab/mingo/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func newDescribePolicyRequest(req *policy.DescribePolicyRequest) (*describePolicyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	return &describePolicyRequest{req}, nil
}

type describePolicyRequest struct {
	*policy.DescribePolicyRequest
}

func (req *describePolicyRequest) String() string {
	return fmt.Sprintf("policy: %s", req.Id)
}

func (req *describePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Id != "" {
		filter["_id"] = req.Id
	}
	return filter
}

func newDeletePolicyRequest(req *policy.DeletePolicyRequest) (*deletePolicyRequest, error) {
	return &deletePolicyRequest{
		DeletePolicyRequest: req,
	}, nil
}

type deletePolicyRequest struct {
	*policy.DeletePolicyRequest
}

func (r *deletePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	filter["domain"] = r.Domain

	if r.Id != "" {
		filter["_id"] = r.Id
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}
	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}
	if r.NamespaceId != "" {
		filter["namespace_id"] = r.NamespaceId
	}
	if r.Type != policy.PolicyType_NULL {
		filter["type"] = r.Type
	}

	return filter
}