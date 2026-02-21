package main

import (
	"go.uber.org/zap"
	"log/slog"
)

func main() {
	slog.Info("Bad message with")

	logger, _ := zap.NewProduction()
	logger.Info("Bad message with")
}
