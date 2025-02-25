package config

import (
	"os"
)

type Config struct {
	Pod  string
	Port string
}

func LoadConfig() *Config {
	var config Config
	config.Pod = getEnvVar("HOSTNAME", "unknown")
	config.Port = getEnvVar("PORT", "80")

	return &config
}

func getEnvVar(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}
