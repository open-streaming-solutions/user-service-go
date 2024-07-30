package database

import (
	"context"
	"fmt"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/open-streaming-solutions/user-service/internal/config"
	"github.com/open-streaming-solutions/user-service/internal/errors"
	"github.com/open-streaming-solutions/user-service/internal/logging"
	"github.com/open-streaming-solutions/user-service/internal/repository"
	"github.com/open-streaming-solutions/user-service/pkg/util"
	"go.uber.org/fx"
	"log/slog"
	"net/url"
	"os"
)

var Module = fx.Provide(NewDatabase)

// Database structure
type Database struct {
	db *pgxpool.Pool
}

// NewDatabase Saves a new database instance
func NewDatabase(lx fx.Lifecycle, logger logging.Logger, env config.Env) (*Database, repository.DBTX) {
	dbURL, err := GetDatabaseURL(env)
	if err != nil {
		logger.Error("Unable to get database URL", slog.String("error", err.Error()))
		os.Exit(1)
	}

	devDbURL, err := GetDevDatabaseURL(env)
	if err != nil {
		logger.Error("Unable to get dev database URL", slog.String("error", err.Error()))
	}

	parseConfig, err := pgxpool.ParseConfig(dbURL.String())
	if err != nil {
		logger.Error("Unable to parse config for pgx", slog.String("error", err.Error()))
		os.Exit(1)
	}
	parseConfig.ConnConfig.Tracer = logging.NewTraceLog(logger)

	conn, err := pgxpool.NewWithConfig(context.Background(), parseConfig)
	if err != nil {
		logger.Error("Unable to create connection pool", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := DoMigration(logger, dbURL, devDbURL, GetSchemaURL()); err != nil {
		logger.Error("Unable to perform database migration", slog.String("error", err.Error()))
	}

	lx.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			conn.Close()
			return nil
		},
	})

	if err := conn.Ping(context.Background()); err != nil {
		logger.Error("Unable to ping database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	db := &Database{conn}

	return db, db
}

func (db Database) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	cmd, err := db.db.Exec(ctx, sql, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.ConvertPgError(err)
	}
	return cmd, nil
}

func (db Database) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	rows, err := db.db.Query(ctx, sql, args...)
	if err != nil {
		return rows, errors.ConvertPgError(err)
	}
	return rows, nil
}

func (db Database) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return util.MockRow{Row: db.db.QueryRow(ctx, sql, args...)}
}

func (db Database) Begin(ctx context.Context) (pgx.Tx, error) {
	tx, err := db.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx}, nil
}

func GetDatabaseURL(env config.Env) (*url.URL, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	return url.Parse(dsn)
}

func GetDevDatabaseURL(env config.Env) (*url.URL, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/dev_db?sslmode=disable", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort)
	return url.Parse(dsn)
}

var pathToSchema = &url.URL{Scheme: "file", Path: "/sql/schema.sql"}

// GetSchemaURL Add in v2, get value from ENV
func GetSchemaURL() []*url.URL {
	return []*url.URL{pathToSchema}
}
