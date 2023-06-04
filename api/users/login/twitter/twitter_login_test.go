package twitter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidRedirectURL(t *testing.T) {

	t.Run("if env is unset it should be false", func(t *testing.T) {
		os.Unsetenv("ALLOWED_CLIENTS")
		assert.False(t, isValidRedirectURL("https://anywhere.com/a/b/c"), "invalid return value")
	})

	t.Run("if env is unset it should be false", func(t *testing.T) {
		os.Setenv("ALLOWED_CLIENTS", "ui.tribist.com,de.tribist.com")
		assert.False(t, isValidRedirectURL("https://anywhere.com/a/b/c"), "invalid return value")
	})

	t.Run("if env set it should allow all apt https redirects", func(t *testing.T) {
		assert.True(t, isValidRedirectURL("https://de.tribist.com/a/b/c"),
			"invalid return value")
		assert.True(t, isValidRedirectURL("https://ui.tribist.com/a/b/c"),
			"invalid return value")
		assert.False(t, isValidRedirectURL("http://ui.tribist.com/a/b/c"),
			"http redirect not allowed")
		assert.False(t, isValidRedirectURL("http://de.tribist.com/a/b/c"),
			"http redirect not allowed")
		assert.False(t, isValidRedirectURL("http://fg.tribist.com/a/b/c"),
			"only identical urls should succeed")
	})

	t.Run("empty allowed clients should be skipped", func(t *testing.T) {
		os.Setenv("ALLOWED_CLIENTS", "")
		assert.False(t,
			isValidRedirectURL(
				"https://ui.tribist.com/a/b/c"), "empty allowed_clients succeeded")
		os.Setenv("ALLOWED_CLIENTS", ",")
		assert.False(t,
			isValidRedirectURL(
				"https://ui.tribist.com/a/b/c"), "empty allowed_clients succeeded")
		os.Setenv("ALLOWED_CLIENTS", ",ui.tribist.com")
		assert.True(t,
			isValidRedirectURL(
				"https://ui.tribist.com/a/b/c"), "empty allowed_clients succeeded")
		os.Setenv("ALLOWED_CLIENTS", "ui.tribist.com,")
		assert.True(t,
			isValidRedirectURL(
				"https://ui.tribist.com/a/b/c"), "empty allowed_clients succeeded")
	})
}
