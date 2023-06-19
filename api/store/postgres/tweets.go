package postgres

import (
	"strings"

	"github.com/param108/profile/api/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// InsertTweet Inserts the tweet and all the tags and
// creates the necessary tweet tags as well.
func (db *PostgresDB) InsertTweet(
	tweet *models.Tweet,
	tags []*models.Tag,
) (*models.Tweet, []*models.Tag, error) {

	err := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(tweet).Error; err != nil {
			return err
		}

		if len(tags) > 0 {
			if err := tx.Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(tags).Error; err != nil {
				return err
			}

			query := ""
			tagArray := []interface{}{}

			tagIDNil := false
			// In case of conflicts we should update the ID in
			// the responses with the values found in the db
			for _, tag := range tags {
				if len(query) != 0 {
					query = query + "or "
				}

				// If atleast one tagID is nil then we need
				// to query to find it's ID
				if len(tag.ID) == 0 {
					tagIDNil = true
				} else {
					continue
				}

				query = query + "tag = ? "
				tagArray = append(tagArray, tag.Tag)
			}

			// Only do below if atleast one TagID is Null.
			if tagIDNil {
				foundTags := []*models.Tag{}
				if err := tx.Where(query, tagArray...).Find(&foundTags).Error; err != nil {
					return err
				}

				tagMap := map[string]string{}
				for _, tag := range foundTags {
					tagMap[tag.Tag] = tag.ID
				}

				for idx := range tags {
					if id, ok := tagMap[tags[idx].Tag]; ok {
						tags[idx].ID = id
					}
				}
			}

			// Finally connect tags to tweets using tweet_tags
			tweetTags := []*models.TweetTag{}
			for _, tag := range tags {
				tweetTag := &models.TweetTag{
					Tag:     tag.Tag,
					TweetID: tweet.ID,
					Writer:  tweet.Writer,
					UserID:  tweet.UserID,
				}

				tweetTags = append(tweetTags, tweetTag)
			}

			if err := tx.Create(tweetTags).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return tweet, tags, err
}

func (db *PostgresDB) UpdateTweet(
	tweet *models.Tweet,
	tags []*models.Tag,
	writer string,
) (*models.Tweet, []*models.Tag, error) {
	err := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Tweet{}).
			Where("id = ? and user_id = ? and writer = ?", tweet.ID, tweet.UserID, writer).
			Update("tweet", tweet.Tweet).Error; err != nil {
			return err
		}

		if err := tx.Where(
			"tweet_id = ? AND user_id = ? AND writer = ?", tweet.ID, tweet.UserID,
			writer).Delete(
			&models.TweetTag{}).Error; err != nil {
			return err
		}

		if len(tags) > 0 {
			if err := tx.Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(tags).Error; err != nil {
				return err
			}

			query := ""
			tagArray := []interface{}{}

			tagIDNil := false
			// In case of conflicts we should update the ID in
			// the responses with the values found in the db
			for _, tag := range tags {
				if len(query) != 0 {
					query = query + "or "
				}

				// If atleast one tagID is nil then we need
				// to query to find it's ID
				if len(tag.ID) == 0 {
					tagIDNil = true
				} else {
					continue
				}

				query = query + "tag = ? "
				tagArray = append(tagArray, tag.Tag)
			}

			// Only do below if atleast one TagID is Null.
			if tagIDNil {
				foundTags := []*models.Tag{}
				if err := tx.Where(query, tagArray...).Find(&foundTags).Error; err != nil {
					return err
				}

				tagMap := map[string]string{}
				for _, tag := range foundTags {
					tagMap[tag.Tag] = tag.ID
				}

				for idx := range tags {
					if id, ok := tagMap[tags[idx].Tag]; ok {
						tags[idx].ID = id
					}
				}
			}

			// Finally connect tags to tweets using tweet_tags
			tweetTags := []*models.TweetTag{}
			for _, tag := range tags {
				tweetTag := &models.TweetTag{
					Tag:     tag.Tag,
					TweetID: tweet.ID,
					Writer:  tweet.Writer,
					UserID:  tweet.UserID,
				}

				tweetTags = append(tweetTags, tweetTag)
			}

			if err := tx.Create(tweetTags).Error; err != nil {
				return err
			}
		}
		return nil
	})

	ret, err := db.GetRawTweet(tweet.UserID, tweet.ID, writer)
	if err != nil {
		return nil, tags, err
	}
	return ret, tags, err

}

