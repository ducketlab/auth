module github.com/ducketlab/auth

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/caarlos0/env/v6 v6.4.0
	github.com/ducketlab/mingo v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator/v10 v10.9.0
	github.com/spf13/cobra v1.2.1
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/ducketlab/mingo => ../mingo
