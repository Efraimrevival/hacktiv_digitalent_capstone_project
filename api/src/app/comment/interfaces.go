package comment

import (
	"api/src/app/comment/repository/record"
	"api/src/helper/errs"
)

type Service interface {
	CreateComment(comment *record.Comment) (*record.Comment, errs.MessageErr)
	GetAllComments() ([]record.Comment, errs.MessageErr)
	UpdateComment(id int, message string) (*record.Comment, errs.MessageErr)
	DeleteComment(id int) errs.MessageErr
}

type Repository interface {
	CreateData(data *record.Comment) (*record.Comment, errs.MessageErr)
	GetAllData() ([]record.Comment, errs.MessageErr)
	UpdateData(id int, data string) (*record.Comment, errs.MessageErr)
	DeleteData(id int) errs.MessageErr
}
