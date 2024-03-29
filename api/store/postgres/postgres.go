package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB() (*PostgresDB, error) {

	sqlDB, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		))

	if err != nil {
		return nil, errors.Wrap(err, "cant open db")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "gorm db error")
	}

	return &PostgresDB{db: gormDB}, nil
}

// Delete deletes all entries in a table of a writer
// ONLY USE IN TESTS
func (db *PostgresDB) Delete(table string, writer string) error {
	query := fmt.Sprintf("delete from %s where writer = ?", table)
	return db.db.Exec(query, writer).Error
}
