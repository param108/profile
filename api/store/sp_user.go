package store

import (
	"log"
	"os"

	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) FindOrCreateSPUser(
	phone string, writer string) (*models.SpUser, error) {
	spUser, err := s.db.FindOrCreateSPUser(phone, writer)
	if err != nil {
		return spUser, err
	}

	// Add the user to the yatra group
	// This is a smell, move this out.
	yatraGrpID := os.Getenv("YATRA_GROUP_ID")
	if yatraGrpID == "" {
		return spUser, err
	}

	if _, err := s.db.GetSPGroupUser(spUser.ID, yatraGrpID, writer); err != nil {
		if _, err := s.db.AddSPUserToGroup(&models.SpGroupUser{
			SpGroupID: yatraGrpID,
			SpUserID:  spUser.ID,
			Role:      spUser.Role,
			Writer:    writer,
		}, writer); err != nil {
			log.Printf("Failed to add user to yatra: %s\n", err)
		}
	}

	return spUser, err
}

// GetSPUserByID get a sp User by ID
func (s *StoreImpl) GetSPUserByID(id string, writer string) (*models.SpUser, error) {
	return s.db.GetSPUserByID(id, writer)
}

// UpdateSPUser update a spUser
func (s *StoreImpl) UpdateSPUser(user *models.SpUser) (*models.SpUser, error) {
	return s.db.UpdateSPUser(user)
}
