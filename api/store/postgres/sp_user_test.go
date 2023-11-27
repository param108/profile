package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneRegexp(t *testing.T) {
	invalidPhones := []string{
		"", "9809", "90897a6578", "9808 980987",
		`9880154321
abcdefg`}

	for _, invalidPhone := range invalidPhones {
		assert.False(t, validPhoneRe.MatchString(invalidPhone), "invalid phone accepted")
	}

	validPhones := []string{
		"0909090909",
		"0000000000",
		"9898678767",
	}

	for _, validPhone := range validPhones {
		assert.True(t, validPhoneRe.MatchString(validPhone), "valid phone rejected")
	}

}
