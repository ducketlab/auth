package domain

import "fmt"

func (req *DescribeDomainRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

func (req *CreateDomainRequest) Validate() error {
	return validate.Struct(req)
}
