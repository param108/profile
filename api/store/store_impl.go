package store

import (
	"gorm.io/gorm"
)

type StoreImpl struct {
	db *gorm.DB
}

// NewStore Creates a new Store by initializing
// A StoreImpl object.
func NewStore() (Store, error) {
	db, err := NewPostgresDB()
	if err != nil {
		return nil, err
	}

	return &StoreImpl{
		db: db,
	}, nil
}
