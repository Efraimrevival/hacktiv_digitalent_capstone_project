package response

import (
	"mygram-api/src/app/user/repository/record"
	"time"
)

type RegisterResponse struct {
	ID       int    `json:"id" example:"1"`
	Email    string `json:"email" example:"example@gmail.com"`
	Username string `json:"username" example:"example"`
	Age      int    `json:"age" example:"19"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type UpdateResponse struct {
	ID        int       `json:"id" example:"1"`
	Email     string    `json:"email" example:"example@gmail.com"`
	Username  string    `json:"username" example:"example"`
	Age       int       `json:"age" example:"19"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

func MapToRegisterResponse(data record.User) RegisterResponse {
	return RegisterResponse{
		ID:       data.ID,
		Email:    data.Email,
		Username: data.Username,
		Age:      data.Age,
	}
}

func MapToUpdateResponse(data record.User) UpdateResponse {
	return UpdateResponse{
		ID:        data.ID,
		Email:     data.Email,
		Username:  data.Username,
		Age:       data.Age,
		UpdatedAt: data.UpdatedAt,
	}
}
