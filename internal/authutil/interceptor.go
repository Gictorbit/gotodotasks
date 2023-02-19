package authutil

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServerInterface interface {
	GetAuthManager() *AuthManager
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if service, ok := info.Server.(AuthServerInterface); ok {
			authManager := service.GetAuthManager()
			if authManager == nil {
				return handler(ctx, req)
			}
			claims, err := authManager.ExtractContext(ctx)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}
			newCtx := context.WithValue(ctx, ClaimKey, claims)
			return handler(newCtx, req)
		}
		return handler(ctx, req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var newCtx context.Context
		if authSrv, ok := srv.(AuthServerInterface); ok {
			authManager := authSrv.GetAuthManager()
			if authManager != nil {
				claim, err := authManager.ExtractContext(stream.Context())
				if err != nil {
					return status.Error(codes.Unauthenticated, err.Error())
				}
				newCtx = context.WithValue(stream.Context(), ClaimKey, claim)
			} else {
				newCtx = stream.Context()
			}
		} else {
			newCtx = stream.Context()
		}
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}
