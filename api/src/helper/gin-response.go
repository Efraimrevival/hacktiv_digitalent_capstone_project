package helper

import (
	"api/src/helper/structs"

	"github.com/gin-gonic/gin"
)

func CreateMessageResponse(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, structs.Message{
		Message: message,
	})
}
