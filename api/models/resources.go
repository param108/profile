package models

type Resource struct {
	ID     string `json:"id" gorm:"default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	T      string `json:"t"`
	Value  int    `json:"value"`
	Max    int    `json:"max"`
	Writer string `json:"writer"`
}
