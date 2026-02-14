package lowercasefail

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// Мок ZAP логгера
	var zapLogger *zap.SugaredLogger
	_ = zapLogger

	zapLogger.Infow("Starting server on port 8080")  // want "log message should start with a lowercase letter"
	zapLogger.Infow("Failed to connect to database") // want "log message should start with a lowercase letter"
	slog.Info("Starting server on port 8080")        // want "log message should start with a lowercase letter"
	slog.Info("Failed to connect to database")       // want "log message should start with a lowercase letter"
}
