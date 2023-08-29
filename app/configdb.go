package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func DBConnection() *sql.DB {

	env := GetEnv()

	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
		env.DBSSLMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to connect DB" + err.Error())
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	log.Println("Successfully connected DB!")
	return db
}
