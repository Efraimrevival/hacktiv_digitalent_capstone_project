package repository

import (
	"api/src/app/socialmedia"
	"api/src/app/socialmedia/repository/record"
	"api/src/helper/errs"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (repo *repository) CreateData(data *record.SocialMedia) (*record.SocialMedia, errs.MessageErr) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return data, nil
}

func (repo *repository) GetAllData() ([]record.SocialMedia, errs.MessageErr) {
	records := []record.SocialMedia{}
	if err := repo.db.Preload("User").Find(&records).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return records, nil
}

func (repo *repository) UpdateData(id int, data *record.SocialMedia) (*record.SocialMedia, errs.MessageErr) {
	if err := repo.db.Where("id", id).Updates(data).Scan(data).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return data, nil
}

func (repo *repository) DeleteData(id int) errs.MessageErr {
	if err := repo.db.Delete(new(record.SocialMedia), "id", id).Error; err != nil {
		return errs.NewError(http.StatusInternalServerError)
	}

	return nil
}

func NewGORMRepository(db *gorm.DB) socialmedia.Repository {
	return &repository{db}
}
