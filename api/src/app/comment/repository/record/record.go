package record

import "time"

type Comment struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	PhotoID   int
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User  `gorm:"constraint:OnDelete:CASCADE"`
	Photo     Photo `gorm:"constraint:OnDelete:CASCADE"`
}

type User struct {
	ID       int `gorm:"primaryKey"`
	Email    string
	Username string
}

type Photo struct {
	ID       int `gorm:"primaryKey"`
	UserID   int
	Title    string
	Caption  string
	PhotoUrl string
}
