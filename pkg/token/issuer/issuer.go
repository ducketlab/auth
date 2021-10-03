package issuer

import (
	"context"
	"github.com/ducketlab/auth/pkg/token"
)

type Issuer interface {
	CheckClient(ctx context.Context, clientID, clientSecret string) error
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
