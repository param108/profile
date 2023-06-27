package store

import (
	"fmt"
	"log"
	"testing"

	"github.com/param108/profile/api/models"
	"github.com/stretchr/testify/assert"
)

const threadWriter = "7ed0494d-2740-425f-a006-2c277e2710ec"

func threadTeardown(writer string) {
	err := testDB.(*StoreImpl).db.Delete("thread_tweets", writer)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("threads", writer)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("tags", writer)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("tweet_tags", writer)
	if err != nil {
		log.Fatalf("failed delete tweet_tags: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("tweets", writer)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}

}

func TestThreadCRUD(t *testing.T) {
	threadTeardown(threadWriter)

	userID := "5b2a9860-5898-46e7-9abb-4254ca170b3a"
	var firstThread *models.Thread
	var secondThread *models.Thread

	t.Run("Can create a thread", func(t *testing.T) {
		thread, err := testDB.CreateThread(userID, threadWriter)
		assert.Nil(t, err, "failed create thread")
		assert.Equal(t, userID, thread.UserID, "invalid user id")
		assert.NotEmpty(t, thread.CreatedAt, "invalid created at")
		assert.NotEmpty(t, thread.ID, "invalid thread id")
		firstThread = thread
	})

	t.Run("Can create second thread", func(t *testing.T) {
		thread, err := testDB.CreateThread(userID, threadWriter)
		assert.Nil(t, err, "failed create thread")
		assert.Equal(t, userID, thread.UserID, "invalid user id")
		assert.NotEmpty(t, thread.CreatedAt, "invalid created at")
		assert.NotEmpty(t, thread.ID, "invalid thread id")
		secondThread = thread
		assert.NotEqual(t, firstThread.ID, secondThread.ID, "same thread seen")
	})

	t.Run("Get first  thread", func(t *testing.T) {
		threadData, err := testDB.GetThread(userID, firstThread.ID, threadWriter)
		assert.Nil(t, err, "failed get thread")
		assert.Equal(t, userID, threadData.UserID, "invalid user_id")
		assert.Equal(t, firstThread.ID, threadData.ID, "invalid id")
		assert.Zero(t, len(threadData.Tweets), "non empty tweets")
	})

	t.Run("Get second thread", func(t *testing.T) {
		threadData, err := testDB.GetThread(userID, secondThread.ID, threadWriter)
		assert.Nil(t, err, "failed get thread")
		assert.Equal(t, userID, threadData.UserID, "invalid user_id")
		assert.Equal(t, secondThread.ID, threadData.ID, "invalid id")
		assert.Zero(t, len(threadData.Tweets), "non empty tweets")
	})

	firstThreadTweets := []string{}

	t.Run("Add 5 tweets to thread", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			tweet, _, _ := testDB.(*StoreImpl).InsertTweet(userID,
				fmt.Sprintf(`#display
The first #tweet is a short
#Hello #World%d.`, i), "", threadWriter)
			firstThreadTweets = append(firstThreadTweets, tweet.ID)
			err := testDB.AddTweetToThread(userID, tweet.ID, firstThread.ID, threadWriter)
			assert.Nil(t, err, "Failed to add Tweet")
		}
	})

	secondThreadTweets := []string{}

	t.Run("Add 5 tweets to second thread", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			tweet, _, _ := testDB.(*StoreImpl).InsertTweet(userID,
				fmt.Sprintf(`#display
The first #tweet is a short
#Hello #World%d.`, i), "", threadWriter)
			secondThreadTweets = append(secondThreadTweets, tweet.ID)
			err := testDB.AddTweetToThread(userID, tweet.ID, secondThread.ID, threadWriter)
			assert.Nil(t, err, "Failed to add Tweet")
		}
	})

	t.Run("Get first  thread with tweets", func(t *testing.T) {
		threadData, err := testDB.GetThread(userID, firstThread.ID, threadWriter)
		assert.Nil(t, err, "failed get thread")
		assert.Equal(t, userID, threadData.UserID, "invalid user_id")
		assert.Equal(t, firstThread.ID, threadData.ID, "invalid id")
		assert.Equal(t, 5, len(threadData.Tweets), "non empty tweets")
	})

	t.Run("Get second thread", func(t *testing.T) {
		threadData, err := testDB.GetThread(userID, secondThread.ID, threadWriter)
		assert.Nil(t, err, "failed get thread")
		assert.Equal(t, userID, threadData.UserID, "invalid user_id")
		assert.Equal(t, secondThread.ID, threadData.ID, "invalid id")
		assert.Equal(t, 5, len(threadData.Tweets), "non empty tweets")
	})

	t.Run("Delete a tweet from a thread", func(t *testing.T) {
		err := testDB.DelTweetFromThread(userID, secondThreadTweets[0], secondThread.ID, threadWriter)
		assert.Nil(t, err, "failed to delete tweet")

		threadData, err := testDB.GetThread(userID, secondThread.ID, threadWriter)
		assert.Equal(t, 4, len(threadData.Tweets), "invalid number of tweets")
	})

	t.Run("Delete a thread and ThreadTweets should go", func(t *testing.T) {
		_, err := testDB.DeleteThread(userID, secondThread.ID, threadWriter)
		assert.Nil(t, err, "cant delete")

		_, err = testDB.(*StoreImpl).db.GetThreadTweets(userID, secondThread.ID, threadWriter)
		assert.NotNil(t, err, "no error")
	})

}
