package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	password := "123456"
	apiKey := "123456"
	token := "123456"

	zapLogger.Infow("user password: " + password) // want "log message should not contain sensitive data"
	zapLogger.Debug("api_key=" + apiKey)          // want "log message should not contain sensitive data"
	zapLogger.Infow("token: " + token)            // want "log message should not contain sensitive data"

	slog.Info("user password: " + password) // want "log message should not contain sensitive data"
	slog.Debug("api_key=" + apiKey)         // want "log message should not contain sensitive data"
	slog.Info("token: " + token)            // want "log message should not contain sensitive data"
}
