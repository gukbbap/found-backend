package entity

import "time"

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	UID       string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	Letters  []Letter
	Feelings []Feeling
}
