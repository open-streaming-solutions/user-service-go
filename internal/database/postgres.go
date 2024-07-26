package database

import (
	"context"
	"fmt"
	"github.com/Open-Streaming-Solutions/user-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"log/slog"
	"os"
)

var Module = fx.Provide(NewDatabase)

// Database structure
type Database struct {
	DB *pgxpool.Pool
}

// NewDatabase Saves a new database instance
func NewDatabase(env config.Env) Database {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		slog.Error("Unable to create connection pool", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		slog.Error("Unable to ping database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return Database{
		DB: dbpool,
	}
}
