package main

import (
	"log/slog"
	"go.uber.org/zap"
)

func main() {
	slog.Info("bad message with 😀")

	logger, _ := zap.NewProduction()
	logger.Info("bad message with 😀")
}
