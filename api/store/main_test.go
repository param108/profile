package store

import (
	"github.com/joho/godotenv"
)

var testDB *StoreImpl

func TestMain() {
	godotenv.Load("test.env")
	testDB, _ = NewStore()
}
