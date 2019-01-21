package config

// Config is application of configuration.
type Config interface {
	GetString(string) string
}
