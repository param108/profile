package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeJWT(t *testing.T) {
	os.Setenv("TRIBIST_JWT_KEY", "6bde337e-1ab0-4dad-8ac3-b9bdefabc88e")
	jwtstr, err := CreateSignedToken("param", "007")
	assert.Nil(t, err, "failed create token")
	claims, err := parseToken(jwtstr)
	assert.Nil(t, err, "failed parsing")

	assert.Equal(t, "param", claims.Username, "invalid username")
	assert.Equal(t, "007", claims.UserID, "invalid userid")
}
