package policy

import (
	"context"
	"fmt"
	"github.com/ducketlab/auth/pkg/namespace"
	"github.com/ducketlab/auth/pkg/role"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/types/ftime"
	"hash/fnv"
)

func New(tk *token.Token, req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if tk == nil {
		return nil, exception.NewUnauthorized("token required")
	}

	p := &Policy{
		CreateAt:    ftime.Now().Timestamp(),
		UpdateAt:    ftime.Now().Timestamp(),
		Creater:     tk.Account,
		Domain:      tk.Domain,
		NamespaceId: req.NamespaceId,
		Account:     req.Account,
		RoleId:      req.RoleId,
		Scope:       req.Scope,
		ExpiredTime: req.ExpiredTime,
		Type:        req.Type,
	}
	p.genID()

	return p, nil
}

func NewDefaultPolicy() *Policy {
	return &Policy{}
}

func (p *Policy) genID() {
	h := fnv.New32a()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s",
		p.Domain, p.NamespaceId, p.Account, p.RoleId)

	h.Write([]byte(hashedStr))
	p.Id = fmt.Sprintf("%x", h.Sum32())
}

func (p *Policy) CheckDependence(ctx context.Context, u user.UserServiceServer, r role.RoleServiceServer, ns namespace.NamespaceServiceServer) (*user.User, error) {
	account, err := u.DescribeAccount(ctx, user.NewDescribeAccountRequestWithAccount(p.Account))
	if err != nil {
		return nil, fmt.Errorf("check user error, %s", err)
	}

	_, err = r.DescribeRole(ctx, role.NewDescribeRoleRequestWithId(p.RoleId))
	if err != nil {
		return nil, fmt.Errorf("check role error, %s", err)
	}

	if !p.IsAllNamespace() {
		_, err = ns.DescribeNamespace(ctx, namespace.NewNewDescribeNamespaceRequestWithId(p.NamespaceId))
		if err != nil {
			return nil, fmt.Errorf("check namespace error, %s", err)
		}
	}

	return account, nil
}

func (p *Policy) IsAllNamespace() bool {
	return p.NamespaceId == "*"
}

func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}
