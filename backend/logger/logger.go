package logger

type (
	Field struct {
		Key   string
		Value interface{}
	}

	Logger interface {
		Info(msg string, fields ...Field)
		Error(msg string, fields ...Field)
		Debug(msg string, fields ...Field)
		Warn(msg string, fields ...Field)
		Fatal(msg string, fields ...Field)
		Sync() error
	}
)

func WithField(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}
