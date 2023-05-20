package common

import (
	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
)

func FindOrCreateTPUser(db store.Store, handle, profile, writer string) (*models.User, error) {
	return db.FindOrCreateUser(handle, profile, models.RoleUser, writer)
}
