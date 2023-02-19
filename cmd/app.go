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
	"time"
)

var (
	HostAddress    string
	PortGRPC       uint
	PortHTTP       uint
	DatabaseURL    string
	LogRequest     bool
	SecretKey      string
	TokenValidTime time.Duration
	Issuer         string
)

func main() {
	randSecret, err := GenerateRandomSecretKey(10)
	if err != nil {
		log.Fatal(err)
	}
	app := &cli.App{
		Name:  "gotodotask",
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
				Value:       3000,
				DefaultText: "3000",
				Aliases:     []string{"gp"},
				EnvVars:     []string{"GRPC_PORT"},
				Destination: &PortGRPC,
			},
			&cli.UintFlag{
				Name:        "http-port",
				Usage:       "http server port",
				Value:       4000,
				DefaultText: "4000",
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
			&cli.StringFlag{
				Name:        "secret",
				Usage:       "jwt secret",
				Value:       randSecret,
				DefaultText: randSecret,
				EnvVars:     []string{"JWT_SECRET"},
				Destination: &SecretKey,
			},
			&cli.StringFlag{
				Name:        "issuer",
				Usage:       "issuer name",
				Value:       "go.task.todo",
				DefaultText: "go.task.todo",
				EnvVars:     []string{"JWT_SECRET"},
				Destination: &Issuer,
			},
			&cli.DurationFlag{
				Name:        "valid-time",
				Usage:       "jwt toke valid time duration",
				Value:       time.Hour * 48,
				DefaultText: "48 hour",
				EnvVars:     []string{"JWT_VALID_TIME"},
				Destination: &TokenValidTime,
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
			{
				Name:  "user",
				Usage: "run user service",
				Action: func(cliCtx *cli.Context) error {
					grpcAddr := net.JoinHostPort(HostAddress, fmt.Sprintf("%d", PortGRPC))
					httpAddr := net.JoinHostPort(HostAddress, fmt.Sprintf("%d", PortHTTP))
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					grpcServer := RunUserGRPCServer(DatabaseURL, grpcAddr)
					httpServer := RunUserHTTPService(ctx, grpcAddr, httpAddr)

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
