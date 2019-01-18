package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProvideLogger provide logger adapter.
func ProvideLogger() (*zap.Logger, error) {
	if false {
		conf := zap.NewProductionConfig()
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		conf.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		return conf.Build()
	}
	return zap.NewDevelopment()
}
