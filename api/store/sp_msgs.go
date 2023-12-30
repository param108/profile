package store

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (s *StoreImpl) GetSPUserMessagesByDay(userID string, start time.Time, limit int,
	writer string) (map[string][]*models.SpGroupMsgSend, error) {
	return s.db.GetSPUserMessagesByDay(userID, start, limit, writer)
}

func (s *StoreImpl) AddSpMessage(
	msg *models.SpMessage, writer string) (*models.SpMessage, error) {
	return s.db.AddSpMessage(msg, writer)
}
