package cmd

import (
	"errors"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/protocol"
	"github.com/ducketlab/mingo/logger"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	confType string
	confFile string
)

var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "start service",
	Long:  "start service",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		if err := pkg.InitService(); err != nil {
			return err
		}

		config := config.C()

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

		svr, err := newService(config)
		if err != nil {
			return err
		}

		go svr.waitSign(ch)

		if err := svr.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}

		return nil
	},
}

type service struct {
	grpc *protocol.GrpcService
	log  logger.Logger
}

func (s *service) waitSign(sign chan os.Signal) {
	for {
		select {
		case sg := <-sign:
			switch v := sg.(type) {
			default:
				s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
				if err := s.grpc.Stop(); err != nil {
					s.log.Errorf("grpc graceful shutdown err: %s, force exit", err)
				}
				return
			}
		}
	}
}

func (s *service) start() error {
	s.log.Infof("loaded domain pkg: %v", pkg.LoadedService())

	return s.grpc.Start()
}

func newService(config *config.Config) (*service, error) {
	grpc := protocol.NewGrpcService()

	service := &service{
		grpc: grpc,
	}

	return service, nil
}

func loadGlobalConfig(configType string) error {
	switch configType {
	case "file":
		err := config.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := config.LoadConfigFromEnv()
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown config type")
	}
	return nil
}

func init() {
	serviceCmd.Flags().StringVarP(&confType, "config-type",
		"t", "file", "the service config type [file/env/etcd]")
	serviceCmd.Flags().StringVarP(&confFile, "config-file",
		"f", "etc/auth.toml", "the service config from file")
	RootCmd.AddCommand(serviceCmd)
}
