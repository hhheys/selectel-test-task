package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("server started")
	zapLogger.Errorw("connection failed")
	zapLogger.Warnw("something went wrong")

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
}
