package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
type Main struct {
	Database *DatabaseConfig
	Log	  *LogConfig
}

type DatabaseConfig struct {
	Url string
}

type LogConfig struct {
	Level int
	Type  string
}

func NewMainConfig() *Main {
	return &Main{
		Database: NewDatabaseConfig(),
		Log:	  NewLogConfig(),
	}
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: String("DATABASE_URL", ""),
	}
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level: Int("LOG_LEVEL", 0),
		Type:  String("LOG_TYPE", "text"),
	}
}

func (dbc *DatabaseConfig) DbUrl() string {
	return dbc.Url
}

func Init() {
	if err := godotenv.Load(); err != nil {
		slog.Error("cannot load .env file")

		return
	}

	slog.Info("Config initialized")
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
