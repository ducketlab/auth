module github.com/ducketlab/auth

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/caarlos0/env/v6 v6.4.0
	github.com/ducketlab/mingo v0.0.0-20210925164512-04b256925ece
	github.com/go-playground/validator/v10 v10.9.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/rs/xid v1.3.0
	github.com/spf13/cobra v1.2.1
	go.mongodb.org/mongo-driver v1.7.1
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/ducketlab/mingo => ../mingo
