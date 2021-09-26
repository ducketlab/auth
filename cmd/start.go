package cmd

import (
	"errors"
	"fmt"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/auth/protocol"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/ducketlab/auth/pkg/all"
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

		if err := loadGlobalComponent(); err != nil {
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
	http *protocol.HttpService
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
				s.log.Info("grpc service stop complete")

				if err := s.http.Stop(); err != nil {
					s.log.Errorf("http graceful shutdown err: %s, force exit", err)
				}
				s.log.Infof("http service stop complete")
				return
			}
		}
	}
}

func (s *service) start() error {
	s.log.Infof("loaded service pkg: %v", pkg.LoadedService())
	s.log.Infof("loaded http service: %s", pkg.LoadedHttp())

	go s.grpc.Start()
	return s.http.Start()
}

func newService(config *config.Config) (*service, error) {

	log := zap.L().Named("cli")

	grpc := protocol.NewGrpcService()

	http := protocol.NewHTTPService()

	service := &service{
		grpc: grpc,
		log: log,
		http: http,
	}

	return service, nil
}

func loadGlobalComponent() error {
	if err := loadGlobalLogger(); err != nil {
		return err
	}
	return nil
}

func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	lc := config.C().Log
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level

	switch lc.To {
	case config.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case config.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	switch lc.Format {
	case config.JSONFormat:
		zapConfig.Json = true
	}

	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("init").Info(logInitMsg)
	return nil
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
