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

type MovieController interface {
	Save(gc *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindBySearch(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type MovieControllerImpl struct {
	MovieService services.MovieService
}

func NewMovieControllerImpl(movieService services.MovieService) MovieController {
	return &MovieControllerImpl{MovieService: movieService}
}

func (c *MovieControllerImpl) Save(gc *gin.Context) {
	var r web.MovieModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = c.MovieService.Save(context.Background(), &r)
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
		Message: "Successfully created movies",
	})
	return
}

func (c *MovieControllerImpl) Update(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	var r web.MovieModelRequest
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
	err = c.MovieService.Update(gc.Request.Context(), &r)
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
		Message: fmt.Sprintf("Success update movies with ID %d", ID),
	})
	return
}

func (c *MovieControllerImpl) Delete(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = c.MovieService.Delete(gc.Request.Context(), ID)
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
	return
}

func (c *MovieControllerImpl) FindByID(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := c.MovieService.FindByID(gc.Request.Context(), id)
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
		Message: "Success get data movies by id",
		Data:    result,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *MovieControllerImpl) FindBySearch(gc *gin.Context) {
	title := gc.Query("title")
	var movies []*web.MovieModelResponse

	if title != "" {
		result, err := c.MovieService.FindByTitle(gc.Request.Context(), title)
		if err != nil {
			gc.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		movies = append(movies, result)
	}

	webResponse := web.ResponseGetSuccess{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data movies by search",
		Data:    movies,
	}

	gc.JSON(http.StatusOK, webResponse)
	return
}

func (c *MovieControllerImpl) FindAll(gc *gin.Context) {

	var responses []*web.MovieModelResponse
	results, err := c.MovieService.FindAll(gc.Request.Context())
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data movies",
		})
		return
	}
	for _, result := range results {
		response := web.MovieModelResponse{
			ID:          result.ID,
			Title:       result.Title,
			ReleaseDate: result.ReleaseDate,
			Duration:    result.Duration,
			Plot:        result.Plot,
			PosterUrl:   result.PosterUrl,
			TrailerUrl:  result.TrailerUrl,
			Language:    result.Language,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
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
