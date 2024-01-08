package store

import "github.com/param108/profile/api/models"

func (s *StoreImpl) GetSPServices(writer string) ([]*models.SpService, error) {
	return s.db.GetSPServices(writer)
}
