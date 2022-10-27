package record

import "time"

type Photo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"constraint:OnDelete:CASCADE"`
}

type User struct {
	ID       int `gorm:"primaryKey"`
	Email    string
	Username string
}
