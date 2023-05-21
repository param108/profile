package postgres

import (
	"errors"
	"time"

	"github.com/param108/profile/api/models"
)

// SetOneTime Creates a new onetime table record.
func (db *PostgresDB) SetOneTime(val, writer string) (*models.Onetime, error) {
	onetime := &models.Onetime{
		Data:   val,
		Writer: writer,
	}

	if err := db.db.Table("onetime").Create(onetime).Error; err != nil {
		return nil, err
	}
	return onetime, nil
}

// GetOneTime Returns a one time value if it hasn't expired
func (db *PostgresDB) GetOneTime(
	id string, expiry time.Duration, writer string) (*models.Onetime, error) {
	rets := []*models.Onetime{}
	expired := time.Now().UTC().Add(-expiry)
	if err := db.db.Table("onetime").Where(
		"created_at > ? and id = ? and writer = ?",
		expired, id, writer).Find(&rets).Error; err != nil {
		return nil, err
	}

	if len(rets) == 0 {
		return nil, errors.New("not found")
	}

	return rets[0], nil
}

func (db *PostgresDB) SaveOneTime(onetime *models.Onetime) error {
	return db.db.Table("onetime").Save(onetime).Error
}
