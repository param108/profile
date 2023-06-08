package store

import (
	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) CreateUser(handle string, profile string, role models.UserRole, writer string) (*models.User, error) {
	user := &models.User{
		Handle:  handle,
		Role:    role,
		Profile: profile,
		Writer:  writer,
	}
	if err := s.db.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *StoreImpl) FindOrCreateUser(
	handle string,
	profile string,
	role models.UserRole,
	writer string) (*models.User, error) {
	user := &models.User{
		Handle:  handle,
		Role:    role,
		Profile: profile,
		Writer:  writer,
	}
	if su, err := s.db.FindOrCreateUser(user); err != nil {
		return nil, err
	} else {
		return su, nil
	}
}

func (s *StoreImpl) GetUser(userID string, writer string) (*models.User, error) {
	return s.db.GetUser(userID, writer)
}

func (s *StoreImpl) GetUserByHandle(username string, writer string) (*models.User, error) {
	return s.db.GetUserByHandle(username, writer)
}
