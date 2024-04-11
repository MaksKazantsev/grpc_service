package app

import (
	"github.com/MaksKazantsev/grpc_service/user/internal/config"
	"github.com/MaksKazantsev/grpc_service/user/internal/log"
	"log/slog"
)

func MustStart(cfg *config.Config) {
	l := log.MustSetup(cfg.Env)

	l.Info("starting app with", slog.Any("cfg:", cfg))
}
