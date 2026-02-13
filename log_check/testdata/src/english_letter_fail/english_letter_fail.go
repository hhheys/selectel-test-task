package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("запуск сервера")                    // want "log message should not contain non-English letters"
	zapLogger.Errorw("ошибка подключения к базе данных") // want "log message should not contain non-English letters"
	slog.Info("запуск сервера")                          // want "log message should not contain non-English letters"
	slog.Error("ошибка подключения к базе данных")       // want "log message should not contain non-English letters"
}
