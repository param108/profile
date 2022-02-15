package postgres

import (
	"github.com/param108/profile/api/models"
	"gorm.io/gorm"
)

func (db *PostgresDB) CreateInvalidToken(it *models.InvalidToken) error {
	return db.db.Create(it).Error
}

func (db *PostgresDB) IsInvalidToken(token string, writer string) (bool, error) {
	it := &models.InvalidToken{}

	if err := db.db.First(it, "token = ?", token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	// No error so we found one!
	return true, nil
}
