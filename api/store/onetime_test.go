package store

import (
	"log"
	"testing"
	"time"

	"github.com/param108/profile/api/models"
	"github.com/stretchr/testify/assert"
)

const onetimeWriter = "3e3c614b-58ea-4c17-a458-1b7236bbf9dd"

func onetimeTeardown() {
	err := testDB.(*StoreImpl).db.Delete("onetime", onetimeWriter)
	if err != nil {
		log.Fatalf("failed delete onetimes: %s", err.Error())
	}
}

func TestOneTime(t *testing.T) {
	onetimeTeardown()
	var oldOnetime *models.Onetime
	t.Run("set one time value", func(t *testing.T) {
		onetime, err := testDB.SetOneTime("beauty is a key", onetimeWriter)
		assert.Nil(t, err, "Failed to set one time")

		assert.NotEmpty(t, onetime.ID, "id not set")
		assert.NotEmpty(t, onetime.CreatedAt, "created_at not set")
		assert.Equal(t, "beauty is a key", onetime.Data)
		oldOnetime = onetime
	})

	oldOnetime.CreatedAt = oldOnetime.CreatedAt.Add(-time.Second * 20)
	testDB.(*StoreImpl).db.SaveOneTime(oldOnetime)

	t.Run("get unexpired record", func(t *testing.T) {
		onetime, err := testDB.GetOneTime(oldOnetime.ID, time.Second*50, onetimeWriter)
		assert.Nil(t, err, "failed to get one time")
		assert.Equal(t, "beauty is a key", onetime.Data, "invalid data found")
	})

	t.Run("get expired record", func(t *testing.T) {
		onetime, err := testDB.GetOneTime(oldOnetime.ID, time.Second*5, onetimeWriter)
		assert.Equal(t, "not found", err.Error(), "invalid error seen")
		assert.Nil(t, onetime, "onetime not nil")
	})

	t.Run("delete old onetimes", func(t *testing.T) {
		err := testDB.DeleteOldOneTimes(time.Second*5, onetimeWriter)
		assert.Nil(t, err, "failed to delete old onetime")

		onetime, err := testDB.GetOneTime(oldOnetime.ID, time.Second*50, onetimeWriter)
		assert.Equal(t, "not found", err.Error(), "invalid error seen")
		assert.Nil(t, onetime, "onetime not nil")
	})
}
