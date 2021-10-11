package role

import (
	"fmt"
	"github.com/ducketlab/mingo/exception"
	"net/http"
)

func NewDescribeRoleRequestWithId(id string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Id:              id,
		WithPermissions: false,
	}
}

func NewAddPermissionToRoleRequest() *AddPermissionToRoleRequest {
	return &AddPermissionToRoleRequest{
		Permissions: []*CreatePermissionRequest{},
	}
}

func (req *AddPermissionToRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DescribeRoleRequest) Validate() error {
	if req.Id == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}

func (req *DescribePermissionRequest) Validate() error {
	if req.Id == "" {
		return exception.NewBadRequest("id required")
	}
	return nil
}

func (req *QueryPermissionRequest) Validate() error {
	return validate.Struct(req)
}

func NewQueryPermissionRequest() *QueryPermissionRequest {
	return &QueryPermissionRequest{}
}


func NewQueryPermissionRequestFromHTTP(r *http.Request) *QueryPermissionRequest {
	req := NewQueryPermissionRequest()
	return req
}