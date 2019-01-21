package config

import (
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
	"github.com/spf13/viper"
)

var (
	defaultConfFile = "config/app.yaml"
)

// NewConfig create config.
func NewConfig(l logger.Logger) *viper.Viper {
	v := viper.New()
	file := os.Getenv("CONFIG_FILE")
	if file == "" {
		file = defaultConfFile
	}

	v.SetDefault("PORT", "8080")
	v.SetConfigFile(file)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		l.Warn(err.Error(), zap.String("file", file))
	}
	return v
}
