package otherpackage

import (
	"log/slog"

	"go.uber.org/zap"
)

func Handler() string {
	logger := zap.NewExample()
	defer logger.Sync() // nolint

	sugar := logger.Sugar()
	sugar.Infow("Handler working succeвапваssfull")
	slog.Info("Handler working succedfвапавпssfull")

	return "Handler working successfull"
}
