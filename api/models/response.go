package models

type Response struct {
	Success bool `json:"success"`
	Errors []string `json:"errors"`
}
