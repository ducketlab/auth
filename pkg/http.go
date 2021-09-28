package pkg

import (
	"github.com/ducketlab/mingo/http/router"
	"strings"
)

var (
	httpAPIs = make(map[string]HttpApi)
)

func LoadedHttp() []string {
	var apis []string
	for k := range httpAPIs {
		apis = append(apis, k)
	}
	return apis
}

type HttpApi interface {
	Registry(router.SubRouter)
	Config() error
}

func RegistryV1Http(name string, api HttpApi) {
	if _, ok := httpAPIs[name]; ok {
		panic("http api " + name + " has registry")
	}
	httpAPIs[name] = api
}

func InitV1HttpApi(pathPrefix string, root router.Router) error {
	for _, api := range httpAPIs {
		if err := api.Config(); err != nil {
			return err
		}

		if pathPrefix != "" && !strings.HasPrefix(pathPrefix, "/") {
			pathPrefix = "/" + pathPrefix
		}

		api.Registry(root.SubRouter(pathPrefix + "/v1"))
	}

	return nil
}