package all

import (
	_ "github.com/ducketlab/auth/pkg/domain/grpc"
	_ "github.com/ducketlab/auth/pkg/domain/http"
	_ "github.com/ducketlab/auth/pkg/token/grpc"
	_ "github.com/ducketlab/auth/pkg/token/http"
	_ "github.com/ducketlab/auth/pkg/user/grpc"
)
