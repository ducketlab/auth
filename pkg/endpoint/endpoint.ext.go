package endpoint

import (
	"fmt"
	"github.com/ducketlab/mingo/pb/http"
	"github.com/ducketlab/mingo/types/ftime"
	"hash/fnv"
)

func NewDefaultEndpoint() *Endpoint {
	return &Endpoint{}
}

func NewEndpoint(serviceId, version string, entry http.Entry) *Endpoint {
	return &Endpoint{
		Id:        GenHashId(serviceId, entry.Path),
		CreateAt:  ftime.Now().Timestamp(),
		UpdateAt:  ftime.Now().Timestamp(),
		ServiceId: serviceId,
		Version:   version,
		Entry:     &entry,
	}
}

func GenHashId(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

func NewEndpointSet() *Set {
	return &Set{
		Items: []*Endpoint{},
	}
}

func (s *Set) Add(e *Endpoint) {
	s.Items = append(s.Items, e)
}



