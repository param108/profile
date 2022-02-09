package store

import (
	"github.com/param108/profile/api/models"
)


func (s *StoreImpl) CreateUser(handle string, role models.UserRole) (*models.User, error) {
	user := &models.User{
		Handle: handle,
		Role: role,
	}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
