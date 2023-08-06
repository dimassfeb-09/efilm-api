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

type NationalController interface {
	Save(gc *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindBySearch(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type NationalControllerImpl struct {
	NationalService services.NationalService
}

func NewNationalControllerImpl(nationalService services.NationalService) NationalController {
	return &NationalControllerImpl{NationalService: nationalService}
}

func (c *NationalControllerImpl) Save(gc *gin.Context) {
	var r web.NationalModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = c.NationalService.Save(context.Background(), &r)
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
		Message: "Successfully create data national",
	})
	return

}

func (c *NationalControllerImpl) Update(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	var r web.NationalModelRequest
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
	err = c.NationalService.Update(gc.Request.Context(), &r)
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
		Message: fmt.Sprintf("Success update nationals with ID %d", ID),
	})
}

func (c *NationalControllerImpl) Delete(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = c.NationalService.Delete(gc.Request.Context(), ID)
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

func (c *NationalControllerImpl) FindByID(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := c.NationalService.FindByID(gc.Request.Context(), id)
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
		Message: "Success get data nationals by id",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *NationalControllerImpl) FindBySearch(gc *gin.Context) {
	name := gc.Query("name")

	var nationals []*web.NationalModelResponse
	if name != "" {
		result, err := c.NationalService.FindByName(gc.Request.Context(), name)
		if err != nil {
			gc.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		nationals = append(nationals, result)
	} else {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Query name is required",
		})
		return
	}

	webResponse := web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data nationals by parameter",
		Data:    nationals,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *NationalControllerImpl) FindAll(gc *gin.Context) {

	var responses []*web.NationalModelResponse
	results, err := c.NationalService.FindAll(gc.Request.Context())
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data nationals",
		})
		return
	}
	for _, result := range results {
		response := web.NationalModelResponse{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
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
