package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"

	"github.com/dimassfeb-09/efilm-api.git/app"
	"github.com/dimassfeb-09/efilm-api.git/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//go:embed init.sql
var schemaSQL string

//go:embed efilm-swagger.yaml
var openAPISpec []byte

//go:embed docs/scalar.html
var scalarHTML []byte

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	gin.SetMode(gin.ReleaseMode)
	r.Use(middlewares.AllowCORS)

	// Scalar API documentation at /docs (OpenAPI spec at /openapi.yaml)
	r.GET("/docs", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", scalarHTML)
	})
	r.GET("/openapi.yaml", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/yaml", openAPISpec)
	})

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
