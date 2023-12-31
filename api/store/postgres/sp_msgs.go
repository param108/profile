package postgres

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (db *PostgresDB) GetSPUserMessagesByDay(userID string, start time.Time, tz string, limit int,
	writer string) (map[string][]*models.SpGroupMsgSend, error) {
	msgs := []*models.SpMessage{}
	err := db.db.Where(
		"sp_user_id = ? and writer = ? and created_at < ?",
		userID, writer, start).Order("created_at DESC").Limit(limit).Find(&msgs).Error
	if err != nil {
		return nil, err
	}

	ret := map[string][]*models.SpGroupMsgSend{}

	if len(msgs) == 0 {
		return ret, nil
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}

	for _, m := range msgs {
		createdAt := m.CreatedAt.In(loc)

		d := createdAt.Format("Mon,02-Jan-06")
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
			CreatedAt:      createdAt,
			Writer:         m.Writer,
			SpUserPhotoURL: m.SpUserPhotoURL,
		})
	}
	return ret, nil
}

func (db *PostgresDB) AddSpMessage(
	msg *models.SpMessage, writer string) (*models.SpMessage, error) {

	msg.Writer = writer

	if err := db.db.Create(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}
