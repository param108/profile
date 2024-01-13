package models

type SpService struct {
	ID               string `json:"id" gorm:"default:uuid_generate_v4()"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Unit             string `json:"unit"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Question         string `json:"question"`
	PhotoURL         string `json:"photo_url"`
}
