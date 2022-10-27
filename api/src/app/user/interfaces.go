package user

import (
	"api/src/app/user/handler/request"
	"api/src/app/user/repository/record"
	"api/src/helper/errs"
)

type Service interface {
	RegisterUser(user *record.User) (*record.User, errs.MessageErr)
	LoginUser(login request.LoginRequest) (*string, errs.MessageErr)
	UpdateUser(id int, user *record.User) (*record.User, errs.MessageErr)
	DeleteUser(id int) errs.MessageErr
}

type Repository interface {
	CreateData(data *record.User) (*record.User, errs.MessageErr)
	FindDataByEmail(email string) (*record.User, errs.MessageErr)
	UpdateData(id int, data *record.User) (*record.User, errs.MessageErr)
	DeleteData(id int) errs.MessageErr
}
