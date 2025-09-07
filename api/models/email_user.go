package models

type EmailUser struct {
	ID           string `gorm:"default:uuid_generate_v4();primaryKey"`
	UserName     string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Writer       string `gorm:"not null"`
}