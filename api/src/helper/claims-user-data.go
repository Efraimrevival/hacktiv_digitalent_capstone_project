package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type data struct {
	ID int
}

func GetUserData(c *gin.Context) data {
	userData := c.MustGet("userData").(jwt.MapClaims)
	return data{
		ID: int(userData["id"].(float64)),
	}
}
