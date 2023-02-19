package main

import (
	"context"
	"crypto/rand"
	"fmt"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	userdb "github.com/Gictorbit/gotodotasks/internal/db/postgres/user"
	"github.com/Gictorbit/gotodotasks/internal/service/user"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"runtime/debug"
)

func GenerateRandomSecretKey(len int) (string, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}

func RunUserGRPCServer(databaseURL, grpcAddr string) *grpc.Server {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
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
	userDatabase, err := userdb.NewUserDBFromDsn(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	authManager := authutil.NewAuthManager(SecretKey, Issuer, TokenValidTime)
	userServer := user.NewUserService(logger, userDatabase, authManager)

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamServerOptions...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryServerOptions...)),
	)
	userpb.RegisterUserServiceServer(grpcServer, userServer) // register your service implementation

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Printf("starting gRPC server on %s", grpcAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return grpcServer
}

func RunUserHTTPService(ctx context.Context, grpcAddr, httpAddr string) *http.Server {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil { // register the gRPC-Gateway handler
		log.Fatalf("failed to register gRPC-Gateway: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	httpServer := &http.Server{Addr: httpAddr, Handler: mux}

	go func() {
		log.Printf("starting HTTP server on %s", httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return httpServer
}
