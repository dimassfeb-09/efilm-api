package controller

import (
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieDirectorController interface {
	Save(gc *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type MovieDirectorControllerImpl struct {
	MovieDirectorService services.MovieDirectorService
}

func NewMovieDirectorControllerImpl(directorService services.MovieDirectorService) MovieDirectorController {
	return &MovieDirectorControllerImpl{MovieDirectorService: directorService}
}

func (controller *MovieDirectorControllerImpl) Save(gc *gin.Context) {

	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format movie_id",
		})
		return
	}

	var movieActor web.MovieDirectorModelRequestPost
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

	err = controller.MovieDirectorService.Save(gc.Request.Context(), &movieActor)
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
		Message: "Successfully created directors at movie",
	})
	return
}

func (controller *MovieDirectorControllerImpl) Delete(gc *gin.Context) {
	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format movie_id",
		})
		return
	}

	directorID, err := strconv.Atoi(gc.Param("director_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format director_id",
		})
		return
	}

	err = controller.MovieDirectorService.Delete(gc.Request.Context(), movieID, directorID)
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
		Message: "Successfully deleted directors from movie",
	})
	return
}

func (controller *MovieDirectorControllerImpl) FindByID(gc *gin.Context) {
	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := controller.MovieDirectorService.FindByID(gc.Request.Context(), movieID)
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
		Message: "Success get data directors from movies",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}
