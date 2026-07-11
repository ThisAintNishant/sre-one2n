package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found, using environment variables")
	}

	return &Config{
		Port: getEnv("PORT", "8080"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "admin"),
		DBPass: getEnv("DB_PASSWORD", "password"),
		DBName: getEnv("DB_NAME", "srebootcamp"),
	}
}

func getEnv(key string, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}