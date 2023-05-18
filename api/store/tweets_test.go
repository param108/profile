package store

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tweetWriter = "17165df5-ee3a-4f25-9c9f-3b8f5fbcc5ac"

func tweetTeardown() {
	err := testDB.(*StoreImpl).db.Delete("tweets", tweetWriter)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("tags", tweetWriter)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}

	err = testDB.(*StoreImpl).db.Delete("tweet_tags", tweetWriter)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}

}

func TestInsertTweet(t *testing.T) {
	tweetTeardown()
	// userID for tweet tests
	userID := "80e2663c-3842-431b-886e-ad440bc29850"

	oldTweetID := ""
	secondTweetID := ""
	t.Run("Insert a tweet and a few tags", func(t *testing.T) {
		tw, tags, err := testDB.(*StoreImpl).InsertTweet(userID,
			`#display
The first #tweet is a short
#Hello #World.`, "", tweetWriter)
		assert.Nil(t, err, "failed to insert tweets and tags")
		assert.NotNil(t, tw.CreatedAt, "empty created_at")
		assert.NotEmpty(t, tw.ID, "ID is empty")
		assert.Equal(t, 3, len(tags), "incorrect number of tags")
		for _, tag := range tags {
			assert.NotEmpty(t, tag.ID, "empty ID")
			assert.NotNil(t, tag.CreatedAt, "empty created at")
		}
		oldTweetID = tw.ID
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
		assert.Equal(t, 2, len(tweets))
	})

	t.Run("get tweets after the first one", func(t *testing.T) {
		tweets, err := testDB.GetTweets(userID, 1, 10, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 1, len(tweets))
		// Make sure its the earliest one
		assert.Equal(t, oldTweetID, tweets[0].ID)
	})

	t.Run("delete the first tweet", func(t *testing.T) {
		tweet, err := testDB.DeleteTweet(userID, oldTweetID, tweetWriter)
		assert.Nil(t, err, "failed to delete")
		assert.Equal(t, oldTweetID, tweet.ID, "returned incorrect tweet")
	})

	t.Run("get tweets. Only one should be returned", func(t *testing.T) {
		tweets, err := testDB.GetTweets(userID, 0, 10, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 1, len(tweets))
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

	tweetTeardown()
	t.Run("multiple tags and single user", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			testDB.(*StoreImpl).InsertTweet(userID,
				fmt.Sprintf(`#display
The first tweet is #tweet_%d
#Hello #World.`, i), "", tweetWriter)
		}
		tweets, err := testDB.SearchTweetsByTags(
			userID,
			[]string{"tweet_1", "tweet_3", "tweet_5", "tweet_7"}, tweetWriter)
		assert.Nil(t, err, "failed to get tweets")
		assert.Equal(t, 4, len(tweets))
		suffixes := []int{7, 5, 3, 1}
		for idx, tweet := range tweets {
			// check that the tweet has the correct tag in it
			assert.True(t, strings.Contains(
				tweet.Tweet,
				fmt.Sprintf(
					"tweet_%d",
					suffixes[idx])), "invalid tweet found")
		}
	})

}
