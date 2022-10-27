package service

import (
	"api/src/app/user"
	"api/src/app/user/handler/request"
	"api/src/app/user/repository/record"
	"api/src/helper"
	"api/src/helper/errs"
	"net/http"
)

type service struct {
	repo user.Repository
}

// RegisterUser register the new user by create new data
func (s *service) RegisterUser(user *record.User) (*record.User, errs.MessageErr) {
	return s.repo.CreateData(user)
}

// LoginUser login the user and return the jwt if request valid
func (s *service) LoginUser(login request.LoginRequest) (*string, errs.MessageErr) {
	record, err := s.repo.FindDataByEmail(login.Email)
	if err != nil {
		return nil, err
	}

	if !helper.ValidateHash(login.Password, record.Password) {
		return nil, errs.NewCustomError(http.StatusUnauthorized, "invalid email or password")
	}

	token := helper.GenerateJWT(record.ID)
	return &token, nil
}

// UpdateUser update user data by the given id
func (s *service) UpdateUser(id int, user *record.User) (*record.User, errs.MessageErr) {
	return s.repo.UpdateData(id, user)
}

// DeleteUser delete the user data by the given id
func (s *service) DeleteUser(id int) errs.MessageErr {
	return s.repo.DeleteData(id)
}

func NewService(repo user.Repository) user.Service {
	return &service{repo}
}
