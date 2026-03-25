package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string
}

func NewConfig() (*Config, error) {
	// Attempt to load .env file; it's okay if it doesn't exist (e.g., in docker/production)
	_ = godotenv.Load()

	config := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=urlshortener port=5432 sslmode=disable"),
		Env:         getEnv("ENV", "development"),
	}

	return config, nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
