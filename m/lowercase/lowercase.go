package main

import (
	"log/slog"
	"go.uber.org/zap"
)

func main() {
	slog.Info("Bad message with")

	logger, _ := zap.NewProduction()
	logger.Info("Bad message with")
}
