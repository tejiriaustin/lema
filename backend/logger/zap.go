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

func (z *zapLogger) convertToZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

func (z *zapLogger) Info(msg string, fields ...Field) {
	z.log.Info(msg, z.convertToZapFields(fields)...)
}

func (z *zapLogger) Error(msg string, fields ...Field) {
	z.log.Error(msg, z.convertToZapFields(fields)...)
}

func (z *zapLogger) Debug(msg string, fields ...Field) {
	z.log.Debug(msg, z.convertToZapFields(fields)...)
}

func (z *zapLogger) Warn(msg string, fields ...Field) {
	z.log.Warn(msg, z.convertToZapFields(fields)...)
}

func (z *zapLogger) Fatal(msg string, fields ...Field) {
	z.log.Fatal(msg, z.convertToZapFields(fields)...)
}

func (z *zapLogger) Sync() error {
	return z.log.Sync()
}
