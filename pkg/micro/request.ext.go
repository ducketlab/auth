package micro

import "errors"

func (req *CreateMicroRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DescribeMicroRequest) Validate() error {
	if req.Id == "" && req.Name == "" && req.ClientId == "" {
		return errors.New("id, name or client_id is required")
	}

	return nil
}

func NewDescribeServiceRequest() *DescribeMicroRequest {
	return &DescribeMicroRequest{}
}

func NewDescribeServiceRequestWithClientId(clientId string) *DescribeMicroRequest {
	req := NewDescribeServiceRequest()
	req.ClientId = clientId
	return req
}

func NewValidateClientCredentialRequest(clientID, clientSecret string) *ValidateClientCredentialRequest {
	return &ValidateClientCredentialRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}
}