package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("user authenticated successfully")
	zapLogger.Debug("api request completed")
	zapLogger.Infow("token validated")

	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")
}
