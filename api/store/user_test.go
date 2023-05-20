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
}
