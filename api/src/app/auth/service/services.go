package service

import (
	"api/src/app/auth"
	"api/src/helper/errs"
)

type service struct {
	repo auth.Repository
}

func (s *service) GetUserIDByPhotoID(photoId int) (int, errs.MessageErr) {
	return s.repo.GetUserIDByFeatureID("photos", photoId)
}

func (s *service) GetUserIDByCommentID(commentId int) (int, errs.MessageErr) {
	return s.repo.GetUserIDByFeatureID("comments", commentId)
}

func (s *service) GetUserIDBySocialMediaID(socialMediaId int) (int, errs.MessageErr) {
	return s.repo.GetUserIDByFeatureID("social_media", socialMediaId)
}

func NewService(repo auth.Repository) auth.Service {
	return &service{repo}
}
