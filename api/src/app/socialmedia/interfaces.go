package socialmedia

import (
	"api/src/app/socialmedia/repository/record"
	"api/src/helper/errs"
)

type Service interface {
	CreateSocialMedia(socialMedia *record.SocialMedia) (*record.SocialMedia, errs.MessageErr)
	GetAllSocialMedias() ([]record.SocialMedia, errs.MessageErr)
	UpdateSocialMedia(id int, socialMedia *record.SocialMedia) (*record.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(id int) errs.MessageErr
}

type Repository interface {
	CreateData(data *record.SocialMedia) (*record.SocialMedia, errs.MessageErr)
	GetAllData() ([]record.SocialMedia, errs.MessageErr)
	UpdateData(id int, data *record.SocialMedia) (*record.SocialMedia, errs.MessageErr)
	DeleteData(id int) errs.MessageErr
}
