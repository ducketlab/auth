package issuer

import "context"

func (i *issuer) CheckClient(ctx context.Context, clientID, clientSecret string) error {
	return nil
}
