package postgres

import (
	"errors"
	"regexp"

	"github.com/param108/profile/api/models"
)

var validPhoneRe = regexp.MustCompile(`^[0-9]{10}$`)

func (db *PostgresDB) FindOrCreateSPUser(
	phone string, writer string) (*models.SpUser, error) {
	spUser := &models.SpUser{}

	if !validPhoneRe.MatchString(phone) {
		return nil, errors.New("invalid phone")
	}

	// best effort delete any previous otp
	if err := db.db.Where("phone = ? and writer = ?", phone, writer).Assign(
		models.SpUser{
			Phone:  phone,
			Writer: writer,
		}).FirstOrCreate(spUser).Error; err != nil {
		return nil, err
	}
	return spUser, nil
}
