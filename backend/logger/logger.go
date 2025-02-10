package logger

import "go.uber.org/zap"

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, err error, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Fatal(msg string, err error, fields ...zap.Field)
	Sync() error
}
