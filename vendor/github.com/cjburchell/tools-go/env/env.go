package env

import (
	"os"
	"strconv"
)

// GetInt from environment variables
func GetInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.Atoi(value); err == nil {
			return result
		}
		return fallback
	}
	return fallback
}

// GetInt64 from environment variables
func GetInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseInt(value, 10, 64); err == nil {
			return result
		}
		return fallback
	}
	return fallback
}

// Get string from environment variables
func Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetBool from environment variables
func GetBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseBool(value); err == nil {
			return result
		}
	}
	return fallback
}
