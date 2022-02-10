package store

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	testDB.CreateUser("param")
}
