package log

import (
	"log/slog"
	"os"
)

const (
	ENV_LOCAL = "local"
	ENV_PROD  = "prod"
)

var log *slog.Logger

type Logger struct {
	*slog.Logger
}

func GetLogger() Logger {
	return Logger{log}
}

func MustSetup(env string) *slog.Logger {
	switch env {
	case ENV_LOCAL:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case ENV_PROD:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	default:
		panic("unknown env")
	}

	return log
}
