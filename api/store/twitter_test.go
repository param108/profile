package store

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	testWriter := "951295a6-dc7e-4855-a505-352ee351d427"

	func() {
		err := testDB.(*StoreImpl).db.Delete("twitter_challenges", testWriter)
		if err != nil {
			log.Fatalf("failed delete twitter_challenges: %s", err.Error())
		}
	}()

	var twitterKey1 string
	t.Run("Should be able to create challenge", func(t *testing.T) {
		key, err := testDB.CreateTwitterChallenge("JustForTesting1", "url1", testWriter)
		assert.Nil(t, err, "couldnt create twitter challenge")
		assert.Greater(t, len(key), 0, "bad key returned")
		twitterKey1 = key
	})

	var twitterKey2 string
	t.Run("Should be able to create second challenge", func(t *testing.T) {
		key, err := testDB.CreateTwitterChallenge("JustForTesting2", "url2", testWriter)
		assert.Nil(t, err, "couldnt create twitter challenge")
		assert.Greater(t, len(key), 0, "bad key returned")
		twitterKey2 = key
	})

	t.Run("Should get challenge for key 1", func(t *testing.T) {
		token, urlStr, err := testDB.GetTwitterChallenge(twitterKey1, testWriter)
		assert.Nil(t, err, "couldnt get twitter challenge")
		assert.Equal(t, "JustForTesting1", token, "bad key returned")
		assert.Equal(t, "url1", urlStr, "invalid redirect url")
	})

	t.Run("Should get challenge for key 2", func(t *testing.T) {
		token, urlStr, err := testDB.GetTwitterChallenge(twitterKey2, testWriter)
		assert.Nil(t, err, "couldnt get twitter challenge")
		assert.Equal(t, "JustForTesting2", token, "bad key returned")
		assert.Equal(t, "url2", urlStr, "invalid redirect url")
	})

	time.Sleep(time.Second * 5)

	t.Run("Should delete older than 1 second", func(t *testing.T) {
		testDB.DeleteOldTwitterChallenges(time.Second)

		_, _, err := testDB.GetTwitterChallenge(twitterKey1, testWriter)
		assert.NotNil(t, err, "twitter challenge not deleted")
		_, _, err = testDB.GetTwitterChallenge(twitterKey2, testWriter)
		assert.NotNil(t, err, "twitter challenge not deleted")
	})

}
