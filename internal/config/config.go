package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Domain      string
	Port        string
	Env         string
}

func Load() (*Config, error) {
	_ = godotenv.Load

	cfg := &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:password@localhost/urlshortener?sslmode=disable"),
		Domain:      getEnv("DOMAIN", "http://localhost:8080"),
		Port:        getEnv("PORT", "8080"),
		Env:         getEnv("ENV", "development"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
