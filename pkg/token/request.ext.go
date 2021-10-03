package token

import (
	"errors"
	"fmt"
	"github.com/ducketlab/mingo/http/request"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var (
	validate = validator.New()
)

func NewDescribeTokenRequest() *DescribeTokenRequest {
	return &DescribeTokenRequest{}
}

func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

func (m *IssueTokenRequest) WithUserAgent(userAgent string) {
	m.UserAgent = userAgent
}

func (m *IssueTokenRequest) WithRemoteIPFromHTTP(r *http.Request) {
	m.RemoteIp = request.GetRemoteIP(r)
}

func NewDescribeTokenRequestWithAccessToken(at string) *DescribeTokenRequest {
	req := NewDescribeTokenRequest()
	req.AccessToken = at
	return req
}

func (m *DescribeTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	if m.AccessToken == "" && m.RefreshToken == "" {
		return errors.New(
			"describe token request validate error, access_token and refresh_token required one")
	}

	return nil
}

func (m *IssueTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	switch m.GrantType {
	case GrantType_PASSWORD:
		if m.Username == "" || m.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", GrantType_PASSWORD)
		}
	case GrantType_REFRESH:
		if m.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", GrantType_REFRESH)
		}
		if m.RefreshToken == "" {
			return fmt.Errorf("use %s grant type, refresh_token required", GrantType_REFRESH)
		}
	case GrantType_ACCESS:
		if m.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", GrantType_ACCESS)
		}
	case GrantType_LDAP:
		if m.Username == "" || m.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", GrantType_LDAP)
		}
	case GrantType_CLIENT:
	case GrantType_AUTH_CODE:
		if m.AuthCode == "" {
			return fmt.Errorf("use %s grant type, code required", GrantType_AUTH_CODE)
		}
	default:
		return fmt.Errorf("unknown grant type %s", m.GrantType)
	}

	return nil
}