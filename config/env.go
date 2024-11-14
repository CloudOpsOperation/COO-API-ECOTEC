package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	ApiPort    string
	DBUser     string
	DBPwd      string
	DBName     string
	DBAddr     string
}

var Envs = NewConfig()

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		ApiPort:    getEnv("API_PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPwd:      getEnv("DB_PWD", "root"),
		DBName:     getEnv("DB_NAME", "ecotec"),
		DBAddr:     getEnv("DB_ADDR", "localhost:3306"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
