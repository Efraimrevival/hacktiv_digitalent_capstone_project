package request

import "api/src/app/photo/repository/record"

type Request struct {
	Title    string `json:"title" binding:"required" example:"holiday"`
	Caption  string `json:"caption" example:"no caption"`
	PhotoUrl string `json:"photo_url" binding:"required" example:"/example"`
}

func (req *Request) MapPostToRecord(userId int) *record.Photo {
	return &record.Photo{
		UserID:   userId,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}
}

func (req *Request) MapUpdateToRecord() *record.Photo {
	return &record.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}
}
