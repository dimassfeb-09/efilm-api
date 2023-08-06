package app

import (
	"database/sql"
	"github.com/dimassfeb-09/efilm-api.git/controller"
	"github.com/dimassfeb-09/efilm-api.git/repository"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
)

func InitialozedRoute(r *gin.Engine, db *sql.DB) *gin.Engine {

	api := r.Group("/api")

	actorRepository := repository.NewActorRepository()
	actorService := services.NewActorService(db, actorRepository)
	actorController := controller.NewActorControllerImpl(actorService)

	api.POST("/actors", actorController.Save)
	api.GET("/actors", actorController.FindAll)
	api.GET("/actors/search", actorController.FindBySearch)
	api.GET("/actors/:id", actorController.FindByID)
	api.PUT("/actors/:id", actorController.Update)
	api.DELETE("/actors/:id", actorController.Delete)

	nationalRepository := repository.NewNationalRepository()
	nationalService := services.NewNationalService(db, nationalRepository)
	nationalController := controller.NewNationalControllerImpl(nationalService)

	api.POST("/nationals", nationalController.Save)
	api.GET("/nationals", nationalController.FindAll)
	api.GET("/nationals/search", nationalController.FindBySearch)
	api.GET("/nationals/:id", nationalController.FindByID)
	api.PUT("/nationals/:id", nationalController.Update)
	api.DELETE("/nationals/:id", nationalController.Delete)

	return r
}
