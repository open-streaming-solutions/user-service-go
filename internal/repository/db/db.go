package db

import (
	"github.com/Open-Streaming-Solutions/user-service/internal/dbContext"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewDatabase)

type Database interface {
}

type database struct {
	log logging.Logger
	db  dbContext.Database
}

func NewDatabase(log logging.Logger, db dbContext.Database) Database {
	return &database{
		log: log,
		db:  db,
	}
}
