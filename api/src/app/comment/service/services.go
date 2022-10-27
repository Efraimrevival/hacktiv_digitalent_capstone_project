package service

import (
	"api/src/app/comment"
	"api/src/app/comment/repository/record"
	"api/src/helper/errs"
)

type service struct {
	repo comment.Repository
}

func (s *service) CreateComment(comment *record.Comment) (*record.Comment, errs.MessageErr) {
	return s.repo.CreateData(comment)
}

func (s *service) GetAllComments() ([]record.Comment, errs.MessageErr) {
	return s.repo.GetAllData()
}

func (s *service) UpdateComment(id int, message string) (*record.Comment, errs.MessageErr) {
	return s.repo.UpdateData(id, message)
}

func (s *service) DeleteComment(id int) errs.MessageErr {
	return s.repo.DeleteData(id)
}

func NewService(repo comment.Repository) comment.Service {
	return &service{repo}
}
