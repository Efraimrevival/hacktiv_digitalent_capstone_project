package middleware

import (
	"
	
	
	
	api/src/app/auth"
	"
	
	
	
	api/src/app/auth/repository"
	"
	
	
	
	api/src/app/auth/service"
	"
	
	
	
	api/src/database"
	"
	
	
	
	api/src/helper"
	"
	
	
	
	api/src/helper/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	service auth.Service
}

func newAuthMiddleware(service auth.Service) *middleware {
	return &middleware{service}
}

func InitAuthMiddleware() *middleware {
	db := database.GetPostgresConnection()
	repo := repository.NewGORMRepository(db)
	serv := service.NewService(repo)
	return newAuthMiddleware(serv)
}

func (m *middleware) PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		photoId, _ := strconv.Atoi(c.Param("photoId"))
		if photoId < 1 {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				structs.Message{Message: "param must be a number greater than 0"})
			return
		}

		userId, err := m.service.GetUserIDByPhotoID(photoId)
		if err != nil {
			c.AbortWithStatusJSON(
				err.Status(),
				structs.Message{Message: err.Message()},
			)
			return
		}

		userData := helper.GetUserData(c)
		if userId != userData.ID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				structs.Message{Message: http.StatusText(http.StatusForbidden)},
			)
			return
		}

		c.Next()
	}
}

func (m *middleware) CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, _ := strconv.Atoi(c.Param("commentId"))
		if commentId < 1 {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				structs.Message{Message: "param must be a number greater than 0"})
			return
		}

		userId, err := m.service.GetUserIDByCommentID(commentId)
		if err != nil {
			c.AbortWithStatusJSON(
				err.Status(),
				structs.Message{Message: err.Message()},
			)
			return
		}

		userData := helper.GetUserData(c)
		if userId != userData.ID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				structs.Message{Message: http.StatusText(http.StatusForbidden)},
			)
			return
		}

		c.Next()
	}
}

func (m *middleware) SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
		if socialMediaId < 1 {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				structs.Message{Message: "param must be a number greater than 0"})
			return
		}

		userId, err := m.service.GetUserIDBySocialMediaID(socialMediaId)
		if err != nil {
			c.AbortWithStatusJSON(
				err.Status(),
				structs.Message{Message: err.Message()},
			)
			return
		}

		userData := helper.GetUserData(c)
		if userId != userData.ID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				structs.Message{Message: http.StatusText(http.StatusForbidden)},
			)
			return
		}

		c.Next()
	}
}
