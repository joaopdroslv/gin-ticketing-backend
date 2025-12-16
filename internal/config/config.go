package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort             string
	DockerDatabaseURL    string
	LocalhostDatabaseURL string
}

func Load() *Config {
	LoadEnvFile()

	return &Config{
		HTTPPort:             getEnv("HTTP_PORT", ":8080"),
		DockerDatabaseURL:    getEnv("DOCKER_DATABASE_URL", ""),
		LocalhostDatabaseURL: getEnv("LOCALHOST_DATABASE_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func LoadEnvFile() {
	_ = godotenv.Load()
}
