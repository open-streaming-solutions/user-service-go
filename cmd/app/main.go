package main

import (
	"context"
	"fmt"
	"github.com/Open-Streaming-Solutions/user-service/internal/config"
	"github.com/Open-Streaming-Solutions/user-service/internal/dbContext"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/repository/db"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logging.Module,
		config.Module,
		dbContext.Module,
		db.Module,
		// TODO Implement the following
		//fx.WithLogger(func() fxevent.Logger {
		//	return logging.GetFxLogger()
		//}),
	)

	ctx := context.Background()

	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		// TODO Implement the following
		//logger.Fatal(err)
	}
	fmt.Println("Hello, World!")
}
