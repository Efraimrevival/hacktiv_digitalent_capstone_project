package request

import "api/src/app/comment/repository/record"

type PostRequest struct {
	Message string `json:"message" binding:"required" example:"example"`
	PhotoID int    `json:"photo_id" binding:"required" example:"1"`
}

type UpdateRequest struct {
	Message string `json:"message" binding:"required" example:"example"`
}

func (req *PostRequest) MapToRecord(userId int) *record.Comment {
	return &record.Comment{
		UserID:  userId,
		Message: req.Message,
		PhotoID: req.PhotoID,
	}
}
