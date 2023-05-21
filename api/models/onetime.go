package models

import (
	"time"
)

type Onetime struct {
	ID        string `gorm:"default:uuid_generate_v4()"`
	Data      string
	CreatedAt time.Time `gorm:"default now()"`
	Writer    string
}
