package logging

import (
	"github.com/Totus-Floreo/otris"
	foxtris "github.com/Totus-Floreo/otris/fx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
	"os"
)

const layout = "15:04"

var Module = fx.Module("logging",
	fx.Provide(NewLogger),
)

type Logger struct {
	*slog.Logger
}

func NewLogger() Logger {
	handler := otris.NewHandlerBuilder().WithPretty().WithWriter(os.Stdout).WithTimeLayout(layout).Build()
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return Logger{logger}
}

func NewFxLogger(log Logger) fxevent.Logger {
	return foxtris.NewSlogLogger(log.Logger)
}
