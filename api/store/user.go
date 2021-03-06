package store

import (
	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) CreateUser(handle string, role models.UserRole, writer string) (*models.User, error) {
	user := &models.User{
		Handle: handle,
		Role:   role,
		Writer: writer,
	}
	if err := s.db.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *StoreImpl) GetUser(userID string, writer string) (*models.User, error) {
	return s.db.GetUser(userID, writer)
}
