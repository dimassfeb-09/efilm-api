package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/dimassfeb-09/efilm-api.git/app"
	"github.com/dimassfeb-09/efilm-api.git/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//go:embed init.sql
var schemaSQL string

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	gin.SetMode(gin.ReleaseMode)
	r.Use(middlewares.AllowCORS)

	db := app.DBConnection()
	defer db.Close()

	if _, err := db.Exec(schemaSQL); err != nil {
		log.Fatalf("Failed to initialize database schema: %s", err.Error())
	}

	r = app.InitialozedRoute(r, db)

	port := os.Getenv("APP_PORT")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Cannot run at port 8080: %s", err.Error())
	}

	log.Printf("Success run at port 8080")
}
