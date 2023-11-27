package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func spUserTeardown(writer string) {
	err := testDB.(*StoreImpl).db.Delete("sp_users", writer)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}
}

func TestFindOrCreateSPUser(t *testing.T) {
	writer := "4e151bce-7b29-42dc-99fd-475f16422e74"

	spUserTeardown(writer)
	phone1 := "9876556789"
	phone2 := "7687678787"
	user1ID := ""
	user2ID := ""
	t.Run("can create spuser from scratch", func(t *testing.T) {
		user1, err := testDB.FindOrCreateSPUser(phone1, writer)
		assert.Nil(t, err, "Failed to find or create user")
		assert.Equal(t, phone1, user1.Phone, "wrong phone")
		assert.Equal(t, writer, user1.Writer, "wrong phone")
		assert.NotEmpty(t, user1.ID, "ID not populated")
		user1ID = user1.ID
	})

	t.Run("can create second spuser from scratch", func(t *testing.T) {
		user2, err := testDB.FindOrCreateSPUser(phone2, writer)
		assert.Nil(t, err, "Failed to find or create user")
		assert.Equal(t, phone2, user2.Phone, "wrong phone")
		assert.Equal(t, writer, user2.Writer, "wrong phone")
		assert.NotEmpty(t, user2.ID, "ID not populated")
		user2ID = user2.ID
	})

	t.Run("can find first spuser", func(t *testing.T) {
		user1, err := testDB.FindOrCreateSPUser(phone1, writer)
		assert.Nil(t, err, "Failed to find or create user")
		assert.Equal(t, phone1, user1.Phone, "wrong phone")
		assert.Equal(t, writer, user1.Writer, "wrong phone")
		assert.Equal(t, user1ID, user1.ID, "ID incorrect")
	})

	t.Run("can find second spuser", func(t *testing.T) {
		user2, err := testDB.FindOrCreateSPUser(phone2, writer)
		assert.Nil(t, err, "Failed to find or create user")
		assert.Equal(t, phone2, user2.Phone, "wrong phone")
		assert.Equal(t, writer, user2.Writer, "wrong phone")
		assert.Equal(t, user2ID, user2.ID, "ID incorrect")
	})

}
