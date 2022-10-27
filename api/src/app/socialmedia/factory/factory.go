package factory

import (
	"api/src/app/socialmedia/handler"
	"api/src/app/socialmedia/repository"
	"api/src/app/socialmedia/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
