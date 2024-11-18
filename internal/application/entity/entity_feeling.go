package entity

import "time"

type FeelingType string

const (
	Pleasure   FeelingType = "PLEASURE"
	Unpleasant FeelingType = "UNPLEASANT"
)

type Feeling struct {
	ID        int         `gorm:"primaryKey;autoIncrement"`
	Type      FeelingType `sql:"type:ENUM('PLEASURE', 'UNPLEASANT')" gorm:"column:feeling_type"`
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	UserID   int
	LetterID int
}
