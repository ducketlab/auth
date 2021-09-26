package pkg

import (
	"context"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

var (
	interceptor = newGrpcAuther()
)

type grpcAuther struct {
	l logger.Logger
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return interceptor.Auth
}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if err != nil {
			switch t := err.(type) {
			case exception.APIException:
				err = status.Errorf(codes.Code(t.ErrorCode()), t.Error())
				trailer := metadata.Pairs(
					ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
					ResponseReasonHeader, t.Reason(),
					ResponseDescHeader, t.Error(),
				)
				if err := grpc.SetTrailer(ctx, trailer); err != nil {
					a.log().Errorf("send grpc trailer error, %s", err)
				}
			}
		}
	}()

	rctx, err := GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	rctx.ClearInternal()

	resp, err = handler(rctx.ClearInternal().Context(), req)

	return resp, err
}

func (a *grpcAuther) log() logger.Logger {
	if a.l == nil {
		a.l = zap.L().Named("Grpc Auther")
	}

	return a.l
}

