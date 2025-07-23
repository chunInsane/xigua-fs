package config

import (
	"os"
)

type Config struct {
	Port         string
	DatabasePath string
	StoragePath  string
	JWTSecret    string
	MaxFileSize  int64
}

func New() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		DatabasePath: getEnv("DATABASE_PATH", "./storage/xigua.db"),
		StoragePath:  getEnv("STORAGE_PATH", "./storage/files"),
		JWTSecret:    getEnv("JWT_SECRET", "******************"),
		MaxFileSize:  100 * 1024 * 1024, // 100MB
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
