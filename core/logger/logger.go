package logger

import (
	"go.uber.org/zap"
)

// Logger is logging client.
type Logger interface {
	Debug(string, ...zap.Field)
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	Fatal(string, ...zap.Field)
}
