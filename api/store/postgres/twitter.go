package postgres

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (db *PostgresDB) CreateTwitterChallenge(token, writer string) (string, error) {
	var id string
	err := db.db.Raw(`INSERT INTO twitter_challenges (challenge, writer)
               VALUES (?,?) RETURNING id`, token, writer).Scan(&id).Error
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *PostgresDB) GetTwitterChallenge(key, writer string) (string, error) {
	ret := &models.TwitterChallenge{}
	if err := db.db.First(ret, "id = ? and writer = ?", key, writer).Error; err != nil {
		return "", err
	}

	return ret.Challenge, nil
}

func (db *PostgresDB) DeleteOldTwitterChallenges(checkpoint time.Duration) {
	db.db.Exec(`delete from twitter_challenges where created_at < ?`, time.Now().Add(-1*checkpoint))
}
