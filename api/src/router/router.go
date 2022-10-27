package router

import (
	"api/src/adapter"
	"api/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func StartServer() *gin.Engine {
	router := gin.Default()
	handler := adapter.Init()
	auth := middleware.InitAuthMiddleware()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to MyGram API"})
	})

	user := router.Group("/users")
	user.POST("/register", handler.User.RegisterUserHandler)
	user.POST("/login", handler.User.LoginUserHandler)
	user.PUT("", middleware.Authentication(), handler.User.UpdateUserHandler)
	user.DELETE("", middleware.Authentication(), handler.User.DeleteUserHandler)

	photo := router.Group("/photos").Use(middleware.Authentication())
	photo.POST("", handler.Photo.PostPhotoHandler)
	photo.GET("", handler.Photo.GetAllPhotosHandler)
	photo.PUT("/:photoId", auth.PhotoAuthorization(), handler.Photo.UpdatePhotoHandler)
	photo.DELETE("/:photoId", auth.PhotoAuthorization(), handler.Photo.DeletePhotoHandler)

	socialMedia := router.Group("/socialmedias").Use(middleware.Authentication())
	socialMedia.POST("", handler.SocialMedia.CreateSocialMediaHandler)
	socialMedia.GET("", handler.SocialMedia.GetAllSocialMediasHandler)
	socialMedia.PUT("/:socialMediaId", auth.SocialMediaAuthorization(), handler.SocialMedia.UpdateSocialMediaHandler)
	socialMedia.DELETE("/:socialMediaId", auth.SocialMediaAuthorization(), handler.SocialMedia.DeleteSocialMediaHandler)

	comment := router.Group("/comments").Use(middleware.Authentication())
	comment.POST("", handler.Comment.CreateCommentHandler)
	comment.GET("", handler.Comment.GetAllCommentsHandler)
	comment.PUT("/:commentId", auth.CommentAuthorization(), handler.Comment.UpdateCommentHandler)
	comment.DELETE("/:commentId", auth.CommentAuthorization(), handler.Comment.DeleteCommentHandler)

	return router
}
