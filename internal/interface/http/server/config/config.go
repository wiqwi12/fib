package config

import (
	"log/slog"

	slogfiber "github.com/samber/slog-fiber"
)

type Config struct {
	Adr   string `env:"ADR_PATH"`
	Token string `env:"API_TOKEN"`
}

var LoggerCfg = slogfiber.Config{
	DefaultLevel:     slog.LevelInfo,
	ClientErrorLevel: slog.LevelWarn,
	ServerErrorLevel: slog.LevelError,
	WithRequestBody:  true,
	WithResponseBody: true,
}
