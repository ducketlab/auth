package pkg

import (
	"context"
	"fmt"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/user"
	"google.golang.org/grpc/metadata"
)

const (
	InternalCallTokenHeader = "internal-call-token"
	ClientIdHeader          = "client-id"
	ClientSecretHeader      = "client-secret"
	OauthTokenHeader        = "x-oauth-token"
	RequestIdHeader         = "x-request-id"
)

type grpcCtx struct {
	md metadata.MD
}

func newGrpcCtx(md metadata.MD) *grpcCtx {
	return &grpcCtx{md: md}
}

func (c *grpcCtx) get(key string) string {
	return c.getWithIndex(key, 0)
}

func (c *grpcCtx) getWithIndex(key string, index int) string {
	if val, ok := c.md[key]; ok {
		if len(val) > index {
			return val[index]
		}
	}
	return ""
}

func (c *grpcCtx) set(key string, values ...string) {
	c.md.Set(key, values...)
}

func (c *grpcCtx) SetAccessToken(ak string) {
	c.set(OauthTokenHeader, ak)
}

func (c *grpcCtx) SetClientCredentials(clientID, clientSecret string) {
	c.set(ClientIdHeader, clientID)
	c.set(ClientSecretHeader, clientSecret)
}

func (c *grpcCtx) SetRequestID(rid string) {
	c.set(RequestIdHeader, rid)
}

type GrpcInCtx struct {
	*grpcCtx
}

func NewGrpcInCtx() *GrpcInCtx {
	return &GrpcInCtx{newGrpcCtx(metadata.Pairs())}
}

func (c *GrpcInCtx) Context() context.Context {
	return metadata.NewIncomingContext(context.Background(), c.md)
}

func (c *GrpcInCtx) GetClientID() string {
	return c.get(ClientIdHeader)
}

func (c *GrpcInCtx) GetClientSecret() string {
	return c.get(ClientSecretHeader)
}

func (c *GrpcInCtx) GetAccessToKen() string {
	return c.get(OauthTokenHeader)
}

func (c *GrpcInCtx) GetRequestID() string {
	return c.get(RequestIdHeader)
}

func (c *GrpcInCtx) SetIsInternalCall(account, domain string) {
	c.set(InternalCallTokenHeader, account, domain)
}

func (c *GrpcInCtx) IsInternalCall() bool {
	if _, ok := c.md[InternalCallTokenHeader]; ok {
		return true
	}

	return false
}

func NewInternalMockGrpcCtx(account string) *GrpcInCtx {
	ctx := NewGrpcInCtx()
	ctx.SetIsInternalCall(account, domain.InternalDomainName)
	return ctx
}

func (c *GrpcInCtx) ClearInternal() *GrpcInCtx {
	delete(c.md, InternalCallTokenHeader)
	return c
}

func (c *GrpcInCtx) GetToken() (*token.Token, error) {
	req := token.NewDescribeTokenRequestWithAccessToken(c.GetAccessToKen())
	ctx := NewInternalMockGrpcCtx("internal").Context()
	return Token.DescribeToken(ctx, req)
}

func GetGrpcInCtx(ctx context.Context) (*GrpcInCtx, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}
	return &GrpcInCtx{newGrpcCtx(md)}, nil
}

func GetTokenFromGrpcInCtx(ctx context.Context) (*token.Token, error) {
	rctx, err := GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if rctx.IsInternalCall() {
		return rctx.InternalCallToken(), nil
	}

	return rctx.GetToken()
}

func (c *GrpcInCtx) InternalCallToken() *token.Token {
	tk := &token.Token{UserType: user.UserType_INTERNAL}
	tk.Account = c.getWithIndex(InternalCallTokenHeader, 0)
	tk.Domain = c.getWithIndex(InternalCallTokenHeader, 1)
	return tk
}
