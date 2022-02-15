package store

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) CreateInvalidToken(token string, writer string) error {
	invalidToken := &models.InvalidToken{
		Token:     token,
		Writer:    writer,
		CreatedAt: time.Now().UTC(),
	}

	return s.db.CreateInvalidToken(invalidToken)
}

func (s *StoreImpl) IsInvalidToken(token string, writer string) (bool, error) {
	return s.db.IsInvalidToken(token, writer)
}
