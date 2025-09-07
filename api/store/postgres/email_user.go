package postgres

import (
	"github.com/param108/profile/api/models"
)

func (pdb *PostgresDB) CreateEmailUser(username, passwordHash, writer string) (*models.EmailUser, error) {
	emailUser := &models.EmailUser{
		UserName:     username,
		PasswordHash: passwordHash,
		Writer:       writer,
	}

	if err := pdb.db.Create(emailUser).Error; err != nil {
		return nil, err
	}

	return emailUser, nil
}

func (pdb *PostgresDB) GetEmailUserByUserName(username, writer string) (*models.EmailUser, error) {
	var emailUser models.EmailUser
	
	if err := pdb.db.Where("user_name = ? AND writer = ?", username, writer).First(&emailUser).Error; err != nil {
		return nil, err
	}

	return &emailUser, nil
}