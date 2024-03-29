package models

import "time"

type SpOtp struct {
	ID      string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Phone   string    `json:"phone"`
	Code    string    `json:"code"`
	Expiry  time.Time `json:"expiry"`
	Writer  string    `json:"writer"`
	Retries int       `json:"retries"`
}

type CreateOTPRequest struct {
	Phone  string `json:"phone"`
	APIKey string `json:"api_key"`
}

type CheckOTPRequest struct {
	Phone    string `json:"phone"`
	Code     string `json:"code"`
	APIKey   string `json:"api_key"`
	IsSignUp bool   `json:"is_sign_up"`
}

type CheckOTPResponse struct {
	SpUser       *SpUser `json:"user"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}
