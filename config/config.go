package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

type Config struct {
	AppKey      string
	AppURL      string
	MainAppPort string
	DB          DBConfig
	Redis       RedisConfig
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Файл .env не найден, используются системные переменные")
	}

	return Config{
		AppKey:      getEnv("APP_KEY", ""),
		AppURL:      getEnv("APP_URL", "http://localhost"),
		MainAppPort: getEnv("MAIN_APP_PORT", "8081"),
		DB:          loadDBConfig(),
		Redis:       loadRedisConfig(),
	}
}

func loadDBConfig() DBConfig {
	return DBConfig{
		Type:     getEnv("DB_CONNECTION", "postgres"),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "quantum"),
		Password: getEnv("DB_PASSWORD", "quantum"),
		Name:     getEnv("DB_DATABASE", "quantum"),
	}
}

func loadRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     getEnv("REDIS_PORT", "6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
