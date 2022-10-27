package repository

import (
	"api/src/app/comment"
	"api/src/app/comment/repository/record"
	"api/src/helper/errs"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (repo *repository) CreateData(data *record.Comment) (*record.Comment, errs.MessageErr) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return data, nil
}

func (repo *repository) GetAllData() ([]record.Comment, errs.MessageErr) {
	records := []record.Comment{}
	if err := repo.db.Preload("User").Preload("Photo").Find(&records).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return records, nil
}

func (repo *repository) UpdateData(id int, message string) (*record.Comment, errs.MessageErr) {
	data := &record.Comment{}
	if err := repo.db.Model(&data).Where("id", id).Update("message", message).
		Scan(data).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return data, nil
}

func (repo *repository) DeleteData(id int) errs.MessageErr {
	if err := repo.db.Delete(new(record.Comment), "id", id).Error; err != nil {
		return errs.NewError(http.StatusInternalServerError)
	}

	return nil
}

func NewGORMRepository(db *gorm.DB) comment.Repository {
	return &repository{db}
}
