package permission

import "fmt"

func (req *CheckPermissionRequest) Validate() error {
	if req.NamespaceId == "" {
		return fmt.Errorf("namespace required")
	}

	if req.EndpointId == "" && (req.ServiceId == "" || req.Path == "") {
		return fmt.Errorf("endpoint_id or (service_id and path) required when check")
	}

	return nil
}

func NewQueryRoleRequest(namespaceId string) *QueryRoleRequest {
	return &QueryRoleRequest{
		NamespaceId: namespaceId,
	}
}

func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{
	}
}