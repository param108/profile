package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/param108/profile/api/models"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (db *PostgresDB) CreateUser(u *models.User) error {
	return db.db.Create(u).Error
}

func (db *PostgresDB) GetUser(userID string, writer string) (*models.User, error) {
	user := &models.User{}
	err := db.db.Find(user).Where("id = ? and writer = ?", userID, writer).Error
	return user, err
}

// Delete deletes all entries in a table of a writer
// ONLY USE IN TESTS
func (db *PostgresDB) Delete(table string, writer string) error {
	query := fmt.Sprintf("delete from %s where writer = ?", table)
	return db.db.Exec(query, writer).Error
}

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

		if err := tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(tags).Error; err != nil {
			return err
		}

		if len(tags) > 0 {
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

		return nil
	})

	return tweet, tags, err
}

func (db *PostgresDB) UpdateTweet(
	tweet *models.Tweet,
	tags []*models.Tag) (*models.Tweet, []*models.Tag, error) {
	err := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(tweet).Error; err != nil {
			return err
		}

		if err := tx.Where(
			"tweet_id = ?", tweet.ID).Delete(
			&models.TweetTag{}).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(tags).Error; err != nil {
			return err
		}

		if len(tags) > 0 {
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

		return nil
	})

	return tweet, tags, err

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
	tags []string, writer string) ([]*models.Tweet, error) {

	query := ""
	// arguments to the query are
	// user_id, writer, tags...
	args := []interface{}{userID, writer}
	for _, tag := range tags {
		if len(query) > 0 {
			query = query + " OR"
		}
		query = query + " tag = ?"
		args = append(args, tag)
	}

	if len(query) > 0 {
		query = "AND ( " + query + " )"
	}

	query = "tweets.user_id = ? AND tweets.writer = ? AND tweets.deleted = FALSE " + query

	tweets := []*models.Tweet{}

	if err := db.db.Joins(
		"Join tweet_tags on tweet_tags.tweet_id = tweets.id").Where(
		query, args...).Order("tweets.created_at DESC").Find(&tweets).Error; err != nil {
		return nil, err
	}

	return tweets, nil
}
