package main

import (
	"context"
	"crypto/rand"
	"fmt"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	userdb "github.com/Gictorbit/gotodotasks/internal/db/postgres/user"
	"github.com/Gictorbit/gotodotasks/internal/service/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
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
	userDatabase, err := userdb.NewUserDBFromDsn(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	authManager := authutil.NewAuthManager(SecretKey, Issuer, TokenValidTime)
	userServer := user.NewUserService(logger, userDatabase, authManager)

	grpcServer := NewGrpcServer(grpcAddr, logger)
	userpb.RegisterUserServiceServer(grpcServer, userServer) // register your service implementation

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Printf("starting user gRPC server on %s", grpcAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return grpcServer
}

func RunUserHTTPService(ctx context.Context, grpcAddr, httpAddr string) *http.Server {
	gwMux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		md := metadata.Pairs("token", request.Header.Get("token"))
		return md
	}))
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil {
		log.Fatalf("failed to register gRPC-Gateway: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	httpServer := &http.Server{Addr: httpAddr, Handler: mux}

	go func() {
		log.Printf("starting user HTTP server on %s", httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return httpServer
}
