package logging

import (
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Provide(NewLogger)

type Logger struct {
	*slog.Logger
}

// TODO Set up slog handler https://github.com/golang/example/blob/master/slog-handler-guide/README.md
func NewLogger() Logger {
	return Logger{
		Logger: slog.Default(),
	}
}
