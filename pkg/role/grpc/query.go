package grpc

import (
	"fmt"
	"github.com/ducketlab/auth/pkg/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newDescribeRoleRequest(req *role.DescribeRoleRequest) (*describeRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describeRoleRequest{req}, nil
}

type describeRoleRequest struct {
	*role.DescribeRoleRequest
}

func (req *describeRoleRequest) String() string {
	return fmt.Sprintf("role: %s", req.Name)
}

func (req *describeRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

func (req *describeRoleRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}

func newDescribePermissionRequest(req *role.DescribePermissionRequest) (*describePermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describePermissionRequest{req}, nil
}

type describePermissionRequest struct {
	*role.DescribePermissionRequest
}

func (req *describePermissionRequest) String() string {
	return fmt.Sprintf("permission: %s", req.Id)
}

func (req *describePermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	return filter
}

func (req *describePermissionRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}

func newQueryPermissionRequest(req *role.QueryPermissionRequest) (*queryPermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryPermissionRequest{
		QueryPermissionRequest: req}, nil
}

type queryPermissionRequest struct {
	*role.QueryPermissionRequest
}

func (r *queryPermissionRequest) FindOptions() *options.FindOptions {
	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
	}

	return opt
}

func (r *queryPermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}

	return filter
}