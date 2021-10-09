package all

import (
	_ "github.com/ducketlab/auth/pkg/domain/grpc"
	_ "github.com/ducketlab/auth/pkg/domain/http"
	_ "github.com/ducketlab/auth/pkg/endpoint/grpc"
	_ "github.com/ducketlab/auth/pkg/endpoint/http"
	_ "github.com/ducketlab/auth/pkg/micro/grpc"
	_ "github.com/ducketlab/auth/pkg/micro/http"
	_ "github.com/ducketlab/auth/pkg/namespace/grpc"
	_ "github.com/ducketlab/auth/pkg/namespace/http"
	_ "github.com/ducketlab/auth/pkg/token/grpc"
	_ "github.com/ducketlab/auth/pkg/token/http"
	_ "github.com/ducketlab/auth/pkg/user/grpc"
	_ "github.com/ducketlab/auth/pkg/user/http"
)
