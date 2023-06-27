package postgres

import (
	"time"

	"github.com/param108/profile/api/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (db *PostgresDB) CreateThread(userID, writer string) (*models.Thread, error) {
	ret := &models.Thread{
		UserID:    userID,
		Writer:    writer,
		CreatedAt: time.Now().UTC(),
		Deleted:   false,
	}
	if err := db.db.Create(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (db *PostgresDB) DeleteThread(userID, threadID, writer string) (*models.Thread, error) {
	ret := []*models.Thread{}
	err := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(
			"user_id = ? AND id = ? AND writer = ?",
			userID, threadID, writer).Find(&ret).Error; err != nil {
			return err
		}

		if len(ret) == 0 {
			return errors.New("not found")
		}

		if err := tx.Delete(ret[0]).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ? AND thread_id = ? AND writer = ?",
			userID, threadID, writer).Delete(&models.ThreadTweet{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return ret[0], nil
}

func (db *PostgresDB) AddTweetToThread(userID, tweetID, threadID, writer string) error {

	tweets := []*models.Tweet{}
	// Check if the tweet and thread are owned by the user.
	if err := db.db.Where(
		"id=? AND user_id=? AND writer = ? AND deleted = FALSE",
		tweetID, userID, writer).Find(&tweets).Error; err != nil {
		return err
	}

	if len(tweets) == 0 {
		return errors.New("not found")
	}

	threads := []*models.Thread{}
	if err := db.db.Where(
		"id = ? AND user_id = ? AND writer = ? AND deleted = FALSE",
		threadID, userID, writer).Find(&threads).Error; err != nil {
		return err
	}

	if len(threads) == 0 {
		return errors.New("not found")
	}

	threadTweet := &models.ThreadTweet{
		ThreadID: threadID,
		UserID:   userID,
		TweetID:  tweetID,
		Writer:   writer,
	}
	return db.db.Create(threadTweet).Error
}

func (db *PostgresDB) DelTweetFromThread(userID, tweetID, threadID, writer string) error {
	return db.db.Where("user_id = ? AND tweet_id = ? AND thread_id = ? AND writer = ?",
		userID, tweetID, threadID, writer).Delete(&models.ThreadTweet{}).Error
}

func (db *PostgresDB) GetThread(userID, threadID, writer string) (*models.ThreadData, error) {
	threads := []*models.Thread{}
	if err := db.db.Where(
		"id = ? AND user_id = ? AND writer = ? AND deleted = FALSE",
		threadID, userID, writer).Find(&threads).Error; err != nil {
		return nil, err
	}

	if len(threads) == 0 {
		return nil, errors.New("not found")
	}

	rawData := []*models.ThreadRawData{}
	if err := db.db.Table("thread_tweets").Joins(
		"left join tweets on thread_tweets.tweet_id = tweets.id ").Joins(
		"left join threads on thread_tweets.thread_id = threads.id").Where(
		"thread_tweets.user_id = ? AND thread_tweets.thread_id = ? "+
			"AND thread_tweets.writer = ? "+
			"AND thread_tweets.deleted = FALSE AND tweets.deleted = FALSE "+
			"AND threads.deleted = FALSE",
		userID, threadID, writer).Select("threads.id as id, threads.user_id as user_id," +
		"threads.created_at as created_at, threads.deleted as deleted, threads.writer as writer," +
		"tweets.id as tweet_id, tweets.user_id as tweet_user_id, " +
		"tweets.tweet as tweet_tweet, tweets.flags as tweet_flags, " +
		"tweets.writer as tweet_writer, tweets.created_at as tweet_created_at ").Order(
		"tweets.created_at ASC",
	).Scan(
		&rawData).Error; err != nil {
		return nil, err
	}

	ret := &models.ThreadData{
		Thread: models.Thread{
			ID:        threads[0].ID,
			UserID:    threads[0].UserID,
			CreatedAt: threads[0].CreatedAt,
			Deleted:   threads[0].Deleted,
			Writer:    threads[0].Writer,
		},
		Tweets: []*models.Tweet{},
	}

	for _, data := range rawData {
		ret.Tweets = append(ret.Tweets, &models.Tweet{
			ID:        data.TweetID,
			UserID:    data.TweetUserID,
			Tweet:     data.TweetTweet,
			Flags:     data.TweetFlags,
			Writer:    data.TweetWriter,
			CreatedAt: data.TweetCreatedAt})
	}
	return ret, nil
}

func (db *PostgresDB) GetThreadTweets(
	userID, threadID, writer string) ([]*models.ThreadTweet, error) {
	ret := []*models.ThreadTweet{}

	if err := db.db.Where("user_id = ? AND thread_id = ? AND writer = ?",
		userID, threadID, writer).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, errors.New("not found")
	}

	return ret, nil
}
