package pkg

import (
	"context"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	"github.com/ducketlab/mingo/pb/http"
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

	if err := a.validateServiceCredential(rctx); err != nil {
		return nil, err
	}

	entry := GetPathEntry(info.FullMethod)
	if entry == nil {
		return nil, status.Errorf(codes.Internal,
			"entry grpc path: %s not found, check is registry", info.FullMethod)
	}

	if entry.AuthEnable {
		tk, err := a.checkToken(rctx)
		if err != nil {
			return nil, err
		}

		if err := a.validatePermission(tk, entry, req); err != nil {
			return nil, err
		}
	}

	rctx.ClearInternal()

	resp, err = handler(rctx.ClearInternal().Context(), req)

	return resp, err
}

func (a *grpcAuther) validateServiceCredential(ctx *GrpcInCtx) error {
	clientID := ctx.GetClientID()
	clientSecret := ctx.GetClientSecret()

	if clientID == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	if Micro == nil {
		return status.Errorf(codes.Internal, "micro service is not initial")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := Micro.ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}

func (a *grpcAuther) checkToken(ctx *GrpcInCtx) (*token.Token, error) {
	accessToken := ctx.GetAccessToken()
	req := token.NewValidateTokenRequest()
	if accessToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "access_token meta required")
	}
	req.AccessToken = accessToken
	return Token.ValidateToken(context.Background(), req)
}

func (a *grpcAuther) log() logger.Logger {
	if a.l == nil {
		a.l = zap.L().Named("grpc-auther")
	}

	return a.l
}

func (a *grpcAuther) validatePermission(tk *token.Token, entry *http.Entry, req interface{}) error {
	return nil
}