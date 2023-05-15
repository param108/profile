package models

import "time"

type Tweet struct {
	ID string `gorm:"default:uuid_generate_v4()"`
	UserID string
	Tweet string
	Flags string
	Writer string
	CreatedAt time.Time
}

type Tag struct {
	ID string `gorm:"default:uuid_generate_v4()"`
	UserID string
	Tag string
	Writer string
	CreatedAt time.Time
}

type TweetTag struct {
	ID int
	Tag string
	TweetID string
	Writer string
	UserID string
}
