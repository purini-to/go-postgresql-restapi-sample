package logger

import (
	"go.uber.org/zap"
)

// ProvideLogger provide logger adapter.
func ProvideLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
