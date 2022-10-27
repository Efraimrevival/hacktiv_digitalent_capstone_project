package adapter

import (
	comment_factory "api/src/app/comment/factory"
	comment_handler "api/src/app/comment/handler"
	photo_factory "api/src/app/photo/factory"
	photo_handler "api/src/app/photo/handler"
	social_media_factory "api/src/app/socialmedia/factory"
	social_media_handler "api/src/app/socialmedia/handler"
	user_factory "api/src/app/user/factory"
	user_handler "api/src/app/user/handler"
	"api/src/database"
)

type handlers struct {
	User        user_handler.Handler
	Photo       photo_handler.Handler
	SocialMedia social_media_handler.Handler
	Comment     comment_handler.Handler
}

func Init() handlers {
	conn := database.GetPostgresConnection()
	return handlers{
		User:        user_factory.Factory(conn),
		Photo:       photo_factory.Factory(conn),
		SocialMedia: social_media_factory.Factory(conn),
		Comment:     comment_factory.Factory(conn),
	}
}
