package store

import (
	"errors"
	"time"
)

func (s *StoreImpl) CreateTwitterChallenge(token, writer string) (string, error) {
	if len(token) == 0 {
		return "", errors.New("invalid challenge")
	}

	if key, err := s.db.CreateTwitterChallenge(token, writer); err != nil {
		return "", err
	} else {
		return key, nil
	}

}

func (s *StoreImpl) GetTwitterChallenge(key, writer string) (string, error) {
	if token, err := s.db.GetTwitterChallenge(key, writer); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func (s *StoreImpl) DeleteOldTwitterChallenges(checkpoint time.Duration) {
	s.db.DeleteOldTwitterChallenges(checkpoint)
}
