package logger

import (
	"log/slog"
	"os"

	"homework-fiber/config"
)

type Service struct {
	Info  *slog.Logger
	Error *slog.Logger
}

func NewService(lc *config.LogConfig) *Service {
	if lc.Type == "json" {
		return &Service{
			Info:  slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(lc.Level)})),
			Error: slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(lc.Level)})),
		}
	}

	return &Service{
		Info:  slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(lc.Level)})),
		Error: slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(lc.Level)})),
	}

}
