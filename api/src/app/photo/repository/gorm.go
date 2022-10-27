package repository

import (
	"api/src/app/photo"
	"api/src/app/photo/repository/record"
	"api/src/helper/errs"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// CreateData create new data from the given input
func (repo *repository) CreateData(data *record.Photo) (*record.Photo, errs.MessageErr) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return data, nil
}

// GetAllData find all data
func (repo *repository) GetAllData() ([]record.Photo, errs.MessageErr) {
	records := []record.Photo{}
	if err := repo.db.Preload("User").Find(&records).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return records, nil
}

// UpdateData Update the data by the given id
func (repo *repository) UpdateData(id int, data *record.Photo) (*record.Photo, errs.MessageErr) {
	if err := repo.db.Where("id", id).Updates(data).Scan(data).Error; err != nil {
		return nil, errs.NewError(http.StatusInternalServerError)
	}

	data.ID = id
	return data, nil
}

// DeleteData delete the data by the given id
func (repo *repository) DeleteData(id int) errs.MessageErr {
	if err := repo.db.Delete(new(record.Photo), "id", id).Error; err != nil {
		return errs.NewError(http.StatusInternalServerError)
	}

	return nil
}

func NewGORMRepository(db *gorm.DB) photo.Repository {
	return &repository{db}
}
