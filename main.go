package main

import (
	"github.com/dimassfeb-09/efilm-api.git/app"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	db, err := app.DBConnection()
	defer db.Close()

	r = app.InitialozedRoute(r, db)
	errAddress := r.Run(":8080")
	if errAddress != nil {
		log.Fatal(err)
	}

}
