package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFontFound(t *testing.T) {

	t.Run("Can find font when it is there", func(t *testing.T) {
		fonts, err := ExtractFonts(`#font:old-font
Hello this is a tweet.`)
		assert.Nil(t, err, "Error is nil")
		assert.Equal(t, 1, len(fonts))
		assert.Equal(t, "old-font", fonts[0].Name)
	})

	t.Run("Cant find font when there isnt one", func(t *testing.T) {
		fonts, err := ExtractFonts(`Hello this is a tweet.`)
		assert.Nil(t, err, "Error is nil")
		assert.Equal(t, 0, len(fonts))
	})

}
