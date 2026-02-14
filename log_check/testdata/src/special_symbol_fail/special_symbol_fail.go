package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("server started! 🚀 ")               // want "log message should not contain special symbols or emoji"
	zapLogger.Errorw("connection failed!!!")            // want "log message should not contain special symbols or emoji"
	zapLogger.Warnw("warning: something went wrong...") // want "log message should not contain special symbols or emoji"

	slog.Info("server started! 🚀 ")               // want "log message should not contain special symbols or emoji"
	slog.Error("connection failed!!!")            // want "log message should not contain special symbols or emoji"
	slog.Warn("warning: something went wrong...") // want "log message should not contain special symbols or emoji"
}
