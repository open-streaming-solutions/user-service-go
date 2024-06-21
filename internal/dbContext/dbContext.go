package dbContext

import (
	"fmt"
	"github.com/Open-Streaming-Solutions/user-service/internal/config"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var Module = fx.Provide(NewDatabase)

// Database structure
type Database struct {
	*gorm.DB
}

// NewDatabase Saves a new database instance
func NewDatabase(env config.Env, logger logging.Logger) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		logger.Error("Open database error", err)
	}

	// Get *sql.DB object from GORM connection
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Get db object from GORM error", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(env.DBMaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(env.DBMaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(env.DBMaxLifetime) * time.Second)

	logger.Info("Database connection established")

	if err := DoMigrations(db); err != nil {
		logger.Error("Do migrations error", err)
	}

	return Database{
		DB: db,
	}
}

func DoMigrations(db *gorm.DB) error {
	return db.AutoMigrate()
}
