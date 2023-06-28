package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractThread(t *testing.T) {
	t.Run("Correctly extracts the thread_id and seq of thread",
		func(t *testing.T) {
			tweet := `#display #thread:c6340c54-da8e-4c8f-b1e4-4840be4d4dfc:0
This is the start of a tweet`
			details, err := ExtractThreads(tweet)
			assert.Nil(t, err, "failed extract")
			assert.Equal(t, 1, len(details))
			assert.Equal(t, "c6340c54-da8e-4c8f-b1e4-4840be4d4dfc", details[0].ID, "invalid id")
			assert.Equal(t, 0, details[0].Seq)
		})

	t.Run("Correctly extracts the thread_id and seq of thread not head",
		func(t *testing.T) {
			tweet := `#display #thread:c6340c54-da8e-4c8f-b1e4-4840be4d4dfc:100
This is the start of a tweet`
			details, err := ExtractThreads(tweet)
			assert.Nil(t, err, "failed extract")
			assert.Equal(t, 1, len(details))
			assert.Equal(t, "c6340c54-da8e-4c8f-b1e4-4840be4d4dfc", details[0].ID, "invalid id")
			assert.Equal(t, 100, details[0].Seq)
		})

	t.Run("Correctly extracts multiple threads",
		func(t *testing.T) {
			tweet := `#display #thread:c6340c54-da8e-4c8f-b1e4-4840be4d4dfc:100 #thread:f11e386f-ae1a-4927-9dfd-146f12498b0c:5
This is the start of a tweet`
			details, err := ExtractThreads(tweet)
			assert.Nil(t, err, "failed extract")
			assert.Equal(t, 2, len(details))
			assert.Equal(t, "c6340c54-da8e-4c8f-b1e4-4840be4d4dfc", details[0].ID, "invalid id")
			assert.Equal(t, 100, details[0].Seq)
			assert.Equal(t, "f11e386f-ae1a-4927-9dfd-146f12498b0c", details[1].ID, "invalid id")
			assert.Equal(t, 5, details[1].Seq)

		})

	t.Run("Ignore invalid thread_id",
		func(t *testing.T) {
			tweet := `#display #thread:junk:100 #thread:f11e386f-ae1a-4927-9dfd-146f12498b0c:5
This is the start of a tweet`
			details, err := ExtractThreads(tweet)
			assert.Nil(t, err, "failed extract")
			assert.Equal(t, 1, len(details))
			assert.Equal(t, "f11e386f-ae1a-4927-9dfd-146f12498b0c", details[0].ID, "invalid id")
			assert.Equal(t, 5, details[0].Seq)

		})

	t.Run("Ignore invalid seq",
		func(t *testing.T) {
			tweet := `#display #thread:c6340c54-da8e-4c8f-b1e4-4840be4d4dfc:abc #thread:f11e386f-ae1a-4927-9dfd-146f12498b0c:5
This is the start of a tweet`
			details, err := ExtractThreads(tweet)
			assert.Nil(t, err, "failed extract")
			assert.Equal(t, 1, len(details))
			assert.Equal(t, "f11e386f-ae1a-4927-9dfd-146f12498b0c", details[0].ID, "invalid id")
			assert.Equal(t, 5, details[0].Seq)

		})

	t.Run("No thread should return empty return without errors",
		func(t *testing.T) {
			tweet := `#display
This is the start of a tweet`
			details, err := ExtractThreads(tweet)
			assert.Nil(t, err, "failed extract")
			assert.Equal(t, 0, len(details))

		})

}
