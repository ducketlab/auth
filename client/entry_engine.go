package client

import (
	"github.com/ducketlab/auth/pkg/endpoint"
	"github.com/ducketlab/auth/pkg/micro"
	"github.com/ducketlab/auth/pkg/permission"
	"github.com/ducketlab/auth/pkg/token"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/grpc/gcontext"
	"github.com/ducketlab/mingo/logger"
	"github.com/ducketlab/mingo/pb/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func newEntryEngine(client *Client, entry *http.Entry, log logger.Logger) *entryEngine {
	return &entryEngine{
		client:  client,
		Entry:   entry,
		log:     log,
		uniPath: false,
	}
}

type entryEngine struct {
	*http.Entry
	client    *Client
	serviceId string
	log       logger.Logger
	uniPath   bool
}

func (e *entryEngine) UseUniPath() {
	e.uniPath = true
}

func (e *entryEngine) ValidateIdentity(ctx *gcontext.GrpcInCtx) (*token.Token, error) {
	e.log.Debug("start token identity check")

	if !e.AuthEnable {
		e.log.Debug("auth disabled skip")
		return nil, nil
	}

	accessToken := ctx.GetAccessToKen()

	if accessToken == "" {
		e.log.Debugf("[%s] auth enabled, but not get access token", e.Path)
		return nil, exception.NewBadRequest("token required")
	}

	req := token.NewValidateTokenRequest()
	req.AccessToken = accessToken

	outCtx := gcontext.NewGrpcOutCtx()
	outCtx.SetAccessToken(ctx.GetAccessToKen())

	var header, trailer metadata.MD
	tk, err := e.client.Token().ValidateToken(
		outCtx.Context(), req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, gcontext.NewExceptionFromTrailer(trailer, err)
	}

	e.log.Debugf("token check ok, username: %s", tk.Account)
	return tk, nil
}

func (e *entryEngine) ValidatePermission(tk *token.Token, ctx *gcontext.GrpcInCtx) error {
	if !e.AuthEnable {
		return nil
	}

	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	if !e.PermissionEnable {
		e.log.Debugf("[%s] permission disabled skip check!", tk.Account)
		return nil
	}

	if e.RequiredNamespace && tk.Namespace == "" {
		return exception.NewBadRequest("namespace required!")
	}

	path := e.Path
	if e.uniPath {
		path = e.UniquePath()
	}

	outCtx := gcontext.NewGrpcOutCtx()
	outCtx.SetAccessToken(ctx.GetAccessToKen())

	eid, err := e.endpointHashId(outCtx, path)

	if err != nil {
		return err
	}

	e.log.Debugf("[%s] start check permission to auth ...", tk.Account)

	req := permission.NewCheckPermissionRequest()
	req.EndpointId = eid
	req.NamespaceId = tk.Namespace

	var header, trailer metadata.MD

	perm, err := e.client.Permission().CheckPermission(
		outCtx.Context(), req, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		return gcontext.NewExceptionFromTrailer(trailer, err)
	}

	tk.Scope = perm.Scope

	e.log.Debugf("[%s] permission check passed", tk.Account)
	return nil
}

func (e *entryEngine) endpointHashId(ctx *gcontext.GrpcOutCtx, path string) (string, error) {
	if e.serviceId == "" {
		descReq := micro.NewDescribeServiceRequestWithClientId(e.client.GetClientId())
		svr, err := e.client.Micro().DescribeService(ctx.Context(), descReq)
		if err != nil {
			return "", err
		}
		e.serviceId = svr.Id
	}

	return endpoint.GenHashId(e.serviceId, path), nil
}
