package logger

import (
	"github.com/purini-to/go-postgresql-restapi-sample/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProvideLogger provide logger adapter.
func ProvideLogger() (*zap.Logger, error) {
	if context.IsProduction() {
		conf := zap.NewProductionConfig()
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		conf.EncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
		return conf.Build()
	}
	return zap.NewDevelopment()
}
