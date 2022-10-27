package service

import (
	"mygram-api/src/app/photo"
	"mygram-api/src/app/photo/repository/record"
	"mygram-api/src/helper/errs"
)

type service struct {
	repo photo.Repository
}

// PostPhoto post new photo
func (s *service) PostPhoto(photo *record.Photo) (*record.Photo, errs.MessageErr) {
	return s.repo.CreateData(photo)
}

// GetAllPhotos get all photos
func (s *service) GetAllPhotos() ([]record.Photo, errs.MessageErr) {
	return s.repo.GetAllData()
}

// UpdatePhoto update photo data by the given id
func (s *service) UpdatePhoto(id int, photo *record.Photo) (*record.Photo, errs.MessageErr) {
	return s.repo.UpdateData(id, photo)
}

// DeletePhoto delete the photo by the given id
func (s *service) DeletePhoto(id int) errs.MessageErr {
	return s.repo.DeleteData(id)
}

func NewService(repo photo.Repository) photo.Service {
	return &service{repo}
}
