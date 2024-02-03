package postgres

import (
	"strings"
	"time"

	"github.com/param108/profile/api/models"
	"gorm.io/gorm"
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

		d := createdAt.Format("Mon,02-Jan-2006")
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
			Username:       m.Username,
			PhotoURL:       m.PhotoURL,
		})
	}

	if len(msgs) > 0 {
		ret = append(ret, data)
	}

	return ret, nil
}

func (db *PostgresDB) AddSpMessage(
	msg *models.SpMessage, tz, writer string) (*models.SpGroupMsgData, error) {
	ret := &models.SpGroupMsgData{}
	msg.Writer = writer

	err := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(msg).Error; err != nil {
			return err
		}

		spGroups := []*models.SpGroup{}

		if err := tx.Where(
			"sp_user_id = ? and writer = ?",
			msg.SpUserID, writer).Find(&spGroups).Error; err != nil {
			return err
		}

		for _, group := range spGroups {
			if err := tx.Create(&models.SpGroupMessage{
				SpGroupID:   group.ID,
				SpUserID:    msg.SpUserID,
				SpMessageID: msg.ID,
				CreatedAt:   msg.CreatedAt,
				Writer:      msg.Writer,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}

	createdAt := msg.CreatedAt.In(loc)

	d := createdAt.Format("Mon,02-Jan-2006")
	info := strings.Split(d, ",")

	ret.Date = info[1]
	ret.Day = info[0]

	ret.Msgs = append(ret.Msgs, &models.SpGroupMsgSend{
		ID:             msg.ID,
		SpGroupID:      "",
		SpUserID:       msg.SpUserID,
		MsgType:        msg.MsgType,
		MsgValue:       msg.MsgValue,
		MsgText:        msg.MsgText,
		CreatedAt:      createdAt,
		Writer:         msg.Writer,
		SpUserPhotoURL: msg.SpUserPhotoURL,
		Username:       msg.Username,
		PhotoURL:       msg.PhotoURL,
	})

	return ret, nil
}
