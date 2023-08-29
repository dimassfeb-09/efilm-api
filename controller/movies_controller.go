package controller

import (
	"context"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieController interface {
	Save(c *gin.Context)
	Update(c *gin.Context)
	UploadPoster(c *gin.Context)
	Delete(c *gin.Context)
	FindByID(c *gin.Context)
	FindBySearch(c *gin.Context)
	FindAll(c *gin.Context)
}

type MovieControllerImpl struct {
	MovieService services.MovieService
}

func NewMovieControllerImpl(movieService services.MovieService) MovieController {
	return &MovieControllerImpl{MovieService: movieService}
}

func (controller *MovieControllerImpl) Save(c *gin.Context) {
	var r web.MovieModelRequest
	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	id, err := controller.MovieService.Save(context.Background(), &r)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created movies",
		Data: struct {
			MovieID int `json:"movie_id"`
		}{
			MovieID: id,
		},
	})
	return
}

func (controller *MovieControllerImpl) Update(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	var r web.MovieModelRequest
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
	err = controller.MovieService.Update(c.Request.Context(), &r)
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
		Message: fmt.Sprintf("Success update movies with ID %d", ID),
	})
	return
}

func (controller *MovieControllerImpl) UploadPoster(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format movie ID",
		})
		return
	}

	fileHeader, err := c.FormFile("poster_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Cannot process file.",
		})
		return
	}

	contentType := fileHeader.Header.Get("Content-Type")
	isValid := helpers.VerfiyFileType(contentType)
	if isValid == false {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: fmt.Sprintf("File %s not accept, only image/png, image/jpg, image/jpeg", contentType),
		})
		return
	}

	err = controller.MovieService.UploadFile(c.Request.Context(), ID, fileHeader)
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
		Message: "Success upload file.",
	})
	return
}

func (controller *MovieControllerImpl) Delete(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = controller.MovieService.Delete(c.Request.Context(), ID)
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
	return
}

func (controller *MovieControllerImpl) FindByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	result, err := controller.MovieService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	webResponse := web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data movies by id",
		Data:    result,
	}

	c.JSON(http.StatusOK, webResponse)
	return
}

func (controller *MovieControllerImpl) FindBySearch(c *gin.Context) {
	title := c.Query("title")
	var movies []*web.MovieModelResponse

	if title != "" {
		result, err := controller.MovieService.FindByTitle(c.Request.Context(), title)
		if err != nil {
			c.JSON(http.StatusOK, web.ResponseError{
				Code:    http.StatusBadRequest,
				Status:  "Status Bad Request",
				Message: err.Error(),
			})
			return
		}
		movies = append(movies, result)
	}

	webResponse := web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data movies by search",
		Data:    movies,
	}

	c.JSON(http.StatusOK, webResponse)
	return
}

func (controller *MovieControllerImpl) FindAll(c *gin.Context) {

	var responses []*web.MovieModelResponse
	results, err := controller.MovieService.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
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
			GenreIDS:    result.GenreIDS,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		}
		responses = append(responses, &response)
	}

	c.JSON(http.StatusOK, web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get data",
		Data:    responses,
	})
}
