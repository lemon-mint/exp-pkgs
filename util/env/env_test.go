package env

import "testing"

func TestGetEnvOrDefault(t *testing.T) {
	Unsetenv("TEST_ENV_VAR")
	if GetEnvOrDefault("TEST_ENV_VAR", "default") != "default" {
		t.Error("GetEnvOrDefault should return default value if env var is not set")
	}
	Setenv("TEST_ENV_VAR", "value")
	if GetEnvOrDefault("TEST_ENV_VAR", "default") != "value" {
		t.Error("GetEnvOrDefault should return env var value if env var is set")
	}

	Unsetenv("TEST_ENV_VAR")
}

func TestGetEnv(t *testing.T) {
	Unsetenv("TEST_ENV_VAR")
	if Getenv("TEST_ENV_VAR") != "" {
		t.Error("Getenv should return empty string if env var is not set")
	}
	Setenv("TEST_ENV_VAR", "value")
	if Getenv("TEST_ENV_VAR") != "value" {
		t.Error("Getenv should return env var value if env var is set")
	}

	Unsetenv("TEST_ENV_VAR")
}
