package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const resourceUserID = "406204db-2614-483d-9985-760b7a6571bd"

func resourceTearDown(writer string) {
	err := testDB.(*StoreImpl).db.Delete("resources", writer)
	if err != nil {
		log.Fatalf("failed delete tweets: %s", err.Error())
	}
}

func TestIncrDecrResources(t *testing.T) {
	const resourceWriter = "df125b3d-25d3-4c06-9a6f-5220f52451c2"

	// perform cleanup before teardown.
	resourceTearDown(resourceWriter)

	t.Run("Incr new resource", func(t *testing.T) {
		res, err := testDB.LockResource(resourceUserID, "image", resourceWriter)
		assert.Nil(t, err, "failed to lock resource")
		assert.Equal(t, 1, res.Value, "wrong resource count")
		assert.Equal(t, 10000000, res.Max, "wrong resource max")
	})

	t.Run("Incr existing resource", func(t *testing.T) {
		res, err := testDB.LockResource(resourceUserID, "image", resourceWriter)
		assert.Nil(t, err, "failed to lock resource")
		assert.Equal(t, 2, res.Value, "wrong resource count")
		assert.Equal(t, 10000000, res.Max, "wrong resource max")
	})

	t.Run("Incr new resource", func(t *testing.T) {
		res, err := testDB.LockResource(resourceUserID, "image", resourceWriter)
		assert.Nil(t, err, "failed to lock resource")
		assert.Equal(t, 3, res.Value, "wrong resource count")
		assert.Equal(t, 10000000, res.Max, "wrong resource max")
	})

	t.Run("Try and incr more than maximum", func(t *testing.T) {
		var err error
		testDB.SetResources(resourceUserID, "image", 9999995, resourceWriter)
		for i := 0; i < 10; i++ {
			_, err = testDB.LockResource(resourceUserID, "image", resourceWriter)
			if err != nil {
				break
			}
		}

		assert.NotNil(t, err, "successfully incremented above maximum")

		resources, err := testDB.GetResources(resourceUserID, resourceWriter)
		assert.Nil(t, err, "failed to get resources")
		for _, res := range resources {
			if res.T == "image" {
				assert.Equal(t, 10000000, res.Value)
			}
		}
	})

	t.Run("decr new resource", func(t *testing.T) {
		res, err := testDB.UnlockResource(resourceUserID, "tape", resourceWriter)
		assert.Nil(t, err, "failed to lock resource")
		assert.Equal(t, 0, res.Value, "wrong resource count")
		assert.Equal(t, 10, res.Max, "wrong resource max")
	})

	t.Run("decr existing resource", func(t *testing.T) {
		res, err := testDB.UnlockResource(resourceUserID, "image", resourceWriter)
		assert.Nil(t, err, "failed to lock resource")
		assert.Equal(t, 9999999, res.Value, "wrong resource count")
		assert.Equal(t, 10000000, res.Max, "wrong resource max")
	})

	t.Run("decr existing resource", func(t *testing.T) {
		res, err := testDB.UnlockResource(resourceUserID, "image", resourceWriter)
		assert.Nil(t, err, "failed to lock resource")
		assert.Equal(t, 9999998, res.Value, "wrong resource count")
		assert.Equal(t, 10000000, res.Max, "wrong resource max")
	})

	t.Run("Try and decr less than 0", func(t *testing.T) {
		var err error
		testDB.SetResources(resourceUserID, "image", 5, resourceWriter)

		for i := 0; i < 10; i++ {
			_, err = testDB.UnlockResource(resourceUserID, "image", resourceWriter)
			if err != nil {
				break
			}
		}
		assert.NotNil(t, err, "successfully decremented below 0")
		resources, err := testDB.GetResources(resourceUserID, resourceWriter)
		assert.Nil(t, err, "failed to get resources")
		for _, res := range resources {
			if res.T == "image" {
				assert.Equal(t, 0, res.Value)
			}
		}
	})

	t.Run("Set the value correctly", func(t *testing.T) {
		res, err := testDB.SetResources(
			resourceUserID, "image", 3, resourceWriter)
		assert.Nil(t, err, "failed")
		assert.Equal(t, 3, res.Value, "wrong value")
	})

	t.Run("Cant update greater than limit", func(t *testing.T) {
		_, err := testDB.SetResources(
			resourceUserID, "image", 10000005, resourceWriter)
		assert.NotNil(t, err, "incorrectly succeeded")
	})
}
