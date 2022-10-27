package response

import (
	"api/src/app/comment/repository/record"
	"time"
)

type PostResponse struct {
	ID        int       `json:"id" example:"1"`
	Message   string    `json:"message" example:"wow"`
	PhotoID   int       `json:"photo_id" example:"1"`
	UserID    int       `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

type CommentResponse struct {
	ID        int           `json:"id" example:"1"`
	Message   string        `json:"message" example:"wow"`
	PhotoID   int           `json:"photo_id" example:"1"`
	UserID    int           `json:"user_id" example:"1"`
	CreatedAt time.Time     `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
	UpdatedAt time.Time     `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
	User      UserResponse  `json:"User"`
	Photo     PhotoResponse `json:"Photo"`
}

type UserResponse struct {
	ID       int    `json:"id" example:"1"`
	Email    string `json:"email" example:"example@gmail.com"`
	Username string `json:"username" example:"example"`
}

type PhotoResponse struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"holiday"`
	Caption  string `json:"caption" example:"no caption"`
	PhotoUrl string `json:"photo_url" example:"/example"`
	UserID   int    `json:"user_id" example:"1"`
}

type UpdateResponse struct {
	ID        int       `json:"id" example:"1"`
	Message   string    `json:"message" example:"wow"`
	PhotoID   int       `json:"photo_id" example:"1"`
	UserID    int       `json:"user_id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

func MapToPostResponse(res record.Comment) PostResponse {
	return PostResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}
}

func MapToBatchCommentResponse(res []record.Comment) (responses []CommentResponse) {
	for _, comment := range res {
		responses = append(responses, mapToCommentResponse(comment))
	}
	return
}

func mapToCommentResponse(res record.Comment) CommentResponse {
	return CommentResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		User: UserResponse{
			ID:       res.User.ID,
			Email:    res.User.Email,
			Username: res.User.Username,
		},
		Photo: PhotoResponse{
			ID:       res.Photo.ID,
			Title:    res.Photo.Title,
			Caption:  res.Photo.Caption,
			PhotoUrl: res.Photo.PhotoUrl,
			UserID:   res.Photo.UserID,
		},
	}
}

func MapToUpdateResponse(res record.Comment) UpdateResponse {
	return UpdateResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}
}
