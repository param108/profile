package cmd

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const toFlagsWriter = "cc0e4ba0-0030-4108-8f92-1b8ea1b47f28"
const toFlagsUserID = "dc6c8456-f27d-4bf7-9659-8f12a4916bab"

func toFlagsTeardown(writer string) {
	err := testDB.UnsafeDelete("tweets", writer)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}

	err = testDB.UnsafeDelete("threads", writer)
	if err != nil {
		log.Fatalf("failed delete threads: %s", err.Error())
	}

	err = testDB.UnsafeDelete("thread_tweets", writer)
	if err != nil {
		log.Fatalf("failed delete thread_tweets: %s", err.Error())
	}

	err = testDB.UnsafeDelete("tags", writer)
	if err != nil {
		log.Fatalf("failed delete tags: %s", err.Error())
	}

	err = testDB.UnsafeDelete("tweet_tags", writer)
	if err != nil {
		log.Fatalf("failed delete tweet_tags: %s", err.Error())
	}

}

func TestToFlags(t *testing.T) {
	toFlagsTeardown(toFlagsWriter)

	t.Run("Convert tweets", func(t *testing.T) {
		th, _ := testDB.CreateThread(toFlagsUserID, "to flags test thread", toFlagsWriter)
		// font display tag
		testDB.InsertTweet(toFlagsUserID,
			`#font:kamal
The first tweet is #tweet
#Hello #World.`, "", toFlagsWriter)

		// thread display tag
		testDB.InsertTweet(toFlagsUserID,
			fmt.Sprintf(`#thread:%s:%d
The first tweet is #tweet
#Hello #World.`, th.ID, 1), "", toFlagsWriter)

		// no display tags
		testDB.InsertTweet(toFlagsUserID,
			`The first tweet is #tweet
#Hello #World.`, "", toFlagsWriter)

		// thread display tag
		testDB.InsertTweet(toFlagsUserID,
			fmt.Sprintf(`#thread:%s:%d #font:kamal
The first tweet is #tweet
#Hello #World.`, th.ID, 1), "", toFlagsWriter)

		err := toFlags(testDB, toFlagsWriter)
		assert.Nil(t, err, "failed to convert")

		tws, err := testDB.GetTweets(toFlagsUserID, 0, 10, false, toFlagsWriter)
		assert.Nil(t, err, "Failed to get tweets")
		assert.Equal(t, 4, len(tws), "invalid number of tweets returned")
		for idx, tw := range tws {
			assert.Equal(t, `The first tweet is #tweet
#Hello #World.`, tw.Tweet)
			if idx == 3 {
				assert.Equal(t, "#font:kamal", tw.Flags)
			}

			if idx == 2 {
				assert.Equal(t, fmt.Sprintf("#thread:%s:1", th.ID), tw.Flags)
			}

			if idx == 1 {
				assert.Equal(t, "", tw.Flags)
			}

			if idx == 0 {
				assert.Equal(t, fmt.Sprintf("#thread:%s:1 ", th.ID)+"#font:kamal", tw.Flags)
			}

		}

	})

}
