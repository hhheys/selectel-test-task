package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("starting server on port 8080")
	zapLogger.Infow("failed to connect to database")
	slog.Info("starting server on port 8080")
	slog.Info("failed to connect to database")
}
