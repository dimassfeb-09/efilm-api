package controller

import (
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UsersController interface {
	GetUserInfo(c *gin.Context)
}

type UsersControllerImpl struct {
}

func NewUserController() UsersController {
	return &UsersControllerImpl{}
}

func (controller *UsersControllerImpl) GetUserInfo(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	bearers := strings.Split(authorization, "Bearer")
	token := strings.TrimSpace(bearers[1])
	isValid, userInfo, err := helpers.ValidateTokenJWT(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{
			Code:    400,
			Status:  "Status Bad Request",
			Message: err.Error(),
		})
		return
	}

	if isValid {
		c.JSON(http.StatusOK, web.ResponseSuccessWithData{
			Code:    200,
			Status:  "Status OK",
			Message: "Success to get user info",
			Data: web.UserInfoResponse{
				UserID:   userInfo.UserID,
				Username: userInfo.Username,
			},
		})
		return
	}
}
