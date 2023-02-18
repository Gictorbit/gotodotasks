package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	HostAddress string
	PortGRPC    uint
	PortHTTP    uint
	DatabaseURL string
	LogRequest  bool
)

func main() {
	app := &cli.App{
		Name:  "client",
		Usage: "go file transfer client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Usage:       "host address",
				Value:       "127.0.0.1",
				DefaultText: "127.0.0.1",
				EnvVars:     []string{"HOST_ADDRESS"},
				Destination: &HostAddress,
			},
			&cli.UintFlag{
				Name:        "grpc-port",
				Usage:       "grpc server port",
				Value:       3232,
				DefaultText: "3232",
				Aliases:     []string{"gp"},
				EnvVars:     []string{"GRPC_PORT"},
				Destination: &PortGRPC,
			},
			&cli.UintFlag{
				Name:        "http-port",
				Usage:       "http server port",
				Value:       8686,
				DefaultText: "8686",
				Aliases:     []string{"hp"},
				EnvVars:     []string{"HTTP_PORT"},
				Destination: &PortHTTP,
			},
			&cli.StringFlag{
				Name:        "database",
				Usage:       "database url",
				Aliases:     []string{"db"},
				EnvVars:     []string{"DATABASE_URL"},
				Destination: &DatabaseURL,
			},
			&cli.BoolFlag{
				Name:        "log-request",
				Usage:       "log incoming requests",
				Aliases:     []string{"lgr"},
				Value:       false,
				DefaultText: "false",
				EnvVars:     []string{"LOG_REQUEST"},
				Destination: &LogRequest,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "taskmanager",
				Usage: "run task manager server",
				Action: func(cliCtx *cli.Context) error {
					grpcAddr := net.JoinHostPort(HostAddress, fmt.Sprintf("%d", PortGRPC))
					httpAddr := net.JoinHostPort(HostAddress, fmt.Sprintf("%d", PortHTTP))
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					grpcServer := RunTaskManagerGRPCServer(DatabaseURL, grpcAddr)
					httpServer := RunTaskManagerHTTPService(ctx, grpcAddr, httpAddr)

					stop := make(chan os.Signal, 1)
					signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
					<-stop

					log.Println("shutting down servers")
					grpcServer.GracefulStop()
					if err := httpServer.Shutdown(context.Background()); err != nil {
						return fmt.Errorf("HTTP server shutdown error: %v", err)
					}
					return nil
				},
			},
		},
	}
	if e := app.Run(os.Args); e != nil {
		log.Println("failed to run app", e)
	}
}
