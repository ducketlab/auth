package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/pkg/domain"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/auth/pkg/user"
	"github.com/ducketlab/auth/version"
	"github.com/spf13/cobra"
	"strings"
)

var InitCmd = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		if err := loadGlobalLogger(); err != nil {
			return err
		}

		if err := pkg.InitService(); err != nil {
			return err
		}

		initer, err := NewInitialerFromCLI()
		if err != nil {
			return err
		}
		if err := initer.Run(); err != nil {
			return err
		}
		return nil
	},
}

func NewInitialerFromCLI() (*Initialer, error) {
	i := NewInitialer()

	if err := i.checkIsInit(); err != nil {
		return nil, err
	}

	i.username = "admin"
	i.password = "admin"

	return i, nil
}

func NewInitialer() *Initialer {
	return &Initialer{}
}

type Initialer struct {
	domainDisplayName string
	username          string
	password          string
	tk                *token.Token
}

func (i *Initialer) Run() error {
	fmt.Println("start init...")

	u, err := i.initUser()
	if err != nil {
		return err
	}

	fmt.Printf("user: %s [success]\n", u.Account)

	svr, err := i.initService()
	if err != nil {
		return err
	}

	fmt.Printf("Service: %s %s %s \n",
		svr.Name, svr.ClientId, svr.ClientSecret)

	return nil
}

func (i *Initialer) mockContext(account string) context.Context {
	ctx := pkg.NewGrpcInCtx()
	ctx.SetIsInternalCall(account, domain.InternalDomainName)
	return ctx.Context()
}

func (i *Initialer) checkIsInit() error {
	req := user.NewQueryAccountRequest()
	req.UserType = user.UserType_SUPPER
	userSet, err := pkg.User.QueryAccount(i.mockContext("internal"), req)
	if err != nil {
		return err
	}

	if userSet.Total > 0 {
		return errors.New("supper admin user has exist")
	}
	return nil
}

func (i *Initialer) initUser() (*user.User, error) {
	req := user.NewCreateUserRequest()
	req.UserType = user.UserType_SUPPER
	req.Account = strings.TrimSpace(i.username)
	req.Password = strings.TrimSpace(i.password)
	return pkg.User.CreateAccount(i.mockContext(i.username), req)
}

func (i *Initialer) userContext() context.Context {
	ctx := pkg.NewGrpcInCtx()
	ctx.SetAccessToken(i.tk.AccessToken)
	return ctx.Context()
}

func (i *Initialer) initService() (*micro.Micro, error) {
	req := micro.NewCreateMicroRequest()
	req.Name = version.ServiceName
	req.Type = micro.Type_BUILD_IN
	return pkg.Micro.CreateService(i.mockContext("internal"), req)
}

func init() {
	InitCmd.Flags().StringVarP(&confType, "config-type",
		"t", "file", "the service config type [file/env/etcd]")
	InitCmd.Flags().StringVarP(&confFile, "config-file",
		"f", "etc/auth.toml", "the service config from file")
	RootCmd.AddCommand(InitCmd)
}
