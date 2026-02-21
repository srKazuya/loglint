package main

import (
	"log/slog"
	"go.uber.org/zap"
)

func main() {
	slog.Info("bad message with привет {}")

	logger, _ := zap.NewProduction()
	logger.Info("привет")
}
