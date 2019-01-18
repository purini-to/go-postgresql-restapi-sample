package config

import (
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

// ProvideConfig provide config.
func ProvideConfig(lg *zap.Logger) (*viper.Viper, error) {
	v := viper.New()
	file := os.Getenv("CONFIG_FILE")
	if file == "" {
		file = "config/application.yaml"
	}

	v.SetDefault("PORT", "8080")

	v.SetConfigFile(file)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		lg.Warn(err.Error(), zap.String("file", file))
	}
	return v, nil
}
