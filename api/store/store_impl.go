package store

import (
	"github.com/param108/profile/api/store/postgres"
)

type StoreImpl struct {
	db *postgres.PostgresDB
}

// NewStore Creates a new Store by initializing
// A StoreImpl object.
func NewStore() (Store, error) {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	return &StoreImpl{
		db: db,
	}, nil
}
