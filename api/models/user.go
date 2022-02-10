package models

type UserRole string

const (
	RoleUser  = UserRole("user")
	RoleAdmin = UserRole("admin")
)

type User struct {
	ID      string `gorm:"default:uuid_generate_v4()"`
	Handle  string
	Profile string
	Role    UserRole
	Writer  string
}
