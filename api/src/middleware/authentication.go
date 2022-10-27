package middleware

import (
	"api/src/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := helper.ValidateJWT(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": err.Error()})
			return
		}

		c.Set("userData", token)
		c.Next()
	}
}
