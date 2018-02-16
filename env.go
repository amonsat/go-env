package env

import (
	"os"
	"strconv"
)

func GetString(key, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return value
}

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

func GetInt(key string, def int64) int64 {
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

func GetFloat(key string, def float64) float64 {
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
