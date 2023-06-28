package store

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/param108/profile/api/models"
	"github.com/stretchr/testify/assert"
)

const tweetWriter = "17165df5-ee3a-4f25-9c9f-3b8f5fbcc5ac"
const guestWriter = "75dd2fc8-7417-4558-8da5-0ddce0348c82"

func tweetTeardown(writer string) {
	err := testDB.(*StoreImpl).db.Delete("tweets", writer)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("tags", writer)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}
	err = testDB.(*StoreImpl).db.Delete("tweet_tags", writer)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("thread_tweets", writer)
	if err != nil {
		log.Fatalf("failed delete thread_tweets: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("threads", writer)
	if err != nil {
		log.Fatalf("failed delete threads: %s", err.Error())
	}

}

func TestInsertTweet(t *testing.T) {
	tweetTeardown(tweetWriter)
	// userID for tweet tests
	userID := "80e2663c-3842-431b-886e-ad440bc29850"

	oldTweetID := ""
	secondTweetID := ""
	noTagTweetID := ""

	t.Run("Insert a tweet with no tags", func(t *testing.T) {
		tw, tags, err := testDB.(*StoreImpl).InsertTweet(userID,
			`#display
The first tweet has no tags.
notags.`, "", tweetWriter)
		assert.Nil(t, err, "failed to insert tweets and tags")
		assert.NotNil(t, tw.CreatedAt, "empty created_at")
		assert.NotEmpty(t, tw.ID, "ID is empty")
		assert.Equal(t, 0, len(tags), "incorrect number of tags")
		noTagTweetID = tw.ID
		fmt.Println("INSERT", tw.CreatedAt, tw.Tweet, tw.ID)

	})

	t.Run("Update tweet with no tags", func(t *testing.T) {
		tw, tags, err := testDB.UpdateTweet(userID, noTagTweetID, `
No tweet tag anywhere`, "", tweetWriter)
		assert.Nil(t, err, "failed to update tweet")
		assert.Equal(t, `
No tweet tag anywhere`, tw.Tweet, "invalid updated tweet")
		assert.Zero(t, len(tags), "returned some spurious tags")
		fmt.Println("UPDATED", tw.CreatedAt, tw.Tweet, tw.ID)
	})

	t.Run("Insert a tweet and a few tags", func(t *testing.T) {
		tw, tags, err := testDB.(*StoreImpl).InsertTweet(userID,
			`#display
The first #tweet is a short
#Hello #World.`, "", tweetWriter)
		assert.Nil(t, err, "failed to insert tweets and tags")
		assert.NotNil(t, tw.CreatedAt, "empty created_at")
		assert.NotEmpty(t, tw.ID, "ID is empty")
		assert.Equal(t, 3, len(tags), "incorrect number of tags")
		oldTweetID = tw.ID
		for _, tag := range tags {
			assert.NotEmpty(t, tag.ID, "empty ID")
			assert.NotNil(t, tag.CreatedAt, "empty created at")
		}
	})

	t.Run("Insert a second tweet with repeat tags", func(t *testing.T) {
		tw, tags, err := testDB.(*StoreImpl).InsertTweet(userID,
			`#display
The first #tweet is a short
#Hello #World.`, "", tweetWriter)
		assert.Nil(t, err, "failed to insert tweets and tags")
		assert.NotNil(t, tw.CreatedAt, "empty created_at")
		assert.NotEmpty(t, tw.ID, "ID is empty")

		// A new tweet should have been inserted with a new ID
		assert.NotEqual(t, oldTweetID, tw.ID, "same id returned for tweet")
		secondTweetID = tw.ID
		assert.Equal(t, 3, len(tags), "incorrect number of tags")
		for _, tag := range tags {
			assert.NotEmpty(t, tag.ID, "empty ID")
			assert.NotNil(t, tag.CreatedAt, "empty created at")
		}
	})

	t.Run("Get all the Tags check there are only 3", func(t *testing.T) {
		tags, err := testDB.GetTags(userID, tweetWriter)
		assert.Nil(t, err, "failed to get tags")

		assert.Equal(t, 3, len(tags), "incorrect number of tags")

		foundTags := []string{}
		for _, tag := range tags {
			foundTags = append(foundTags, tag.Tag)
		}

		assert.Equal(t, []string{"hello", "tweet", "world"},
			foundTags, "invalid tweet found")
	})

	t.Run("Check number of tweetTags", func(t *testing.T) {
		tags, err := testDB.GetTweetTags(userID, oldTweetID, tweetWriter)
		assert.Nil(t, err, "failed to get tags")

		assert.Equal(t, 3, len(tags), "incorrect number of tags")

		foundTags := []string{}
		for _, tag := range tags {
			foundTags = append(foundTags, tag.Tag)
		}

		assert.Equal(t, []string{"hello", "tweet", "world"},
			foundTags, "invalid tweet found")

	})

	t.Run("Update tweet with more tags", func(t *testing.T) {
		tw, tags, err := testDB.(*StoreImpl).UpdateTweet(userID,
			oldTweetID,
			`#display
The first #tweet is a short
#Hello #World #Tree.`, "", tweetWriter)
		assert.Nil(t, err, "failed to insert tweets and tags")
		assert.NotNil(t, tw.CreatedAt, "empty created_at")
		assert.NotEmpty(t, tw.ID, "ID is empty")

		// A new tweet should have been inserted with a new ID
		assert.Equal(t, oldTweetID, tw.ID, "same id returned for tweet")
		assert.Equal(t, 4, len(tags), "incorrect number of tags")
		for _, tag := range tags {
			assert.NotEmpty(t, tag.ID, "empty ID")
			assert.NotNil(t, tag.CreatedAt, "empty created at")
		}
	})

	t.Run("Get all the Tags check there are only 4", func(t *testing.T) {
		tags, err := testDB.GetTags(userID, tweetWriter)
		assert.Nil(t, err, "failed to get tags")

		assert.Equal(t, 4, len(tags), "incorrect number of tags")

		foundTags := []string{}
		for _, tag := range tags {
			foundTags = append(foundTags, tag.Tag)
		}

		assert.Equal(t, []string{"hello", "tree", "tweet", "world"},
			foundTags, "invalid tweet found")
	})

	t.Run("Check number of tweetTags", func(t *testing.T) {
		tags, err := testDB.GetTweetTags(userID, oldTweetID, tweetWriter)
		assert.Nil(t, err, "failed to get tags")

		assert.Equal(t, 4, len(tags), "incorrect number of tags")

		foundTags := []string{}
		for _, tag := range tags {
			foundTags = append(foundTags, tag.Tag)
		}

		assert.Equal(t, []string{"hello", "tree", "tweet", "world"},
			foundTags, "invalid tweet found")

	})

	t.Run("Update tweet with less tags", func(t *testing.T) {
		tw, tags, err := testDB.(*StoreImpl).UpdateTweet(userID,
			oldTweetID,
			`#display
The first is a short
#Hello #World #Tree.`, "", tweetWriter)
		assert.Nil(t, err, "failed to insert tweets and tags")
		assert.NotNil(t, tw.CreatedAt, "empty created_at")
		assert.NotEmpty(t, tw.ID, "ID is empty")

		// A new tweet should have been inserted with a new ID
		assert.Equal(t, oldTweetID, tw.ID, "same id returned for tweet")
		assert.Equal(t, 3, len(tags), "incorrect number of tags")
		for _, tag := range tags {
			assert.NotEmpty(t, tag.ID, "empty ID")
			assert.NotNil(t, tag.CreatedAt, "empty created at")
		}
	})

	t.Run("Get all the Tags check there are only 3", func(t *testing.T) {
		tags, err := testDB.GetTags(userID, tweetWriter)
		assert.Nil(t, err, "failed to get tags")

		assert.Equal(t, 4, len(tags), "incorrect number of tags")

		foundTags := []string{}
		for _, tag := range tags {
			foundTags = append(foundTags, tag.Tag)
		}

		assert.Equal(t, []string{"hello", "tree", "tweet", "world"},
			foundTags, "invalid tweet found")
	})

	t.Run("Check number of tweetTags", func(t *testing.T) {
		tags, err := testDB.GetTweetTags(userID, oldTweetID, tweetWriter)
		assert.Nil(t, err, "failed to get tags")

		assert.Equal(t, 3, len(tags), "incorrect number of tags")

		foundTags := []string{}
		for _, tag := range tags {
			foundTags = append(foundTags, tag.Tag)
		}

		assert.Equal(t, []string{"hello", "tree", "world"},
			foundTags, "invalid tweet found")

	})

	t.Run("get tweets", func(t *testing.T) {
		tweets, err := testDB.GetTweets(userID, 0, 10, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 3, len(tweets))
	})

	t.Run("get tweets after the latest one", func(t *testing.T) {
		tweets, err := testDB.GetTweets(userID, 1, 10, tweetWriter)
		fmt.Println(tweets)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 2, len(tweets))
		// Make sure its the earliest one
		assert.Equal(t, oldTweetID, tweets[0].ID)
	})

	t.Run("delete the first tweet", func(t *testing.T) {
		tweet, err := testDB.DeleteTweet(userID, oldTweetID, tweetWriter)
		assert.Nil(t, err, "failed to delete")
		assert.Equal(t, oldTweetID, tweet.ID, "returned incorrect tweet")
	})

	t.Run("get tweets. Only two should be returned", func(t *testing.T) {
		tweets, err := testDB.GetTweets(userID, 0, 10, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 2, len(tweets))
		// Make sure its the second one
		assert.Equal(t, secondTweetID, tweets[0].ID)
	})

	t.Run("Insert 20 tweets, when we get Tweets we should only get the limit mentioned",
		func(t *testing.T) {
			for i := 0; i < 20; i++ {
				testDB.(*StoreImpl).InsertTweet(userID,
					`#display
The first #tweet is a short
#Hello #World.`, "", tweetWriter)
			}
			tweets, err := testDB.GetTweets(userID, 0, 10, tweetWriter)
			assert.Nil(t, err, "failed to get tweets")
			assert.Equal(t, 10, len(tweets))
		})

	tweetTeardown(tweetWriter)
	t.Run("multiple tags and single user", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			testDB.(*StoreImpl).InsertTweet(userID,
				fmt.Sprintf(`#display
The first tweet is #tweet_%d
#Hello #World.`, i), "", tweetWriter)
		}
		tweets, err := testDB.SearchTweetsByTags(
			userID,
			[]string{"tweet_1", "tweet_3", "tweet_5", "tweet_7"}, 20, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 4, len(tweets))

		// tweets will be ordered oldest first
		suffixes := []int{7, 5, 3, 1}
		for idx, tweet := range tweets {
			// check that the tweet has the correct tag in it
			assert.True(t, strings.Contains(
				tweet.Tweet,
				fmt.Sprintf(
					"tweet_%d",
					suffixes[idx])), "invalid tweet found")
		}

		tweets, err = testDB.SearchTweetsByTags(
			userID,
			[]string{"tweet_199"}, 20, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 0, len(tweets))

	})
	tweetTeardown(tweetWriter)

	var threadTweet *models.Tweet
	var threadID string
	t.Run("insert tweet with thread", func(t *testing.T) {
		thread, err := testDB.CreateThread(userID, tweetWriter)
		threadID = thread.ID
		assert.Nil(t, err, "failed to create thread")
		tweetStr := fmt.Sprintf(`#thread:%s:%d
This tweet is part of thread %s`, thread.ID, 0, thread.ID)
		tweet, _, err := testDB.InsertTweet(userID, tweetStr, "", tweetWriter)
		assert.Nil(t, err, "failed insert tweet")
		threadTweet = tweet
		threadData, err := testDB.GetThread(userID, thread.ID, tweetWriter)
		assert.Nil(t, err, "invalid thread")
		assert.Equal(t, 1, len(threadData.Tweets))
		assert.Equal(t, tweet.ID, threadData.Tweets[0].ID, "check tweet returned properly")
		assert.Equal(t, tweetStr, threadData.Tweets[0].Tweet, "invalid saved tweet")
	})

	t.Run("update tweet remove thread", func(t *testing.T) {
		tweetStr := fmt.Sprintf(`
This tweet is not part of thread`)
		_, _, err := testDB.UpdateTweet(userID, threadTweet.ID, tweetStr, "", tweetWriter)
		assert.Nil(t, err, "failed to update tweet")
		threadData, err := testDB.GetThread(userID, threadID, tweetWriter)
		assert.Nil(t, err, "invalid thread")
		assert.Equal(t, 0, len(threadData.Tweets))
	})

	t.Run("update tweet add thread again", func(t *testing.T) {
		tweetStr := fmt.Sprintf(`#thread:%s:%d
This tweet is part of thread %s`, threadID, 0, threadID)
		_, _, err := testDB.UpdateTweet(userID, threadTweet.ID, tweetStr, "", tweetWriter)
		assert.Nil(t, err, "failed to update tweet")
		threadData, err := testDB.GetThread(userID, threadID, tweetWriter)
		assert.Nil(t, err, "invalid thread")
		assert.Equal(t, 1, len(threadData.Tweets))
		assert.Equal(t, threadTweet.ID, threadData.Tweets[0].ID, "check tweet returned properly")
		assert.Equal(t, tweetStr, threadData.Tweets[0].Tweet, "invalid saved tweet")
	})

}

