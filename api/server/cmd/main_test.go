package cmd

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/param108/profile/api/store"
)

var testDB store.Store

func TestMain(m *testing.M) {
	err := godotenv.Load("../../test.env")
	if err != nil {
		log.Fatalf("failed to load env: %s", err.Error())
	}
	tDB, err := store.NewStore()
	if err != nil {
		log.Fatalf("failed to create store %s", err.Error())
		return
	}

	testDB = tDB
	m.Run()
}
