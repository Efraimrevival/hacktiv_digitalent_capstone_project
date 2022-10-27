package auth

import "api/src/helper/errs"

type Service interface {
	GetUserIDByPhotoID(photoId int) (int, errs.MessageErr)
	GetUserIDByCommentID(commentId int) (int, errs.MessageErr)
	GetUserIDBySocialMediaID(socialMediaId int) (int, errs.MessageErr)
}

type Repository interface {
	GetUserIDByFeatureID(feature string, id int) (int, errs.MessageErr)
}
