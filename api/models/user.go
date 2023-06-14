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

const GuestUserID = "85f724d2-0276-4b8c-aa00-529a08333cea"
const GuestUsername = "guest"

const DevUserID = "a592e6ab-91d1-49a7-9435-ab3c04f77ab9"
const DevUsername = "param108"
