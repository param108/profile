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

type SpGroupUser struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpGroupID string    `json:"sp_group_id"`
	SpUserID  string    `json:"sp_user_id"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:NOW()"`
	Writer    string    `json:"writer"`
}

type SpMessage struct {
	ID             string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpUserID       string    `json:"sp_user_id"`
	MsgType        string    `json:"sp_msg_type"`  // The spiritual practice performed
	MsgValue       int       `json:"sp_msg_value"` // The amount of service performed
	MsgText        string    `json:"msg_text"`     // User's realization
	CreatedAt      time.Time `json:"created_at" gorm:"default:NOW()"`
	Writer         string    `json:"writer"`
	SpUserPhotoURL string    `json:"sp_user_photo_url"`
}

// Data structure for sending back messages.
// SpGroupID will be empty string if the Group is the user's
// private group.
type SpGroupMsgSend struct {
	ID             string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpGroupID      string    `json:"sp_group_id"`
	SpUserID       string    `json:"sp_user_id"`
	MsgType        string    `json:"sp_msg_type"`  // The spiritual practice performed
	MsgValue       int       `json:"sp_msg_value"` // The amount of service performed
	MsgText        string    `json:"msg_text"`     // User's realization
	CreatedAt      time.Time `json:"created_at" gorm:"default:NOW()"`
	Writer         string    `json:"writer"`
	SpUserPhotoURL string    `json:"sp_user_photo_url"`
}

type SpGroupMsgData struct {
	Day  string            `json:"day"`
	Date string            `json:"date"`
	Msgs []*SpGroupMsgSend `json:"msgs"`
}

type SpMsgComment struct {
	ID          string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpUserID    string    `json:"sp_user_id"`
	SpMessageID string    `json:"sp_message_id"`
	MsgText     string    `json:"msg_text"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:NOW()"`
	Writer      string    `json:"writer"`
}

type SpGroupMessage struct {
	ID          string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpGroupID   string    `json:"sp_group_id"`
	SpUserID    string    `json:"sp_user_id"`
	SpMessageID string    `json:"sp_message_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:NOW()"`
	Writer      string    `json:"writer"`
}

type RefreshTokenResponse struct {
	SpUser       *SpUser `json:"user"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}
