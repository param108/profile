package twitter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidRedirectURL(t *testing.T) {

	t.Run("only urls beginning with'/' are allowed", func(t *testing.T) {
		os.Unsetenv("ALLOWED_CLIENTS")
		assert.False(t, isValidRedirectURL("https://anywhere.com/a/b/c"), "invalid return value")
		assert.True(t, isValidRedirectURL("/a/b/c"), "invalid return value")

	})

}
