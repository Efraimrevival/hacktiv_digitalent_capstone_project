package request

import "api/src/app/socialmedia/repository/record"

type Request struct {
	Name           string `json:"name" binding:"required" example:"example"`
	SocialMediaUrl string `json:"social_media_url" binding:"required" example:"/example"`
}

func (r *Request) MapPostToRecord(userID int) *record.SocialMedia {
	return &record.SocialMedia{
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserID:         userID,
	}
}

func (r *Request) MapUpdateToRecord() *record.SocialMedia {
	return &record.SocialMedia{
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
	}
}
