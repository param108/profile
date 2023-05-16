package store

import (
	"github.com/param108/profile/api/models"
)

// InsertTweet inserts a tweet and all tags in a transaction
// It will ignore the first line of flags as those are for display only.
// From the rest extract tags and insert them in tweet table
func (s *StoreImpl) InsertTweet(userID string, tweet string,
	flags string, writer string) (*models.Tweet, error) {
	/*tw := &models.Tweet{
		UserID: userID,
		Tweet:  tweet,
		Writer: writer,
		Flags:  flags,
	}

	if err := s.db.CreateTweet(tw); err != nil {
		return nil, err
	}*/
	return nil, nil
}

// UpdateTweet edits a tweet
func (s *StoreImpl) UpdateTweet(tweetID string,
	tweet string, writer string) (*models.Tweet, error) {
	return nil, nil
}

// DeleteTweet deletes a tweet
func (s *StoreImpl) DeleteTweet(tweetID string, writer string) (*models.Tweet, error) {
	return nil, nil
}

// GetTags return all tags for a user
// writer is optional. Empty value is all writers
func (s *StoreImpl) GetTags(userID string, writer string) ([]*models.Tag, error) {
	return nil, nil
}

// SearchTweetsByTags return all tweets for user
// by tag. Return in Chronologically descending order.
// writer is optional. Empty value is all writers.
func (s *StoreImpl) SearchTweetsByTags(userID string,
	tags []string, writer string) ([]*models.Tweet, error) {
	return nil, nil
}
