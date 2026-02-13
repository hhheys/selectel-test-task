package zap

type SugaredLogger struct{}

func (l *SugaredLogger) Infow(msg string, keysAndValues ...interface{})  {}
func (l *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {}
func (l *SugaredLogger) Warnw(msg string, keysAndValues ...interface{})  {}
func (l *SugaredLogger) Debug(msg string, keysAndValues ...interface{})  {}
