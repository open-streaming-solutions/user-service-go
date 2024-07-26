package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log/slog"
)

const EnvPrefix = "user_service"

var Module = fx.Provide(NewEnv)

type Env struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Port       string
}

func NewEnv() (env Env) {
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	env.DBUsername = viper.GetString("DB_USERNAME")
	env.DBPassword = viper.GetString("DB_PASSWORD")
	env.DBHost = viper.GetString("DB_HOST")
	env.DBPort = viper.GetString("DB_PORT")
	env.DBName = viper.GetString("DB_NAME")
	env.Port = viper.GetString("PORT")

	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Info("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	return env
}
