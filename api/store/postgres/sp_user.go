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

func (db *PostgresDB) GetSPUserByID(id string, writer string) (*models.SpUser, error) {
	users := []*models.SpUser{}
	if err := db.db.Where("id = ? and writer = ?", id, writer).Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("not found")
	}

	return users[0], nil
}

func (db *PostgresDB) UpdateSPUser(user *models.SpUser) (*models.SpUser, error) {
	if err := db.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
