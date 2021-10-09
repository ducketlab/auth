package endpoint

import (
	"fmt"
	"github.com/ducketlab/mingo/pb/http"
	"github.com/ducketlab/mingo/types/ftime"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (req *RegistryRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require *router.Entry")
	}

	return validate.Struct(req)
}

func (req *RegistryRequest) Endpoints(serviceId string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			Id:        GenHashId(serviceId, req.Entries[i].Path),
			CreateAt:  ftime.Now().Timestamp(),
			UpdateAt:  ftime.Now().Timestamp(),
			ServiceId: serviceId,
			Version:   req.Version,
			Entry:     req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

func NewRegistryResponse(message string) *RegistryResponse {
	return &RegistryResponse{Message: message}
}

func NewDescribeEndpointRequestWithId(id string) *DescribeEndpointRequest {
	return &DescribeEndpointRequest{Id: id}
}

func (req *DescribeEndpointRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}

func NewDefaultRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Entries: []*http.Entry{},
	}
}