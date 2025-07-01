package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}

	log.Println("Config initialized")
}

func String(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func Int(key string, defaultValue int) int {
	v := os.Getenv(key)

	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return i
}

func Bool(key string, defaultValue bool) bool {
	v := os.Getenv(key)
	b, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}
	return b

}

type DatabaseConfig struct {
	url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: String("DATABASE_URL", ""),
	}
}
