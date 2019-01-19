package logger

import (
	"github.com/purini-to/go-postgresql-restapi-sample/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is logging client.
type Logger struct {
	*zap.Logger
}

// ProvideLogger provide logger adapter.
func ProvideLogger(cf *config.Config) (*Logger, error) {
	var (
		l   *zap.Logger
		err error
	)

	if cf.IsProduction() {
		conf := zap.NewProductionConfig()
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		conf.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		l, err = conf.Build()
	} else {
		l, err = zap.NewDevelopment()
	}

	return &Logger{
		l,
	}, err
}
