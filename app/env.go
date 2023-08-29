package app

import (
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

func GetEnv() *Env {

	return &Env{
		DBHost:    os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_NAME"),
		DBPass:    os.Getenv("DB_PASS"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBSSLMode: os.Getenv("DB_SSL_MODE"),
	}
}
