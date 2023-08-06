package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	DBHost    string
	DBName    string
	DBPass    string
	DBPort    string
	DBUser    string
	DBSSLMode string
}

func getEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{
		DBHost:    os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_NAME"),
		DBPass:    os.Getenv("DB_PASS"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBSSLMode: os.Getenv("DB_SSL_MODE"),
	}
}
