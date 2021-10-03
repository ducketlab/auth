package issuer

import (
	"context"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"time"
)

func NewTokenIssuer() (Issuer, error) {

	issuer := &issuer{
		user:    pkg.User,
		domain:  pkg.Domain,
		token:   pkg.Token,
		log:     zap.L().Named("issuer"),
	}
	return issuer, nil
}

type issuer struct {
	token   token.TokenServiceServer
	user    user.UserServiceServer
	domain  domain.DomainServiceServer
	log     logger.Logger
}

func (i *issuer) checkUserPass(user, pass string) (*user.User, error) {
	ctx := pkg.NewInternalMockGrpcCtx(user).Context()
	u, err := i.getUser(ctx, user)
	if err != nil {
		return nil, err
	}

	if err := u.HashedPassword.CheckPassword(pass); err != nil {
		return nil, err
	}
	return u, nil
}

func (i *issuer) getUser(ctx context.Context, name string) (*user.User, error) {
	req := user.NewDescribeAccountRequest()
	req.Account = name
	return i.user.DescribeAccount(ctx, req)
}

func (i *issuer) IssueToken(ctx context.Context,
	req *token.IssueTokenRequest) (*token.Token, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	switch req.GrantType {
	case token.GrantType_PASSWORD:
		u, checkErr := i.checkUserPass(req.Username, req.Password)
		if checkErr != nil {
			i.log.Debugf("issue password token error, %s", checkErr)
			return nil, checkErr
		}

		tk := i.issueUserToken(u, token.GrantType_PASSWORD)

		switch u.Type {
		case user.UserType_SUPPER, user.UserType_PRIMARY:
			tk.Domain = domain.InternalDomainName
		case user.UserType_SUB:
			tk.Domain = u.Domain
		}

		return tk, nil
	default:
		return nil, exception.NewInternalServerError("unknown grant type %s", req.GrantType)
	}
}

func (i *issuer) issueUserToken(u *user.User, gt token.GrantType) *token.Token {
	now := time.Now()
	tk := &token.Token{
		Type:            token.TokenType_BEARER,
		AccessToken:     token.MakeBearer(24),
		RefreshToken:    token.MakeBearer(32),
		CreateAt:        now.UnixNano() / 1000000,
		Account:         u.Account,
		UserType:        u.Type,
		ClientId:        "",
		GrantType:       gt,
		ApplicationId:   "",
		ApplicationName: "",
	}
	return tk
}
