package postgres

import (
	"github.com/param108/profile/api/models"
	"github.com/pkg/errors"
)

func (db *PostgresDB) CreateUser(u *models.User) error {
	return db.db.Create(u).Error
}

func (db *PostgresDB) FindOrCreateUser(u *models.User) (*models.User, error) {
	ret := &models.User{}
	if err := db.db.Where(
		"handle = ? and profile = ?",
		u.Handle,
		u.Profile).FirstOrCreate(ret).Error; err != nil {
		return nil, err
	}

	return ret, nil
}

func (db *PostgresDB) GetUser(userID string, writer string) (*models.User, error) {
	user := &models.User{}
	err := db.db.Find(user).Where("id = ? and writer = ?", userID, writer).Error
	return user, err
}

func (db *PostgresDB) GetUserByHandle(username, writer string) (*models.User, error) {
	users := []*models.User{}
	if err := db.db.Where(
		"handle = ? and writer = ?",
		username,
		writer).Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("not found")
	}

	return users[0], nil
}
