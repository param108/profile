package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"

	"github.com/urfave/cli/v2"
)

var (
	toFlagsCommand = &cli.Command{
		Name:   "toFlags",
		Usage:  "Move display tags to flags",
		Action: toFlagsCmd,
		Flags:  []cli.Flag{},
	}
)

func toFlags(db store.Store, writer string) error {
	offset := 0
	step := 20

	for {
		tweets, newOffset, err := db.UnsafeGetAllTweets(writer, offset, step)
		if err != nil {
			return err
		}

		for _, tweet := range tweets {
			if len(tweet.Flags) > 0 {
				// ignore tweets whose Flags are already set
				continue
			}

			// parse the first line of the tweet and check if it has
			// a display tag in it.
			fonts, err := utils.ExtractFonts(tweet.Tweet)
			if err != nil {
				fmt.Println("Failed to extract font from tweet:", err.Error())
			}

			threads, err := utils.ExtractThreads(tweet.Tweet)
			if err != nil {
				fmt.Println("Failed to extract threads from tweet:", err.Error())
			}

			if len(fonts) > 0 || len(threads) > 0 {
				lines := strings.Split(tweet.Tweet, "\n")
				if len(lines) > 1 {
					tweet.Tweet = strings.Join(lines[1:], "\n")
				} else {
					tweet.Tweet = ""
				}
				tweet.Flags = lines[0]
			}

			_, _, err = db.UpdateTweet(tweet.UserID,
				tweet.ID, tweet.Tweet, tweet.Flags, tweet.Writer)
			if err != nil {
				fmt.Println("Failed to update tweet", err.Error())
			}
		}

		// If we got less than step tweets then we are done
		if (newOffset - offset) < step {
			break
		}

		offset = newOffset
	}

	return nil
}

func toFlagsCmd(c *cli.Context) error {
	db, err := store.NewStore()
	if err != nil {
		log.Fatalf("failed to connect db:%s", err.Error())
	}

	return toFlags(db, os.Getenv("WRITER"))
}

func init() {
	register(toFlagsCommand)
}
