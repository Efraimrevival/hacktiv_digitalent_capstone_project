package factory

import (
	"api/src/app/user/handler"
	"api/src/app/user/repository"
	"api/src/app/user/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
