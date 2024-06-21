package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const EnvPrefix = "user_service"

var Module = fx.Provide(NewEnv)

type Env struct {
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBMaxLifetime  int64
}

func NewEnv() (env Env) {

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	env.DBUsername = viper.GetString("DB_USERNAME")
	env.DBPassword = viper.GetString("DB_PASSWORD")
	env.DBHost = viper.GetString("DB_HOST")
	env.DBPort = viper.GetString("DB_PORT")
	env.DBName = viper.GetString("DB_NAME")
	env.DBMaxIdleConns = viper.GetInt("DB_MAX_IDLE_CONNS")
	env.DBMaxOpenConns = viper.GetInt("DB_MAX_OPEN_CONNS")
	env.DBMaxLifetime = viper.GetInt64("DB_MAX_LIFETIME")

	return env
}
