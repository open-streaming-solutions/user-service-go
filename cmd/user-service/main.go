package main

import (
	"github.com/open-streaming-solutions/user-service/internal/config"
	"github.com/open-streaming-solutions/user-service/internal/database"
	"github.com/open-streaming-solutions/user-service/internal/handler"
	"github.com/open-streaming-solutions/user-service/internal/logging"
	"github.com/open-streaming-solutions/user-service/internal/repository"
	"github.com/open-streaming-solutions/user-service/internal/server"
	"github.com/open-streaming-solutions/user-service/internal/service"
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
