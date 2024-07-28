package server

import (
	"context"
	"github.com/Open-Streaming-Solutions/user-service/internal/config"
	"github.com/Open-Streaming-Solutions/user-service/internal/handler"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/middleware"
	mdlogging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
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

func NewGrpcServer(lx fx.Lifecycle, logger logging.Logger, env config.Env, handle *handler.GrpcHandler) *GrpcServer {
	opts := []mdlogging.Option{
		mdlogging.WithLogOnEvents(mdlogging.StartCall, mdlogging.FinishCall),
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			mdlogging.UnaryServerInterceptor(logging.InterceptorLogger(logger), opts...),
			recovery.UnaryServerInterceptor(
				recovery.WithRecoveryHandler(middleware.RecoveryHandlerFunc),
			),
		),
	)
	handler.RegisterGrpcHandler(server, handle)

	lx.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listen, err := net.Listen("tcp", ":"+env.Port)
			if err != nil {
				return err
			}

			go func() {
				if err := server.Serve(listen); err != nil {
					logger.Error("Failed to start server", "error", err)
				}
			}()

			logger.Info("grpc server listening on ", slog.String("addr", listen.Addr().String()))

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
