package models

import "time"

type SpUser struct {
	ID             string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Phone          string    `json:"phone"`
	Name           string    `json:"name"`
	PhotoURL       string    `json:"photo_url"`
	CreatedAt      time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"default:NOW()"`
	Deleted        bool      `json:"deleted"`
	Writer         string    `json:"writer"`
	DeletedAt      time.Time `json:"deleted_at"`
	ProfileUpdated bool      `json:"profile_updated"`
}

type SpGroup struct {
	ID          string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Parent      string    `json:"parent"`
	Deleted     bool      `json:"deleted"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:NOW()"`
	DeletedAt   time.Time `json:"deleted_at"`
	Writer      string    `json:"writer"`
}

type SpGroupSend struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Parent      string    `json:"parent"`
	Deleted     bool      `json:"deleted"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Writer      string    `json:"writer"`
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
}

type SpGroupUser struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpGroupID string    `json:"sp_group_id"`
	SpUserID  string    `json:"sp_user_id"`
	Role      string    `json:"role"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:NOW()"`
	Writer    string    `json:"writer"`
}

type RefreshTokenResponse struct {
	SpUser       *SpUser `json:"user"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}
