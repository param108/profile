package postgres

import (
	"errors"
	"time"

	"github.com/param108/profile/api/models"
)

func (db *PostgresDB) GetSPUserMessagesByDay(userID string, start time.Time, limit int,
	writer string) (map[string][]*models.SpGroupMsgSend, error) {
	msgs := []*models.SpMessage{}
	err := db.db.Where(
		"sp_user_id = ? and writer = ? and created_at < ?",
		userID, writer, start).Order("created_at DESC").Limit(limit).Find(&msgs).Error
	if err != nil {
		return nil, err
	}

	if len(msgs) == 0 {
		return nil, errors.New("not found")
	}

	ret := map[string][]*models.SpGroupMsgSend{}
	for _, m := range msgs {
		d := m.CreatedAt.Format("02-Jan-06")
		if _, ok := ret[d]; !ok {
			ret[d] = []*models.SpGroupMsgSend{}
		}
		ret[d] = append(ret[d], &models.SpGroupMsgSend{
			ID:             m.ID,
			SpGroupID:      "",
			SpUserID:       m.SpUserID,
			MsgType:        m.MsgType,
			MsgValue:       m.MsgValue,
			MsgText:        m.MsgText,
			CreatedAt:      m.CreatedAt,
			Writer:         m.Writer,
			SpUserPhotoURL: m.SpUserPhotoURL,
		})
	}
	return nil, nil
}

func (db *PostgresDB) AddSpMessage(
	msg *models.SpMessage, writer string) (*models.SpMessage, error) {

	msg.Writer = writer

	if err := db.db.Create(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}