func (db *PostgresDB) GetTags(userID, writer string) ([]*models.Tag, error) {
	tags := []*models.Tag{}
	if len(writer) == 0 {
		if err := db.db.Where(
			"user_id = ?", userID).Order("tag asc").Find(&tags).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.db.Where(
			"user_id = ? and writer = ?", userID, writer).Order("tag asc").Find(&tags).Error; err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func (db *PostgresDB) GetTweetTags(userID, tweetID, writer string) ([]*models.Tag, error) {
	tags := []*models.Tag{}
	if len(writer) == 0 {
		if err := db.db.Joins("JOIN tweet_tags ON tags.tag = tweet_tags.tag").Where(
			"tweet_tags.user_id = ? AND tweet_tags.tweet_id = ? AND tags.user_id = ?",
			userID, tweetID, userID).Order("tag asc").Find(&tags).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.db.Joins(
			`JOIN tweet_tags on ( tags.tag = tweet_tags.tag and
tags.user_id = tweet_tags.user_id )`).Where(
			`tweet_tags.user_id = ? and tweet_tags.tweet_id = ? and tags.user_id = ?
and tweet_tags.writer = ? and tags.writer = ?`,
			userID, tweetID,
			userID, writer,
			writer).Order("tag asc").Find(&tags).Error; err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func (db *PostgresDB) GetTweet(userID, tweetID, writer string) (*models.Tweet, error) {
	tweets := []*models.Tweet{}
	if err := db.db.Where(
		"user_id = ? AND id = ? AND writer = ? AND deleted = false",
		userID, tweetID, writer).Find(tweets).Error; err != nil {
		return nil, err
	}

	if len(tweets) == 0 {
		return nil, errors.New("tweet not found")
	}
	return tweets[0], nil
}

// GetRawTweet Ignore the delete flag and return deleted flags.
func (db *PostgresDB) GetRawTweet(userID, tweetID, writer string) (*models.Tweet, error) {
	tweet := &models.Tweet{}
	if err := db.db.Where(
		"user_id = ? AND id = ? AND writer = ?",
		userID, tweetID, writer).First(tweet).Error; err != nil {
		return nil, err
	}

	return tweet, nil
}

func (db *PostgresDB) GetTweets(userID string,
	offset, limit int,
	writer string) ([]*models.Tweet, error) {
	tweets := []*models.Tweet{}
	if err := db.db.Where(
		"user_id = ? AND writer = ? AND deleted = false",
		userID, writer).Order("created_at desc").Offset(offset).Limit(limit).Find(&tweets).Error; err != nil {
		return nil, err
	}

	return tweets, nil
}

func (db *PostgresDB) DeleteTweet(userID, tweetID, writer string) (*models.Tweet, error) {
	if err := db.db.Table("tweets").Where(
		"user_id = ? AND id = ? AND writer = ?",
		userID, tweetID, writer).Update("deleted", true).Error; err != nil {
		return nil, err
	}

	return db.GetRawTweet(userID, tweetID, writer)
}

func (db *PostgresDB) SearchTweetsByTags(userID string,
	tags []string, limit int, writer string) ([]*models.Tweet, error) {

	query := ""
	// arguments to the query are
	// user_id, writer, tags...
	args := []interface{}{userID, writer}
	for _, tag := range tags {
		if len(query) > 0 {
			query = query + " OR"
		}
		query = query + " tag = ?"
		args = append(args, strings.ToLower(tag))
	}

	if len(query) > 0 {
		query = "AND ( " + query + " )"
	}

	query = "tweets.user_id = ? AND tweets.writer = ? AND tweets.deleted = FALSE " + query

	tweets := []*models.Tweet{}

	if err := db.db.Joins(
		"Join tweet_tags on tweet_tags.tweet_id = tweets.id").Where(
		query, args...).Order("tweets.created_at DESC").Limit(limit).Find(&tweets).Error; err != nil {
		return nil, err
	}

	return tweets, nil
}

// DeleteGuestData delete tweets if number of tweets > maxTweets
// Delete oldest ones first. Will not touch tags
func (db *PostgresDB) DeleteGuestData(
	userID string,
	maxTweets int,
	writer string) error {

	var totalTweets int64
	// First check how many tweets are there
	if err := db.db.Model(&models.Tweet{}).
		Where("user_id = ? and writer = ?", userID, writer).
		Count(&totalTweets).Error; err != nil {
		return err
	}

	if totalTweets <= int64(maxTweets) {
		return nil
	}

	tweetsToDel := []models.Tweet{}
	// We have more than minTweets
	// Get the Tweets that don't belong to the most recent minTweets
	// and delete them. Hard Delete.
	if err := db.db.Model(&models.Tweet{}).
		Where("user_id = ? and writer = ?", userID, writer).
		Order("tweets.created_at DESC").
		Offset(maxTweets).Find(&tweetsToDel).Error; err != nil {
		return err
	}

	seenTweet := map[string]bool{}
	tweetIDs := []string{}
	for _, tweet := range tweetsToDel {
		if v, ok := seenTweet[tweet.ID]; ok && v {
			continue
		}
		tweetIDs = append(tweetIDs, tweet.ID)
		seenTweet[tweet.ID] = true
	}

	return db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("tweet_id IN ?", tweetIDs).Delete(&models.TweetTag{}).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err := tx.Where("id IN ?", tweetIDs).Delete(&models.Tweet{}).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		return nil
	})
}
