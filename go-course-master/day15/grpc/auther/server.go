package auther

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

// GrpcAuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return newGrpcAuther().Auth
}

// GrpcAuthStreamServerInterceptor returns a new unary server interceptor for auth.
func GrpcAuthStreamServerInterceptor() grpc.StreamServerInterceptor {
	return newGrpcAuther().StreamAuth
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{
		log: zap.L().Named("Grpc Auther"),
	}
}

// internal todo
type grpcAuther struct {
	log logger.Logger
}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	fmt.Println("gprc header info: ", md)

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	resp, err = handler(ctx, req)
	return resp, err
}

func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cids := md.Get(ClientHeaderKey)
	sids := md.Get(ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	if !(clientId == "admin" && clientSecret == "123456") {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret invalidate")
	}

	return nil
}

func (a *grpcAuther) StreamAuth(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	fmt.Println(srv, info)
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return fmt.Errorf("ctx is not an grpc incoming context")
	}
	fmt.Println("gprc header info: ", md)

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return err
	}

	return handler(srv, ss)
}
