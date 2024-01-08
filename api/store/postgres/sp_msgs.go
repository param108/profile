package postgres

import (
	"strings"
	"time"

	"github.com/param108/profile/api/models"
)

func (db *PostgresDB) GetSPUserMessagesByDay(userID string, start time.Time, tz string, limit int,
	writer string) ([]*models.SpGroupMsgData, error) {
	msgs := []*models.SpMessage{}
	err := db.db.Where(
		"sp_user_id = ? and writer = ? and created_at < ?",
		userID, writer, start).Order("created_at DESC").Limit(limit).Find(&msgs).Error
	if err != nil {
		return nil, err
	}

	ret := []*models.SpGroupMsgData{}

	if len(msgs) == 0 {
		return ret, nil
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}

	found := map[string]bool{}

	data := &models.SpGroupMsgData{}

	for _, m := range msgs {
		createdAt := m.CreatedAt.In(loc)

		d := createdAt.Format("Mon,02-Jan-06")
		info := strings.Split(d, ",")

		if _, ok := found[d]; !ok {
			if len(found) != 0 {
				ret = append(ret, data)
				data = &models.SpGroupMsgData{}
			}

			found[d] = true
			data.Date = info[1]
			data.Day = info[0]
		}

		data.Msgs = append(data.Msgs, &models.SpGroupMsgSend{
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

	if len(msgs) > 0 {
		ret = append(ret, data)
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
