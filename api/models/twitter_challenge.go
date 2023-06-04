package models

import "time"

type TwitterChallenge struct {
	ID          string `gorm:"default:uuid_generate_v4()"`
	Challenge   string
	Writer      string
	CreatedAt   time.Time
	RedirectUri string
}
