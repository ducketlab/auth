package client

import (
	"context"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/grpc/gcontext"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/logger/zap"
	httpb "github.com/ducketlab/mingo/pb/http"
	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

type PathEntryHandleFunc func(path string) *httpb.Entry

func NewGrpcAuthAuther(hf PathEntryHandleFunc, c *Client) *GrpcAuther {
	return &GrpcAuther{
		hf:    hf,
		client:     c,
		logger:     zap.L().Named("grpc-interceptor"),
	}
}

type GrpcAuther struct {
	hf    PathEntryHandleFunc
	logger     logger.Logger
	client     *Client
}

func (a *GrpcAuther) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.auth
}

func (a *GrpcAuther) SetLogger(l logger.Logger) {
	a.logger = l
}

func (a *GrpcAuther) auth(
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
					gcontext.ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
					gcontext.ResponseReasonHeader, t.Reason(),
					gcontext.ResponseDescHeader, t.Error(),
				)
				if err := grpc.SetTrailer(ctx, trailer); err != nil {
					a.log().Errorf("send grpc trailer error, %s", err)
				}
			}
		}
	}()

	rctx, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := a.validateServiceCredential(rctx); err != nil {
		return nil, err
	}

	entry := a.hf(info.FullMethod)
	if entry == nil {
		return nil, status.Errorf(codes.Internal, "entry not found, check is registry")
	}

	engine := newEntryEngine(a.client, entry, a.log())

	tk, err := engine.ValidateIdentity(rctx)
	if err != nil {
		return nil, err
	}

	if err := engine.ValidatePermission(tk, rctx); err != nil {
		return nil, err
	}

	rid := rctx.GetRequestID()
	if rid == "" {
		rid = xid.New().String()
		rctx.SetRequestId(rid)
	}

	return handler(rctx.Context(), req)
}

func (a *GrpcAuther) validateServiceCredential(ctx *gcontext.GrpcInCtx) error {
	clientID := ctx.GetClientId()
	clientSecret := ctx.GetClientSecret()
	a.log().Debugf("start check client[%s] credential ...", clientID)

	if clientID == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := a.client.Micro().ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	a.log().Debugf("check client[%s] credential ok", clientID)
	return nil
}

func (a *GrpcAuther) log() logger.Logger {
	if a == nil {
		a.logger = zap.L().Named("grpc-auther")
	}

	return a.logger
}