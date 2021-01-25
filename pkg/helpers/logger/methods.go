package logger

func (l *LoggerImpl) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LoggerImpl) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *LoggerImpl) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *LoggerImpl) Sync() error {
	return l.logger.Sync()
}

func (l *LoggerImpl) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *LoggerImpl) Info(args ...interface{}) {
	l.logger.Info(args...)
}