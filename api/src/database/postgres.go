package database

import (
	comment "api/src/app/comment/repository/record"
	photo "api/src/app/photo/repository/record"
	"api/src/app/socialmedia/repository/record"
	user "api/src/app/user/repository/record"
	"api/src/config/env"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresConn *gorm.DB

func InitPostgres() {
	var err error

	config := env.GetSQLEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)
	postgresConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	postgresConn.AutoMigrate(&user.User{}, &photo.Photo{}, &record.SocialMedia{}, &comment.Comment{})
}

func GetPostgresConnection() *gorm.DB {
	return postgresConn
}
