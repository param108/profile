package store

import "github.com/param108/profile/api/models"

func (s *StoreImpl) CreateThread(userID, name, writer string) (*models.Thread, error) {
	return s.db.CreateThread(userID, name, writer)
}

func (s *StoreImpl) DeleteThread(userID, threadID, writer string) (*models.Thread, error) {
	return s.db.DeleteThread(userID, threadID, writer)
}

func (s *StoreImpl) AddTweetToThread(userID, tweetID, threadID, writer string) error {
	return s.db.AddTweetToThread(userID, tweetID, threadID, writer)
}

func (s *StoreImpl) DelTweetFromThread(userID, tweetID, threadID, writer string) error {
	return s.db.DelTweetFromThread(userID, tweetID, threadID, writer)
}

func (s *StoreImpl) GetThread(userID, threadID, writer string) (*models.ThreadData, error) {
	return s.db.GetThread(userID, threadID, writer)
}
