package client

import "context"

type Authentication struct {
	clientId     string
	clientSecret string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		"client-id":     a.clientId,
		"client-secret": a.clientSecret,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func (a *Authentication) SetClientCredentials(clientID, clientSecret string) {
	a.clientId = clientID
	a.clientSecret = clientSecret
}



