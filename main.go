package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/dimassfeb-09/efilm-api.git/app"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	gin.SetMode(gin.ReleaseMode)
	r.Use(app.AllowCORS)

	godotenv.Load()

	db := app.DBConnection()
	defer db.Close()

	r = app.InitialozedRoute(r, db)

	port := os.Getenv("APP_PORT")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Cannot run at port 8080: %s", err.Error())
	}

	log.Printf("Success run at port 8080")
}
