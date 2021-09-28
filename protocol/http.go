package protocol

import (
	"context"
	"fmt"
	"github.com/ducketlab/auth/client"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/mingo/http/middleware/recovery"
	"github.com/ducketlab/mingo/http/router"
	"github.com/ducketlab/mingo/http/router/httprouter"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"net/http"
	"time"
)

func NewHTTPService() *HttpService {
	r := httprouter.New()
	r.Use(recovery.NewWithLogger(zap.L().Named("recovery")))
	r.EnableApiRoot()
	r.Auth(true)
	r.Permission(true)

	server := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Addr:              config.C().App.HttpAddr(),
		Handler:           r,
	}

	return &HttpService{
		router: r,
		server: server,
		logger: zap.L().Named("http"),
		config: config.C(),
	}
}

type HttpService struct {
	router router.Router
	logger logger.Logger
	config *config.Config
	server *http.Server
}

func (s *HttpService) Start() error {

	if err := s.initGrpcClient(); err != nil {
		return err
	}

	if err := pkg.InitV1HttpApi(s.config.App.HttpPrefix, s.router); err != nil {
		return err
	}

	s.logger.Infof("http listen address: %s", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.logger.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

func (s *HttpService) Stop() error {

	s.logger.Info("start http graceful shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Errorf("graceful shutdown timeout, force exit")
	}

	return nil
}

func (s *HttpService) initGrpcClient() error {

	cf := client.NewDefaultConfig()

	cli, err := client.NewClient(cf)

	cf.SetAddress(s.config.App.GrpcAddr())

	if err != nil {
		return err
	}
	client.SetGlobal(cli)

	return err
}
