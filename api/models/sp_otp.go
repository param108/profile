package models

import "time"

type SpOtp struct {
	ID     string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Phone  string    `json:"phone"`
	Code   string    `json:"code"`
	Expiry time.Time `json:"expiry"`
	Writer string    `json:"writer"`
}
