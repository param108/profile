package store

import (
	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) FindOrCreateSPUser(
	phone string, writer string) (*models.SpUser, error) {
	return s.db.FindOrCreateSPUser(phone, writer)
}
