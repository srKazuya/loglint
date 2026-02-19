package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("bad message with 🤫")
	slog.Info("bad message with привет")
	slog.Info("Bad message with")
	slog.Info("bad message with: password")

	logger, _ := zap.NewProduction()
	logger.Info("Hello")
}
