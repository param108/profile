package store

import (
	"github.com/param108/profile/api/models"
)

// Store This interface encapsulates all storage needs.
// The idea here is that it doesn't matter the implementation of the store.
// What is important is the Data flow.
type Store interface {
	CreateUser(handle string, role models.UserRole, writer string) (*models.User, error)
	GetUser(userID string, writer string) (*models.User, error)
}
