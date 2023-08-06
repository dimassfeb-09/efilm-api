package controller

import (
	"context"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ActorController interface {
	Save(gc *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindBySearch(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type ActorControllerImpl struct {
	ActorService services.ActorService
}

func NewActorControllerImpl(actorService services.ActorService) ActorController {
	return &ActorControllerImpl{ActorService: actorService}
}

func (c *ActorControllerImpl) Save(gc *gin.Context) {
	var r web.ActorModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = c.ActorService.Save(context.Background(), &r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

}

func (c *ActorControllerImpl) Update(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	var r web.ActorModelRequest
	err = gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	r.ID = ID
	err = c.ActorService.Update(gc.Request.Context(), &r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	gc.JSON(http.StatusOK, web.ResponseSuccess{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: fmt.Sprintf("Success update actors with ID %d", ID),
	})
}

func (c *ActorControllerImpl) Delete(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = c.ActorService.Delete(gc.Request.Context(), ID)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	gc.JSON(http.StatusOK, web.ResponseSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: fmt.Sprintf("Success delete data with ID %d", ID),
	})
}

func (c *ActorControllerImpl) FindByID(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := c.ActorService.FindByID(gc.Request.Context(), id)
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
		Message: "Success get data actors by id",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *ActorControllerImpl) FindBySearch(gc *gin.Context) {
	name := gc.Param("name")
	id := gc.Param("national_id")
	idInt, err := strconv.Atoi(id)
	var actors []*web.ActorModelResponse

	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	if name != "" {
		result, err := c.ActorService.FindByName(gc.Request.Context(), name)
		if err != nil {
			gc.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		actors = append(actors, result)
	}

	if id != "" {
		result, err := c.ActorService.FindByID(gc.Request.Context(), idInt)
		if err != nil {
			gc.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		actors = append(actors, result)
	}

	webResponse := web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data actors by id",
		Data:    actors,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *ActorControllerImpl) FindAll(gc *gin.Context) {

	var responses []*web.ActorModelResponse
	results, err := c.ActorService.FindAll(gc.Request.Context())
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data actors",
		})
		return
	}
	for _, result := range results {
		response := web.ActorModelResponse{
			ID:            result.ID,
			Name:          result.Name,
			DateOfBirth:   result.DateOfBirth,
			NationalityID: result.NationalityID,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
		}
		responses = append(responses, &response)
	}

	gc.JSON(http.StatusOK, web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data",
		Data:    responses,
	})
}
