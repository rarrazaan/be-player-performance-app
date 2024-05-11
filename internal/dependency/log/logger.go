package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logLevel string) *Logger {
	zapConfig := zap.NewProductionConfig()

	switch logLevel {
	case "debug":
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	case "info":
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	case "warn":
		zapConfig.Level.SetLevel(zapcore.WarnLevel)
	case "error":
		zapConfig.Level.SetLevel(zapcore.ErrorLevel)
	case "fatal":
		zapConfig.Level.SetLevel(zapcore.FatalLevel)
	case "panic":
		zapConfig.Level.SetLevel(zapcore.PanicLevel)
	default:
		zapConfig.Level.SetLevel(zapcore.InfoLevel)
	}
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	zapLogger, err := zapConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to instantiate logger, err: %s", err))
	}

	return &Logger{logger: zapLogger.Sugar()}
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) ErrorWithFields(msg string, kv ...any) {
	l.logger.Errorw(msg, kv...)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) InfoWithFields(msg string, kv ...any) {
	l.logger.Infow(msg, kv...)
}
