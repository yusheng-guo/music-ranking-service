package log

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

// NewLogger 新建 Logger 结构体
func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	return &Logger{Logger: logger}, err
}

func (l *Logger) INFO(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func (l *Logger) WARN(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func (l *Logger) ERROR(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func (l *Logger) DEBUG(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}
