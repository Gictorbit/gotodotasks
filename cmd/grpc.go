package main

import (
	"context"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

func NewGrpcServer(grpcAddr string, logger *zap.Logger) *grpc.Server {
	unaryServerOptions := []grpc.UnaryServerInterceptor{
		grpcCtxTags.UnaryServerInterceptor(grpcCtxTags.WithFieldExtractor(grpcCtxTags.CodeGenRequestFieldExtractor)),
		grpcZap.UnaryServerInterceptor(logger),
		grpcZap.PayloadUnaryServerInterceptor(logger, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return LogRequest == true
		}),
		grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Error("stack trace from panic " + string(debug.Stack()))
			return status.Errorf(codes.Internal, "%v", p)
		})),
		authutil.UnaryServerInterceptor(),
	}
	streamServerOptions := []grpc.StreamServerInterceptor{
		grpcCtxTags.StreamServerInterceptor(grpcCtxTags.WithFieldExtractor(grpcCtxTags.CodeGenRequestFieldExtractor)),
		grpcZap.StreamServerInterceptor(logger),
		grpcZap.PayloadStreamServerInterceptor(logger, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return LogRequest == true
		}),
		grpcRecovery.StreamServerInterceptor(grpcRecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Error("stack trace from panic " + string(debug.Stack()))
			return status.Errorf(codes.Internal, "%v", p)
		})),
		authutil.StreamServerInterceptor(),
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamServerOptions...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryServerOptions...)),
	)
	return grpcServer
}
