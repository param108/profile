package models

import (
	"time"
)

type InvalidToken struct {
	Token     string
	Writer    string
	CreatedAt time.Time
}