func TestDeleteGuestData(t *testing.T) {
	tweetTeardown(guestWriter)
	t.Run("Insert 20 tweets",
		func(t *testing.T) {
			for i := 0; i < 20; i++ {
				testDB.(*StoreImpl).InsertTweet(models.GuestUserID,
					fmt.Sprintf(`#display
The first #tweet is a short
#Hello #World%d.`, i), "", guestWriter)
			}
			tweets, err := testDB.GetTweets(models.GuestUserID, 0, 40, guestWriter)
			assert.Nil(t, err, "failed to get tweets")
			assert.Equal(t, 20, len(tweets))

			err = testDB.DeleteGuestData(models.GuestUserID, 20, guestWriter)
			assert.Nil(t, err, "couldnt delete")

			tweets, err = testDB.GetTweets(models.GuestUserID, 0, 40, guestWriter)
			assert.Nil(t, err, "failed to get tweets")
			assert.Equal(t, 20, len(tweets))

			err = testDB.DeleteGuestData(models.GuestUserID, 10, guestWriter)
			assert.Nil(t, err, "couldnt delete")

			tweets, err = testDB.GetTweets(models.GuestUserID, 0, 40, guestWriter)
			assert.Nil(t, err, "failed to get tweets")
			assert.Equal(t, 10, len(tweets))

		})

}
