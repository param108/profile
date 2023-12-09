package store

import (
	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) FindOrCreateSPUser(
	phone string, writer string) (*models.SpUser, error) {
	return s.db.FindOrCreateSPUser(phone, writer)
}

// GetSPUserByID get a sp User by ID
func (s *StoreImpl) GetSPUserByID(id string, writer string) (*models.SpUser, error) {
	return s.db.GetSPUserByID(id, writer)
}

// UpdateSPUser update a spUser
func (s *StoreImpl) UpdateSPUser(user *models.SpUser) (*models.SpUser, error) {
	return s.db.UpdateSPUser(user)
}
