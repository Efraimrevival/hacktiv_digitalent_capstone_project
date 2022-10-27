package repository

import (
	"api/src/app/auth"
	"api/src/helper/errs"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// GetUserIDByFeatureID return the UserID by the given feature ID *example: photos, comments, etc
func (repo *repository) GetUserIDByFeatureID(feature string, id int) (int, errs.MessageErr) {
	userId := 0
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE id=%d", feature, id)
	err := repo.db.Raw(query).Scan(&userId).Error
	if err != nil {
		return 0, errs.NewError(http.StatusInternalServerError)
	}

	if err == nil && userId == 0 {
		return 0, errs.NewError(http.StatusNotFound)
	}

	return userId, nil
}

func NewGORMRepository(db *gorm.DB) auth.Repository {
	return &repository{db}
}
