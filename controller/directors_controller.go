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

type DirectorController interface {
	Save(gc *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindBySearch(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type DirectorControllerImpl struct {
	DirectorService services.DirectorService
}

func NewDirectorControllerImpl(directorService services.DirectorService) DirectorController {
	return &DirectorControllerImpl{DirectorService: directorService}
}

func (c *DirectorControllerImpl) Save(gc *gin.Context) {
	var r web.DirectorModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = c.DirectorService.Save(context.Background(), &r)
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
		Message: "Successfully add data director",
	})
	return
}

func (c *DirectorControllerImpl) Update(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	var r web.DirectorModelRequest
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
	err = c.DirectorService.Update(gc.Request.Context(), &r)
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
		Message: fmt.Sprintf("Success update directors with ID %d", ID),
	})
}

func (c *DirectorControllerImpl) Delete(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = c.DirectorService.Delete(gc.Request.Context(), ID)
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

func (c *DirectorControllerImpl) FindByID(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := c.DirectorService.FindByID(gc.Request.Context(), id)
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
		Message: "Success get data directors by id",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *DirectorControllerImpl) FindBySearch(gc *gin.Context) {
	name := gc.Query("name")
	id := gc.Query("national_id")
	var directors []*web.DirectorModelResponse

	if name != "" {
		result, err := c.DirectorService.FindByName(gc.Request.Context(), name)
		if err != nil {
			gc.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		directors = append(directors, result)
	}

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			gc.JSON(http.StatusBadRequest, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: "Invalid format ID",
			})
			return
		}

		result, err := c.DirectorService.FindByID(gc.Request.Context(), idInt)
		if err != nil {
			gc.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		directors = append(directors, result)
	}

	webResponse := web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data directors by id",
		Data:    directors,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *DirectorControllerImpl) FindAll(gc *gin.Context) {

	var responses []*web.DirectorModelResponse
	results, err := c.DirectorService.FindAll(gc.Request.Context())
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data directors",
		})
		return
	}
	for _, result := range results {
		response := web.DirectorModelResponse{
			ID:            result.ID,
			Name:          result.Name,
			DateOfBirth:   result.DateOfBirth,
			NationalityID: result.NationalityID,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
		}
		responses = append(responses, &response)
	}

	gc.JSON(http.StatusOK, web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data",
		Data:    responses,
	})
}
