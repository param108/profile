package store

import (
	"time"

	"github.com/param108/profile/api/models"
)

// Store This interface encapsulates all storage needs.
// The idea here is that it doesn't matter the implementation of the store.
// What is important is the Data flow.
type Store interface {
	CreateUser(handle string, role models.UserRole, writer string) (*models.User, error)
	GetUser(userID string, writer string) (*models.User, error)
	CreateInvalidToken(token string, writer string) error
	IsInvalidToken(token string, writer string) (bool, error)

	// CreateTwitterChallenge Takes a token and returns a uuid
	// which can be passed to GetTwitterChallenge to retrieve it
	CreateTwitterChallenge(token, writer string) (string, error)

	// GetTwitterChallenge Takes a uuid previously returned by CreateTwitterChallenge
	// and returns the saved challenge string
	GetTwitterChallenge(key, writer string) (string, error)

	// DeleteOldTwitterChallenges Deletes TwitterChallenges > 24 hours old
	DeleteOldTwitterChallenges(d time.Duration)
}
