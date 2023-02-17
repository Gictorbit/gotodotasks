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
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start the gRPC server

	logger, _ := zap.NewProduction()
	taskManagerDB := taskdb.NewTodoTaskDB(nil)
	taskSrv := taskmanager.NewTodoTaskManager(logger, taskManagerDB)

	grpcServer := grpc.NewServer()
	taskpb.RegisterTodoTaskServiceServer(grpcServer, taskSrv) // register your service implementation

	grpcAddr := ":9090"
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

	// Start the HTTP server
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	httpAddr := ":8080"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	// Handle shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down servers")

	grpcServer.GracefulStop()
	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
}
