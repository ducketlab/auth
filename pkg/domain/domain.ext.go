package domain

import (
	"errors"
	"github.com/ducketlab/mingo/types/ftime"
)

func New(owner string, req *CreateDomainRequest) (*Domain, error) {

	if owner == "" {
		return nil, errors.New("domain required owner")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Domain{
		CreateAt:        ftime.Now().Timestamp(),
		UpdateAt:        ftime.Now().Timestamp(),
		Owner:           owner,
		Name:            req.Name,
		Enabled:         true,
	}

	return d, nil
}

func NewDomainSet() *Set {
	return &Set{
		Items: []*Domain{},
	}
}

func (s *Set) Add(d *Domain) {
	s.Items = append(s.Items, d)
}