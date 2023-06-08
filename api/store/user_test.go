package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testUserWriter = "9435ff03-4600-4413-a1b2-ab4ed205418c"

func teardown() {
	err := testDB.(*StoreImpl).db.Delete("users", testUserWriter)
	if err != nil {
		log.Fatalf("failed delete users: %s", err.Error())
	}
}

func TestUser(t *testing.T) {
	teardown()
	t.Run("can create user", func(t *testing.T) {
		user, err := testDB.CreateUser("param", "twitter", "user", testUserWriter)
		assert.Nil(t, err, "failed create user")

		fetchedUser, err := testDB.GetUser(user.ID, testUserWriter)
		assert.Nil(t, err, "failed fetch user")
		assert.Exactly(t, user, fetchedUser, "created and fetched do not match")
	})

	t.Run("find or create user", func(t *testing.T) {
		user, err := testDB.FindOrCreateUser("param", "twitter", "user", testUserWriter)
		assert.Nil(t, err, "failed to find user")

		fetchedUser, err := testDB.GetUser(user.ID, testUserWriter)
		assert.Nil(t, err, "failed fetch user")
		assert.Exactly(t, user, fetchedUser, "created and fetched do not match")
	})

	t.Run("find by handle", func(t *testing.T) {
		user, err := testDB.GetUserByHandle("param", testUserWriter)
		assert.Nil(t, err, "failed to find user")
		assert.Equal(t, "param", user.Handle, "Invalid user found")
	})

	t.Run("proper error when not found", func(t *testing.T) {
		user, err := testDB.GetUserByHandle("param1", testUserWriter)
		assert.NotNil(t, err, "found user")
		assert.Nil(t, user, "found non nil user")
	})

}
