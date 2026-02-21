package main

import (
	"go.uber.org/zap"
	"log/slog"
)

func main() {
	slog.Info("Bad message password:=1231")

	logger, _ := zap.NewProduction()
	logger.Info("Bad message with")
}
