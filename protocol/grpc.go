package protocol

import (
	"fmt"
	"github.com/ducketlab/auth/config"
	"github.com/ducketlab/auth/pkg"
	"github.com/ducketlab/mingo/grpc/middleware/recovery"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

type GrpcService struct {
	logger logger.Logger
	config *config.Config
	server *grpc.Server
}

func NewGrpcService() *GrpcService {

	log := zap.L().Named("grpc")

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		rc.UnaryServerInterceptor(),
		pkg.AuthUnaryServerInterceptor(),
	)))

	return &GrpcService{
		config: config.C(),
		logger: log,
		server: grpcServer,
	}
}

func (s *GrpcService) Start() error {

	// Load all grpc service
	pkg.InitGrpcApi(s.server)

	lis, err := net.Listen("tcp", s.config.App.GrpcAddr())
	if err != nil {
		return err
	}

	s.logger.Infof("grpc listen address: %s", s.config.App.GrpcAddr())

	if err := s.server.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.logger.Info("service is stopped")
		}
		return fmt.Errorf("start grpc service error, %s", err.Error())
	}
	return nil
}

func (s *GrpcService) Stop() error {
	s.logger.Info("start grpc service shutdown...")
	s.server.GracefulStop()
	return nil
}
