package store

import (
	"errors"
	"strings"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/utils"
)

// InsertTweet inserts a tweet and all tags in a transaction
// It will ignore the first line of flags as those are for display only.
// From the rest extract tags and insert them in tweet table
func (s *StoreImpl) InsertTweet(userID string, tweet string,
	flags string, writer string) (*models.Tweet, []*models.Tag, error) {
	tw := &models.Tweet{
		UserID: userID,
		Tweet:  tweet,
		Writer: writer,
		Flags:  flags,
	}

	threads, err := utils.ExtractThreads(flags)
	if err != nil {
		return nil, nil, err
	}

	tagStrs, err := utils.ExtractTags(tweet)
	if err != nil {
		return nil, nil, err
	}

	tags := []*models.Tag{}
	for _, tagStr := range tagStrs {
		tag := &models.Tag{
			UserID: userID,
			Tag:    strings.ToLower(tagStr),
			Writer: writer,
		}
		tags = append(tags, tag)
	}

	twD, tags, err := s.db.InsertTweet(tw, tags, threads)
	if err != nil {
		return nil, nil, err
	}

	return twD, tags, nil
}

// UpdateTweet edits a tweet
func (s *StoreImpl) UpdateTweet(userID, tweetID,
	tweet, flags, writer string) (*models.Tweet, []*models.Tag, error) {
	tw := &models.Tweet{
		ID:     tweetID,
		UserID: userID,
		Tweet:  tweet,
		Writer: writer,
		Flags:  flags,
	}

	threads, err := utils.ExtractThreads(flags)
	if err != nil {
		return nil, nil, err
	}

	tagStrs, err := utils.ExtractTags(tweet)
	if err != nil {
		return nil, nil, err
	}

	tags := []*models.Tag{}
	for _, tagStr := range tagStrs {
		if len(strings.TrimSpace(tagStr)) == 0 {
			continue
		}
		tag := &models.Tag{
			UserID: userID,
			Tag:    strings.ToLower(tagStr),
			Writer: writer,
		}
		tags = append(tags, tag)
	}

	twD, tags, err := s.db.UpdateTweet(tw, tags, threads, writer)
	if err != nil {
		return nil, nil, err
	}

	return twD, tags, nil

}

// GetTags return all tags for a user
// writer is optional. Empty value is all writers
func (s *StoreImpl) GetTags(userID, writer string) ([]*models.Tag, error) {
	if len(userID) == 0 {
		return nil, errors.New("user_id mandatory")
	}
	return s.db.GetTags(userID, writer)
}

// SearchTweetsByTags return all tweets for user
// by tag. Return in Chronologically descending order.
// writer is optional. Empty value is all writers.
func (s *StoreImpl) SearchTweetsByTags(userID string,
	tags []string, offset, limit int, reverse bool, writer string) ([]*models.Tweet, error) {
	return s.db.SearchTweetsByTags(userID, tags, offset, limit, reverse, writer)
}

// GetTweetTags return all tags for a user
// writer is optional. Empty value is all writers
func (s *StoreImpl) GetTweetTags(userID, tweetID, writer string) ([]*models.Tag, error) {
	if len(userID) == 0 || len(tweetID) == 0 {
		return nil, errors.New("user_id mandatory, tweetID mandatory")
	}
	return s.db.GetTweetTags(userID, tweetID, writer)
}

func (s *StoreImpl) GetTweets(userID string, offset, limit int, reverse bool, writer string) ([]*models.Tweet, error) {
	return s.db.GetTweets(userID, offset, limit, reverse, writer)
}

func (s *StoreImpl) GetTweet(userID string, tweetID string, writer string) (*models.Tweet, error) {
	return s.db.GetTweet(userID, tweetID, writer)
}

// DeleteTweet soft deletes a tweet
func (s *StoreImpl) DeleteTweet(userID string, tweetID string, writer string) (*models.Tweet, error) {
	return s.db.DeleteTweet(userID, tweetID, writer)
}

func (s *StoreImpl) DeleteGuestData(userID string, maxTweets int,
	writer string) error {
	return s.db.DeleteGuestData(userID, maxTweets, writer)
}

func (s *StoreImpl) UnsafeGetAllTweets(writer string, offset int, count int) ([]*models.Tweet, int, error) {
	return s.db.GetAllTweets(writer, offset, count)
}

func (s *StoreImpl) UnsafeDelete(table, writer string) error {
	return s.db.Delete(table, writer)
}
