package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	log *zap.Logger
}

var _ Logger = (*zapLogger)(nil)

func NewProductionLogger() (Logger, error) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.StacktraceKey = "stacktrace"
	config.EncoderConfig.MessageKey = "message"

	logger, err := config.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}

	return &zapLogger{log: logger}, nil
}

func (l *zapLogger) Info(msg string, fields ...zap.Field) {
	l.log.Info(msg, fields...)
}

func (l *zapLogger) Error(msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	l.log.Error(msg, fields...)
}

func (l *zapLogger) Debug(msg string, fields ...zap.Field) {
	l.log.Debug(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	l.log.Fatal(msg, fields...)
}

func (l *zapLogger) Sync() error {
	return l.log.Sync()
}
