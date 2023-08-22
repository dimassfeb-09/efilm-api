package middlewares

import (
	"net/http"
	"strings"

	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/gin-gonic/gin"
)

func MiddlewareToken(c *gin.Context) {
	if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.ResponseError{
				Code:    http.StatusUnauthorized,
				Status:  "Status Unauthorized",
				Message: "Authorization header not found",
			})
			return
		}
		bearers := strings.Split(authorization, "Bearer")
		if bearers[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.ResponseError{
				Code:    http.StatusUnauthorized,
				Status:  "Status Unauthorized",
				Message: "Token not found",
			})
			return
		} else {
			token := strings.TrimSpace(bearers[1])
			isValid, err := helpers.ValidateTokenJWT(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, web.ResponseError{
					Code:    http.StatusUnauthorized,
					Status:  "Status Unauthorized",
					Message: err.Error(),
				})
				return
			}

			if isValid {
				c.Next()
			}
		}
	}
}
