package role

import (
	"fmt"
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/types/ftime"
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
	"hash/fnv"
	"strings"
)

const (
	MaxPermission = 500
)

var (
	validate = validator.New()
)

func New(tk *token.Token, req *CreateRoleRequest) (*Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &Role{
		Id:          xid.New().String(),
		CreateAt:    ftime.Now().Timestamp(),
		UpdateAt:    ftime.Now().Timestamp(),
		Domain:      tk.Domain,
		Creater:     tk.Account,
		Type:        req.Type,
		Name:        req.Name,
		Meta:        req.Meta,
		Description: req.Description,
	}
	r.Permissions = NewPermission(r.Id, tk.Account, req.Permissions)
	return r, nil
}

func NewDefaultRole() *Role {
	return &Role{
		Permissions: []*Permission{},
		Type:        RoleType_CUSTOM,
		Meta:        map[string]string{},
	}
}

func (r *Role) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	var (
		rok, lok bool
	)
	for i := range r.Permissions {
		rok = r.Permissions[i].MatchResource(ep.ServiceId, ep.Entry.Resource)
		lok = r.Permissions[i].MatchLabel(ep.Entry.Labels)
		if rok && lok {
			return r.Permissions[i], true, nil
		}
	}
	return nil, false, nil
}

func NewDefaultPermission() *Permission {
	return &Permission{}
}

func NewPermission(roleID, creater string, perms []*CreatePermissionRequest) []*Permission {
	set := []*Permission{}
	for i := range perms {
		set = append(set, &Permission{
			Id:           PermissionHash(roleID, perms[i]),
			RoleId:       roleID,
			CreateAt:     ftime.Now().Timestamp(),
			Creater:      creater,
			Effect:       perms[i].Effect,
			ServiceId:    perms[i].ServiceId,
			ResourceName: perms[i].ResourceName,
			LabelKey:     perms[i].LabelKey,
			MatchAll:     perms[i].MatchAll,
			LabelValues:  perms[i].LabelValues,
		})
	}
	return set
}

func PermissionHash(namespace string, perm *CreatePermissionRequest) string {
	h := fnv.New32a()

	h.Write([]byte(namespace + perm.Effect.String() + perm.ServiceId + perm.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Permissions: []*CreatePermissionRequest{},
		Type:        RoleType_CUSTOM,
	}
}

func (req *CreateRoleRequest) Validate() error {
	pc := len(req.Permissions)
	if pc > MaxPermission {
		return fmt.Errorf("role permission overed max count: %d",
			MaxPermission)
	}

	errs := []string{}
	for i := range req.Permissions {
		if err := req.Permissions[i].Validate(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("validate permission error, %s", strings.Join(errs, ","))
	}

	return validate.Struct(req)
}

func (p *CreatePermissionRequest) Validate() error {
	if p.ServiceId == "" || p.ResourceName == "" || p.LabelKey == "" {
		return fmt.Errorf("permisson required service_id, resource_name and label_key")
	}

	if len(p.LabelValues) == 0 {
		return fmt.Errorf("permission label_values required")
	}

	return nil
}

func (p *Permission) ID(namespace string) string {
	return namespace + "." + p.ResourceName
}

func (p *Permission) MatchResource(serviceID, resourceName string) bool {
	if p.ServiceId != "*" && p.ServiceId != serviceID {
		return false
	}

	if p.ResourceName != "*" && p.ResourceName != resourceName {
		return false
	}

	return true
}

func (p *Permission) MatchLabel(label map[string]string) bool {
	for k, v := range label {
		if p.LabelKey == "*" || p.LabelKey == k {
			if p.isMatchAllValue() {
				return true
			}
			for i := range p.LabelValues {
				if p.LabelValues[i] == v {
					return true
				}
			}
		}
	}

	return false
}

func (p *Permission) isMatchAllValue() bool {
	if p.MatchAll {
		return true
	}

	for i := range p.LabelValues {
		if p.LabelValues[i] == "*" {
			return true
		}
	}

	return false
}

func NewPermissionSet() *PermissionSet {
	return &PermissionSet{
		Items: []*Permission{},
	}
}

func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}