package store

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) GetSPUserMessagesByDay(userID string, start time.Time,
	tz string, limit int, writer string) ([]*models.SpGroupMsgData, error) {
	return s.db.GetSPUserMessagesByDay(userID, start, tz, limit, writer)
}

func (s *StoreImpl) AddSpMessage(
	msg *models.SpMessage, tz, writer string) (*models.SpGroupMsgData, error) {
	return s.db.AddSpMessage(msg, tz, writer)
}
