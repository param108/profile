package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidToken(t *testing.T) {
	testWriter := "a141224f-0afe-4335-9381-cfd8c10a1c5f"

	func() {
		err := testDB.(*StoreImpl).db.Delete("invalid_tokens", testWriter)
		if err != nil {
			log.Fatalf("failed delete invalid_tokens: %s", err.Error())
		}
	}()

	t.Run("Should be able to create token", func(t *testing.T) {
		err := testDB.CreateInvalidToken("dummy token", testWriter)
		assert.Nil(t, err, "couldnt create token")
	})

	t.Run("Should be able to create second token", func(t *testing.T) {
		err := testDB.CreateInvalidToken("dummy token 2", testWriter)
		assert.Nil(t, err, "couldnt create token")
	})

	t.Run("Cant create multiple invalid tokens", func(t *testing.T) {
		err := testDB.CreateInvalidToken("dummy token", testWriter)
		assert.NotNil(t, err, "couldnt create token")
	})

	t.Run("Should mark invalid token as invalid", func(t *testing.T) {
		found, err := testDB.IsInvalidToken("dummy token", testWriter)
		assert.Nil(t, err, "failed to check validity")
		assert.True(t, found, "failed find invalid token")
	})

	t.Run("Should mark invalid token as invalid", func(t *testing.T) {
		found, err := testDB.IsInvalidToken("dummy token 2", testWriter)
		assert.Nil(t, err, "failed to check validity")
		assert.True(t, found, "failed find invalid token")
	})

	t.Run("Should return false if invalid token not found", func(t *testing.T) {
		found, err := testDB.IsInvalidToken("dummy token 1", testWriter)
		assert.Nil(t, err, "failed to check validity")
		assert.False(t, found, "failed find invalid token")
	})

}
