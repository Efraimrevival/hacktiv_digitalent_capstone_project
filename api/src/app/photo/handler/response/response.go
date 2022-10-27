package response

import (
	"api/src/app/photo/repository/record"
	"time"
)

type PostResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"holiday"`
	Caption   string    `json:"caption" example:"no caption"`
	PhotoUrl  string    `json:"photo_url" example:"/example"`
	UserID    int       `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

type PhotoResponse struct {
	ID        int          `json:"id" example:"1"`
	Title     string       `json:"title" example:"holiday"`
	Caption   string       `json:"caption" example:"no caption"`
	PhotoUrl  string       `json:"photo_url" example:"/example"`
	UserID    int          `json:"user_id" example:"1"`
	CreatedAt time.Time    `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
	UpdatedAt time.Time    `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
	User      UserResponse `json:"User"`
}

type UserResponse struct {
	Email    string `json:"email" example:"example@gmail.com"`
	Username string `json:"username" example:"example"`
}

type UpdateResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"holiday"`
	Caption   string    `json:"caption" example:"no caption"`
	PhotoUrl  string    `json:"photo_url" example:"/example"`
	UserID    int       `json:"user_id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

func MapToPostResponse(res record.Photo) PostResponse {
	return PostResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		CreatedAt: res.CreatedAt,
	}
}

func MapToBatchPhotoResponse(res []record.Photo) (responses []PhotoResponse) {
	for _, photo := range res {
		responses = append(responses, MapToPhotoResponse(photo))
	}
	return
}

func MapToPhotoResponse(res record.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		User: UserResponse{
			Email:    res.User.Email,
			Username: res.User.Username,
		},
	}
}

func MapToUpdateResponse(res record.Photo) UpdateResponse {
	return UpdateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UpdatedAt: res.UpdatedAt,
	}
}
