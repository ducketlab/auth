package micro

import (
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/types/ftime"
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

var (
	validate = validator.New()
)

func New(req *CreateMicroRequest) (*Micro, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &Micro{
		Id:            xid.New().String(),
		CreateAt:      ftime.Now().Timestamp(),
		UpdateAt:      ftime.Now().Timestamp(),
		Enabled:       true,
		Type:          req.Type,
		Name:          req.Name,
		Label:         req.Label,
		Description:   req.Description,
		ClientId:      token.MakeBearer(24),
		ClientSecret:  token.MakeBearer(32),
		ClientEnabled: true,
	}

	return ins, nil
}

func (m *Micro) ValidateClientCredential(clientSecret string) error {
	if !m.ClientEnabled {
		return exception.NewBadRequest("client not enabled")
	}

	if m.ClientSecret != clientSecret {
		return exception.NewUnauthorized("client credential invalidate")
	}
	return nil
}

func (m *Micro) Desensitize() {
	m.ClientSecret = ""
}

func NewCreateMicroRequest() *CreateMicroRequest {
	return &CreateMicroRequest{
		Label: map[string]string{},
		Type:  Type_CUSTOM,
	}
}