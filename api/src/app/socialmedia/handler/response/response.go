package response

import (
	"api/src/app/socialmedia/repository/record"
	"time"
)

type PostResponse struct {
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"example"`
	SocialMediaUrl string    `json:"social_media_url" example:"/example"`
	UserID         int       `json:"user_id" example:"1"`
	CreatedAt      time.Time `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

type Response struct {
	SocialMedias []SocialMediaResponse `json:"social_medias"`
}

type SocialMediaResponse struct {
	ID             int          `json:"id" example:"1"`
	Name           string       `json:"name" example:"example"`
	SocialMediaUrl string       `json:"social_media_url" example:"/example"`
	UserID         int          `json:"user_id" example:"1"`
	CreatedAt      time.Time    `json:"created_at" example:"2022-10-07T15:54:24.575005+07:00"`
	UpdatedAt      time.Time    `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
	User           UserResponse `json:"User"`
}

type UserResponse struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"example"`
}

type UpdateResponse struct {
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"example"`
	SocialMediaUrl string    `json:"social_media_url" example:"/example"`
	UserID         int       `json:"user_id" example:"1"`
	UpdatedAt      time.Time `json:"updated_at" example:"2022-10-07T15:54:24.575005+07:00"`
}

func MapToPostResponse(res record.SocialMedia) PostResponse {
	return PostResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		CreatedAt:      res.CreatedAt,
	}
}

func MapToBatchSocialMediaResponse(res []record.SocialMedia) Response {
	socialMedias := []SocialMediaResponse{}
	for _, socialMedia := range res {
		socialMedias = append(socialMedias, MapToSocialMediaResponse(socialMedia))
	}

	return Response{SocialMedias: socialMedias}
}

func MapToSocialMediaResponse(res record.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
		User: UserResponse{
			ID:       res.UserID,
			Username: res.User.Username,
		},
	}
}

func MapToUpdateResponse(res record.SocialMedia) UpdateResponse {
	return UpdateResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		UpdatedAt:      res.UpdatedAt,
	}
}
