package controller

import (
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieGenreController interface {
	Save(gc *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type MovieGenreControllerImpl struct {
	MovieGenreService services.MovieGenreService
}

func NewMovieGenreControllerImpl(genreService services.MovieGenreService) MovieGenreController {
	return &MovieGenreControllerImpl{MovieGenreService: genreService}
}

func (controller *MovieGenreControllerImpl) Save(gc *gin.Context) {

	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format movie_id",
		})
		return
	}

	var movieActor web.MovieGenreModelRequestPost
	err = gc.ShouldBind(&movieActor)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	movieActor.MovieID = movieID

	err = controller.MovieGenreService.Save(gc.Request.Context(), &movieActor)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	gc.JSON(http.StatusOK, web.ResponseError{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully created genres at movie",
	})
	return
}

func (controller *MovieGenreControllerImpl) Delete(gc *gin.Context) {
	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format movie_id",
		})
		return
	}

	genreID, err := strconv.Atoi(gc.Param("genre_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format genre_id",
		})
		return
	}

	err = controller.MovieGenreService.Delete(gc.Request.Context(), movieID, genreID)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	gc.JSON(http.StatusOK, web.ResponseError{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully deleted genres from movie",
	})
	return
}

func (controller *MovieGenreControllerImpl) FindByID(gc *gin.Context) {
	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := controller.MovieGenreService.FindByID(gc.Request.Context(), movieID)
	if err != nil {
		gc.JSON(http.StatusOK, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	webResponse := web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data genres from movies",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}
