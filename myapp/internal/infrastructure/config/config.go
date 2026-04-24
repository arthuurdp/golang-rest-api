package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("arquivo .env não encontrado, usando variáveis de ambiente")
    }

    return &Config{
        DBHost:     mustGetEnv("DB_HOST"),
        DBPort:     mustGetEnv("DB_PORT"),
        DBUser:     mustGetEnv("DB_USER"),
        DBPassword: mustGetEnv("DB_PASSWORD"),
        DBName:     mustGetEnv("DB_NAME"),
        ServerPort: mustGetEnv("SERVER_PORT"),
    }
}

func mustGetEnv(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Fatalf("variável de ambiente obrigatória não definida: %s", key)
    }
    return value
}