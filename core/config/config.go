package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	defaultConfFile = "config/app.yaml"
	modeKey         = "MODE"
)

// Config is application of configuration.
type Config struct {
	*viper.Viper
}

// IsProduction return true if application is production.
func (c *Config) IsProduction() bool {
	return c.GetString(modeKey) == "production"
}

// NewConfig create config.
func NewConfig() (*Config, error) {
	v := viper.New()
	file := os.Getenv("CONFIG_FILE")
	if file == "" {
		file = defaultConfFile
	}

	v.SetDefault("PORT", "8080")
	v.SetDefault(modeKey, "development")

	v.SetConfigFile(file)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Config{
		v,
	}, nil
}
