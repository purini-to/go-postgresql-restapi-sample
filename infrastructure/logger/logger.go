package logger

import (
	"github.com/purini-to/go-postgresql-restapi-sample/core/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger create logger adapter.
func NewLogger() (*zap.Logger, error) {
	var (
		l   *zap.Logger
		err error
	)
	if env.IsProduction() {
		conf := zap.NewProductionConfig()
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		conf.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		l, err = conf.Build()
	} else {
		l, err = zap.NewDevelopment()
	}

	return l, err
}
