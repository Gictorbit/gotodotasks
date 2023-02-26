package main

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	taskdb "github.com/Gictorbit/gotodotasks/internal/db/postgres/taskmanager"
	"github.com/Gictorbit/gotodotasks/internal/service/taskmanager"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
)

func RunTaskManagerGRPCServer(databaseURL, grpcAddr string) *grpc.Server {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	taskManagerDB, err := taskdb.NewTodoTaskFromDsn(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := NewGrpcServer(grpcAddr, logger)
	authManager := authutil.NewAuthManager(SecretKey, Domain, TokenValidTime)
	taskSrv := taskmanager.NewTodoTaskManager(logger, taskManagerDB, authManager)
	taskpb.RegisterTodoTaskServiceServer(grpcServer, taskSrv) // register your service implementation

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Printf("starting taskmanager gRPC server on %s", grpcAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return grpcServer
}

func RunTaskManagerHTTPService(ctx context.Context, grpcAddr, httpAddr string) *http.Server {
	gwMux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		md := metadata.Pairs("token", request.Header.Get("token"))
		return md
	}))
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := taskpb.RegisterTodoTaskServiceHandlerFromEndpoint(ctx, gwMux, grpcAddr, opts); err != nil { // register the gRPC-Gateway handler
		log.Fatalf("failed to register gRPC-Gateway: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	httpServer := &http.Server{Addr: httpAddr, Handler: mux}

	go func() {
		log.Printf("starting taskmanager HTTP server on %s", httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return httpServer
}
