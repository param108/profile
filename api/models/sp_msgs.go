package models

import "time"

type SpMessage struct {
	ID             string    `json:"id" gorm:"default:uuid_generate_v4()"`
	SpUserID       string    `json:"sp_user_id"`
	MsgType        string    `json:"sp_msg_type"`  // The spiritual practice performed
	MsgValue       int       `json:"sp_msg_value"` // The amount of service performed
	MsgText        string    `json:"msg_text"`     // User's realization
	CreatedAt      time.Time `json:"created_at" gorm:"default:NOW()"`
	Writer         string    `json:"writer"`
	SpUserPhotoURL string    `json:"sp_user_photo_url"`
	Username       string    `json:"username"`
	PhotoURL       string    `json:"photo_url" gorm:"default:''"`
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
	Username       string    `json:"username"`
	PhotoURL       string    `json:"photo_url"`
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
