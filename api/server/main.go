package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"github.com/param108/profile/api/server/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &cli.App{
    Commands: cmd.GetCommands(),
  }

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
