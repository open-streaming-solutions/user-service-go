package serve

import (
	"context"
	"github.com/Open-Streaming-Solutions/user-service/internal/config"
	"github.com/Open-Streaming-Solutions/user-service/internal/controller/handler"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

var Module = fx.Invoke(NewGrpcServer)

type GrpcServer struct {
	logger logging.Logger
	server *grpc.Server
	env    config.Env
}

func NewGrpcServer(logger logging.Logger, env config.Env, lc fx.Lifecycle) *GrpcServer {
	server := grpc.NewServer()
	handler.RegisterGrpcHandler(server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listen, err := net.Listen("tcp", env.Port)
			if err != nil {
				return err
			}

			go func() {
				if err := server.Serve(listen); err != nil {
					logger.Error("Failed to start server", "error", err)
				}
			}()

			logger.Info("grpc server listening on ", slog.String("addr", listen.Addr().String()), slog.String("port", env.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})

	return &GrpcServer{
		logger: logger,
		env:    env,
		server: server,
	}
}
