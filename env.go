package env

import (
	"os"
	"strconv"
	"time"
)

// GetString - get string environment variable or default value
func GetString(key, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return value
}

// GetBool - get bool environment variable or default value
func GetBool(key string, def bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	b, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}

	return b
}

// GetInt - get int environment variable or default value
func GetInt(key string, def int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}

	return i
}

// GetInt64 - get int64 environment variable or default value
func GetInt64(key string, def int64) int64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}

	return i
}

// GetFloat32 - get float environment variable or default value
func GetFloat32(key string, def float32) float32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return def
	}

	return float32(f)
}

// GetFloat64 - get float64 environment variable or default value
func GetFloat64(key string, def float64) float64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return def
	}

	return f
}

// GetUint - get uint environment variable or default value
func GetUint(key string, def uint) uint {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return def
	}

	return uint(i)
}

// GetUint64 - get uint64 environment variable or default value
func GetUint64(key string, def uint64) uint64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return def
	}

	return i
}

// GetDuration - get time.Duration environment variable or default value
func GetDuration(key string, def time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		return def
	}

	return d
}
