package main

import (
	"github.com/Open-Streaming-Solutions/user-service/internal/config"
	"github.com/Open-Streaming-Solutions/user-service/internal/database"
	"github.com/Open-Streaming-Solutions/user-service/internal/handler"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/repository"
	"github.com/Open-Streaming-Solutions/user-service/internal/server"
	"github.com/Open-Streaming-Solutions/user-service/internal/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.WithLogger(logging.NewFxLogger),
		config.Module,
		logging.Module,
		database.Module,
		repository.Module,
		service.Module,
		handler.Module,
		server.Module,
	)

	app.Run()
}
