package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(gc *gin.Context)
	Login(ctx *gin.Context)
}

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthControllerImpl(authService services.AuthService) AuthController {
	return &AuthControllerImpl{AuthService: authService}
}

func (c *AuthControllerImpl) Register(gc *gin.Context) {
	var r web.AuthModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	err = c.AuthService.Register(context.Background(), &r)
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
		Message: "Successfully created user",
	})
}

func (c *AuthControllerImpl) Login(gc *gin.Context) {

	var r web.AuthModelRequest
	err := gc.ShouldBind(&r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	token, err := c.AuthService.Login(gc.Request.Context(), &r)
	if err != nil {
		gc.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    http.StatusBadRequest,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	gc.JSON(http.StatusOK, web.ResponseSuccessWithData{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: fmt.Sprintf("Success login with username %v", r.Username),
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	})
}
