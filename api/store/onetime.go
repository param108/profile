package store

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) SetOneTime(val, writer string) (*models.Onetime, error) {
	return s.db.SetOneTime(val, writer)
}

func (s *StoreImpl) GetOneTime(
	id string,
	expiry time.Duration,
	writer string) (*models.Onetime, error) {
	return s.db.GetOneTime(id, expiry, writer)
}

func (s *StoreImpl) DeleteOldOneTimes(expiry time.Duration, writer string) error {
	return s.db.DeleteOldOneTimes(expiry, writer)
}
