package factory

import (
	"api/src/app/comment/handler"
	"api/src/app/comment/repository"
	"api/src/app/comment/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
