package env

import (
	"os"
)

var (
	modeKey = "MODE"
)

// IsProduction check production mode.
func IsProduction() bool {
	return os.Getenv(modeKey) == "production"
}
