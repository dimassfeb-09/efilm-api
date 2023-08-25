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

type GenreController interface {
	Save(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindByID(c *gin.Context)
	FindBySearch(c *gin.Context)
	FindAll(c *gin.Context)
	FindAllMoviesByID(c *gin.Context)
}

type GenreControllerImpl struct {
	GenreService services.GenreService
}

func NewGenreControllerImpl(genreService services.GenreService) GenreController {
	return &GenreControllerImpl{GenreService: genreService}
}

func (controller *GenreControllerImpl) Save(c *gin.Context) {
	var r web.GenreModelRequest
	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = controller.GenreService.Save(context.Background(), &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ResponseError{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully create data genre",
	})
	return

}

func (controller *GenreControllerImpl) Update(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	var r web.GenreModelRequest
	err = c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	r.ID = ID
	err = controller.GenreService.Update(c.Request.Context(), &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ResponseSuccess{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: fmt.Sprintf("Success update genres with ID %d", ID),
	})
}

func (controller *GenreControllerImpl) Delete(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = controller.GenreService.Delete(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ResponseSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: fmt.Sprintf("Success delete data with ID %d", ID),
	})
}

func (controller *GenreControllerImpl) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := controller.GenreService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	webResponse := web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data genres by id",
		Data:    result,
	}

	c.JSON(http.StatusOK, webResponse)
	return
}

func (controller *GenreControllerImpl) FindBySearch(c *gin.Context) {
	name := c.Query("name")

	var genres []*web.GenreModelResponse
	if name != "" {
		result, err := controller.GenreService.FindByName(c.Request.Context(), name)
		if err != nil {
			c.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		genres = append(genres, result)
	} else {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Query name is required",
		})
		return
	}

	webResponse := web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data genres by parameter",
		Data:    genres,
	}

	c.JSON(http.StatusOK, webResponse)
	return
}

func (controller *GenreControllerImpl) FindAll(c *gin.Context) {

	var responses []*web.GenreModelResponse
	results, err := controller.GenreService.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data genres",
		})
		return
	}
	for _, result := range results {
		response := web.GenreModelResponse{
			ID:   result.ID,
			Name: result.Name,
		}
		responses = append(responses, &response)
	}

	c.JSON(http.StatusOK, web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data",
		Data:    responses,
	})
}

func (controller *GenreControllerImpl) FindAllMoviesByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	responses, err := controller.GenreService.FindAllMoviesByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data genres",
		})
		return
	}

	c.JSON(http.StatusOK, web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data",
		Data:    responses,
	})
}
