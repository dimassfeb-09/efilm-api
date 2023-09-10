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

type RecommendationMovieController interface {
	Save(gc *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type RecommendationMovieControllerImpl struct {
	RecommendationMovieService services.RecommendationMovieService
}

func NewRecommendationMovieControllerImpl(recommendationService services.RecommendationMovieService) RecommendationMovieController {
	return &RecommendationMovieControllerImpl{RecommendationMovieService: recommendationService}
}

func (c *RecommendationMovieControllerImpl) Save(gc *gin.Context) {
	var r web.RecommendationMovieModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = c.RecommendationMovieService.Save(context.Background(), r.MovieID)
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
		Message: "Successfully created recommendations",
	})
	return
}

func (c *RecommendationMovieControllerImpl) Delete(gc *gin.Context) {
	ID, err := strconv.Atoi(gc.Param("movie_id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Invalid format ID",
		})
		return
	}

	err = c.RecommendationMovieService.Delete(gc.Request.Context(), ID)
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
		Message: fmt.Sprintf("Success delete data with Movie ID %d", ID),
	})
	return
}

func (c *RecommendationMovieControllerImpl) FindAll(gc *gin.Context) {

	var responses []*web.RecommendationMovieModelResponse
	results, err := c.RecommendationMovieService.FindAll(gc.Request.Context())
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: "Failed get all data recommendations",
		})
		return
	}
	for _, result := range results {
		response := web.RecommendationMovieModelResponse{
			ID:          result.ID,
			Title:       result.Title,
			ReleaseDate: result.ReleaseDate,
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
