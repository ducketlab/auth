PROJECT_NAME=auth
MAIN_FILE=main.go
MOD_DIR := $(shell go env GOPATH)
PKG := "github.com/ducketlab/$(PROJECT_NAME)"

codegen: # Init Service
	@protoc -I=. -I${MOD_DIR}/src --go-ext_out=. --go-ext_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} --go-http_out=. --go-http_opt=module=${PKG} pkg/*/pb/*.proto
	@go generate ./...

