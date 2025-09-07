package store

import (
	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) CreateEmailUser(username, passwordHash, writer string) (*models.EmailUser, error) {
	return s.db.CreateEmailUser(username, passwordHash, writer)
}

func (s *StoreImpl) GetEmailUserByUserName(username, writer string) (*models.EmailUser, error) {
	return s.db.GetEmailUserByUserName(username, writer)
}