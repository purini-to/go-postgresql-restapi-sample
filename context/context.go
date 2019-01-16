package context

import "os"

var (
	// ModeKey is environment key.
	ModeKey = "GIN_MODE"
	// ModePrd is production mode.
	ModePrd = "release"
)

// IsProduction is application production.
func IsProduction() bool {
	return os.Getenv(ModeKey) == ModePrd
}
