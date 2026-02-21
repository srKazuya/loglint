package main

import (
	"log"
	"log/slog"
	"go.uber.org/zap"
)

func main() {
	slog.Info("bad message with password")
	log.Fatal("password")
	
	logger, _ := zap.NewProduction()
	logger.Info("bad message with password")
}
