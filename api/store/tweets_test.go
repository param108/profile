package store

import (
	"log"
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

}
