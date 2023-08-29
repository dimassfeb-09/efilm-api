package controller

import (
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieActorController interface {
	Save(gc *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type MovieActorControllerImpl struct {
	MovieActorService services.MovieActorService
}

func NewMovieActorControllerImpl(actorService services.MovieActorService) MovieActorController {
	return &MovieActorControllerImpl{MovieActorService: actorService}
}

func (controller *MovieActorControllerImpl) Save(gc *gin.Context) {
	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format Movie ID",
		})
		return
	}

	var movieActor web.MovieActorModelRequestPost
	err = gc.ShouldBind(&movieActor)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format Movie ID",
		})
		return
	}

	movieActor.MovieID = movieID
	err = controller.MovieActorService.Save(gc.Request.Context(), &movieActor)
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
		Message: "Successfully created actors at movie",
	})
	return
}

func (controller *MovieActorControllerImpl) Update(gc *gin.Context) {

	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format Movie ID",
		})
		return
	}

	actorID, err := strconv.Atoi(gc.Param("actor_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format Actor ID",
		})
		return
	}

	var movieActor web.MovieActorModelRequestPut
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
	movieActor.ActorID = actorID
	err = controller.MovieActorService.Update(gc.Request.Context(), &movieActor)
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
		Message: "Successfully update actors from movie",
	})
	return
}

func (controller *MovieActorControllerImpl) Delete(gc *gin.Context) {
	actorID, err := strconv.Atoi(gc.Param("actor_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = controller.MovieActorService.Delete(gc.Request.Context(), actorID)
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
		Message: "Successfully deleted actors from movie",
	})
	return
}

func (controller *MovieActorControllerImpl) FindByID(gc *gin.Context) {
	movieID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := controller.MovieActorService.FindByID(gc.Request.Context(), movieID)
	if err != nil {
		gc.JSON(http.StatusOK, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	webResponse := web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success deleted data actors from movies",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}
