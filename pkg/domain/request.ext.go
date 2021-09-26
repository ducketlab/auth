package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (req *DescribeDomainRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

func (req *CreateDomainRequest) Validate() error {
	return validate.Struct(req)
}
