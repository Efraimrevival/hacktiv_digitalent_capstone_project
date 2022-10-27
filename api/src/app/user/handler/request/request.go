package request

import (
	"api/src/app/user/repository/record"
	"api/src/helper"
)

type RegisterRequest struct {
	Age      int    `json:"age" binding:"required,min=9" example:"19"`
	Email    string `json:"email" binding:"required,email" example:"example@gmail.com"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
	Username string `json:"username" binding:"required" example:"example"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"example@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type UpdateRequest struct {
	Email    string `json:"email" binding:"required,email" example:"example@gmail.com"`
	Username string `json:"username" binding:"required" example:"example"`
}

func (req *RegisterRequest) MapToRecord() *record.User {
	return &record.User{
		Age:      req.Age,
		Email:    req.Email,
		Password: helper.CreateHash(req.Password),
		Username: req.Username,
	}
}

func (req *UpdateRequest) MapToRecord() *record.User {
	return &record.User{
		Email:    req.Email,
		Username: req.Username,
	}
}
