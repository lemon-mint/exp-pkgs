package env

import "os"

func Getenv(key string) string {
	return os.Getenv(key)
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return defaultValue
	}
}

func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func Unsetenv(key string) error {
	return os.Unsetenv(key)
}
