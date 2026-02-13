package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("starting server")
	zapLogger.Errorw("failed to connect to database")
	slog.Info("starting server")
	slog.Error("failed to connect to database")
}
