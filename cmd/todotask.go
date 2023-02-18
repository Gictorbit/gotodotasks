package main

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	taskdb "github.com/Gictorbit/gotodotasks/internal/db/postgres"
	"github.com/Gictorbit/gotodotasks/internal/service/taskmanager"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	taskSrv := taskmanager.NewTodoTaskManager(logger, taskManagerDB)

	grpcServer := grpc.NewServer()
	taskpb.RegisterTodoTaskServiceServer(grpcServer, taskSrv) // register your service implementation

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

func RunTaskManagerHTTPService(ctx context.Context, grpcAddr, httpAddr string) *http.Server {
	gwMux := runtime.NewServeMux()
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
		log.Printf("starting HTTP server on %s", httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return httpServer
}
