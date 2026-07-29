package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port string 
	DatabaseURL string

	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime int
	DBConnMaxIdleTime int
}

func Load() *Config{
	cfg := &Config{
		Port: getEnv("PORT","8080"),
		DatabaseURL: getEnv("DATABASE_URL",""),

		DBMaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		DBConnMaxLifetime: getEnvAsInt("DB_CONN_MAX_LIFETIME", 30),
		DBConnMaxIdleTime: getEnvAsInt("DB_CONN_MAX_IDLE_TIME", 5),

	}
	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intValue
}