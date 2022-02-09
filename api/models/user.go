package models

import ()

type UserRole string

const (
	RoleUser = UserRole("user")
	RoleAdmin = UserRole("admin")
)

type User struct {
	ID string
	Handle string
	Profile string
	Role UserRole
}
