package entity

import "time"

type Letter struct {
	ID        int `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	UserID   int
	Feelings []Feeling
}
