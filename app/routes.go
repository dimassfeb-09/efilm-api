package app

import (
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/efilm-api.git/controller"
	"github.com/dimassfeb-09/efilm-api.git/middlewares"
	"github.com/dimassfeb-09/efilm-api.git/repository"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
)

func InitialozedRoute(r *gin.Engine, db *sql.DB) *gin.Engine {

	// Index
	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("Server ON!"))
	})

	api := r.Group("/api")

	authRepository := repository.NewAuthRepository()
	authService := services.NewAuthService(db, authRepository)
	authController := controller.NewAuthControllerImpl(authService)

	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)

	actorRepository := repository.NewActorRepository()
	actorService := services.NewActorService(db, actorRepository)
	actorController := controller.NewActorControllerImpl(actorService)

	userController := controller.NewUserController()
	api.POST("/users/info", userController.GetUserInfo)

	api.Use(middlewares.MiddlewareToken)

	api.POST("/actors", actorController.Save)
	api.GET("/actors", actorController.FindAll)
	api.GET("/actors/search", actorController.FindBySearch)
	api.GET("/actors/:id", actorController.FindByID)
	api.PUT("/actors/:id", actorController.Update)
	api.DELETE("/actors/:id", actorController.Delete)

	directorRepository := repository.NewDirectorRepository()
	directorService := services.NewDirectorService(db, directorRepository)
	directorController := controller.NewDirectorControllerImpl(directorService)

	api.POST("/directors", directorController.Save)
	api.GET("/directors", directorController.FindAll)
	api.GET("/directors/search", directorController.FindBySearch)
	api.GET("/directors/:id", directorController.FindByID)
	api.PUT("/directors/:id", directorController.Update)
	api.DELETE("/directors/:id", directorController.Delete)

	nationalRepository := repository.NewNationalRepository()
	nationalService := services.NewNationalService(db, nationalRepository)
	nationalController := controller.NewNationalControllerImpl(nationalService)

	api.POST("/nationals", nationalController.Save)
	api.GET("/nationals", nationalController.FindAll)
	api.GET("/nationals/search", nationalController.FindBySearch)
	api.GET("/nationals/:id", nationalController.FindByID)
	api.PUT("/nationals/:id", nationalController.Update)
	api.DELETE("/nationals/:id", nationalController.Delete)

	movieRepository := repository.NewMovieRepository()
	movieService := services.NewMovieService(db, movieRepository)
	movieController := controller.NewMovieControllerImpl(movieService)

	api.POST("/movies/:movie_id/upload_poster", movieController.UploadPoster)
	api.POST("/movies", movieController.Save)
	api.GET("/movies", movieController.FindAll)
	api.GET("/movies/search", movieController.FindBySearch)
	api.GET("/movies/:movie_id", movieController.FindByID)
	api.PUT("/movies/:movie_id", movieController.Update)
	api.DELETE("/movies/:movie_id", movieController.Delete)

	recommendationMovieRepo := repository.NewRecommendationMovieRepositoryImpl()
	recommendationMovieService := services.NewRecommendationMovieService(db, recommendationMovieRepo)
	recommendationMovieController := controller.NewRecommendationMovieControllerImpl(recommendationMovieService)

	api.GET("/movies/recommendation", recommendationMovieController.FindAll)
	api.POST("/movies/recommendation", recommendationMovieController.Save)
	api.DELETE("/movies/recommendation/:movie_id", recommendationMovieController.Delete)

	genreRepository := repository.NewGenreRepository()
	genreService := services.NewGenreService(db, genreRepository, movieService)
	genreController := controller.NewGenreControllerImpl(genreService)

	api.POST("/genres", genreController.Save)
	api.GET("/genres", genreController.FindAll)
	api.GET("/genres/search", genreController.FindBySearch)
	api.GET("/genres/:id", genreController.FindByID)
	api.GET("/genres/:id/movies", genreController.FindAllMoviesByID)
	api.PUT("/genres/:id", genreController.Update)
	api.DELETE("/genres/:id", genreController.Delete)

	movieActorsRepository := repository.NewMovieActorRepository()
	movieActorsService := services.NewMovieActorService(db, movieActorsRepository)
	movieActorsController := controller.NewMovieActorControllerImpl(movieActorsService)

	api.POST("/movies/:movie_id/actors", movieActorsController.Save)
	api.GET("/movies/:movie_id/actors", movieActorsController.FindByID)
	api.PUT("/movies/:movie_id/actors/:actor_id", movieActorsController.Update)
	api.DELETE("/movies/:movie_id/actors/:actor_id", movieActorsController.Delete)

	movieDirectorsRepository := repository.NewMovieDirectorRepository()
	movieDirectorsService := services.NewMovieDirectorService(db, movieDirectorsRepository)
	movieDirectorsController := controller.NewMovieDirectorControllerImpl(movieDirectorsService)

	api.POST("/movies/:movie_id/directors", movieDirectorsController.Save)
	api.GET("/movies/:movie_id/directors", movieDirectorsController.FindByID)
	api.DELETE("/movies/:movie_id/directors/:director_id", movieDirectorsController.Delete)

	movieGenresRepository := repository.NewMovieGenreRepository()
	movieGenresService := services.NewMovieGenreService(db, movieGenresRepository)
	movieGenresController := controller.NewMovieGenreControllerImpl(movieGenresService)

	api.POST("/movies/:movie_id/genres", movieGenresController.Save)
	api.GET("/movies/:movie_id/genres", movieGenresController.FindByID)
	api.DELETE("/movies/:movie_id/genres/:genre_id", movieGenresController.Delete)

	return r
}
