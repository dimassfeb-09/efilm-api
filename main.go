package main

import (
	"github.com/dimassfeb-09/efilm-api.git/app"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

//	@title			Swagger API E-Film
//	@version		1.0
//	@description	This is a sample server film
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
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
