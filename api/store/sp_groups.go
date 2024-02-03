package store

import "github.com/param108/profile/api/models"

func (s *StoreImpl) AddSPGroup(group *models.SpGroup, writer string) (*models.SpGroup, error) {
	return s.db.AddSPGroup(group, writer)
}

func (s *StoreImpl) GetSPGroupsForUser(userID, writer string) ([]*models.SpGroupSend, error) {
	return s.db.GetSPGroupsForUser(userID, writer)
}

func (s *StoreImpl) AddSPUserToGroup(
	spGroupUser *models.SpGroupUser, writer string) (*models.SpGroupUser, error) {
	return s.db.AddSPUserToGroup(spGroupUser, writer)
}

func (s *StoreImpl) GetSPGroupUser(
	userID, groupID, writer string) (*models.SpGroupUser, error) {
	return s.db.GetSPGroupUser(userID, groupID, writer)
}

func (s *StoreImpl) GetSPGroupUsers(groupID, writer string) ([]*models.SpUser, error) {
	return s.db.GetSPGroupUsers(groupID, writer)
}
