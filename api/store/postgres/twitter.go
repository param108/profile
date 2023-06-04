package postgres

import (
	"time"

	"github.com/param108/profile/api/models"
)

func (db *PostgresDB) CreateTwitterChallenge(token, redirectURI, writer string) (string, error) {
	var id string
	err := db.db.Raw(`INSERT INTO twitter_challenges (challenge, redirect_uri, writer)
               VALUES (?,?,?) RETURNING id`, token, redirectURI, writer).Scan(&id).Error
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *PostgresDB) GetTwitterChallenge(token, writer string) (string, string, error) {
	ret := &models.TwitterChallenge{}
	if err := db.db.First(ret, "id = ? and writer = ?", token, writer).Error; err != nil {
		return "", "", err
	}

	return ret.Challenge, ret.RedirectUri, nil
}

func (db *PostgresDB) DeleteOldTwitterChallenges(checkpoint time.Duration) {
	db.db.Exec(`delete from twitter_challenges where created_at < ?`, time.Now().Add(-1*checkpoint))
}
