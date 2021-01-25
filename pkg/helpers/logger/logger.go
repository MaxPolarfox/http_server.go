package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

//Logger is our contract for the logger
type Logger interface {
	Debug(args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Sync() error
}

type LoggerImpl struct {
	logger *zap.SugaredLogger
}

func NewLogger(out io.Writer, config Options) Logger {
	baseLogger := zap.New(
		zapcore.NewCore(
			getEncoder(config),
			zapcore.Lock(zapcore.AddSync(out)),
			zap.NewAtomicLevelAt(getZapLevel(config.Level)),
		),
	)

	sugarLogger := baseLogger.Sugar()
	return &LoggerImpl{
		logger: sugarLogger,
	}
}

func getZapLevel(configLevel string) zapcore.Level {
	switch configLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getEncoder(config Options) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	if config.Json {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}