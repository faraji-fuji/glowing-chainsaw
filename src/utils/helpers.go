package helpers

import "os"

// GetEnv returns the value of an environment variable, or a default value if not found.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
