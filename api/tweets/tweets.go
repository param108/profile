package tweets

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

const MAX_TWEETS_PER_QUERY = 20

func CreateGetTweetsHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		usersStr := strings.TrimSpace(r.URL.Query().Get("users"))
		tagsStr := strings.TrimSpace(r.URL.Query().Get("tags"))

		if len(usersStr) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "need exactly one user")
			return
		}

		users := strings.Split(usersStr, ",")

		// Split of empty string returns a slice of one element.
		// The element is empty string and will not match any tags.
		// Check if the input is empty before trying to split it.
		var tags []string
		if len(tagsStr) > 0 {
			tags = strings.Split(tagsStr, ",")
		}

		if len(users) != 1 {
			utils.WriteError(rw, http.StatusBadRequest, "need exactly one user")
			return
		}

		user, err := db.GetUserByHandle(users[0], os.Getenv("WRITER"))
		if err != nil {
			status := http.StatusInternalServerError
			if err.Error() == "not found" {
				status = http.StatusNotFound
			}
			utils.WriteError(rw, status, err.Error())
			return
		}

		offset := 0

		offsetStr := r.URL.Query().Get("offset")
		if len(offsetStr) != 0 {
			o, err := strconv.Atoi(offsetStr)
			if err != nil {
				utils.WriteError(rw, http.StatusBadRequest, "invalid offset")
				return
			}
			offset = o
		}

		var tweets []*models.Tweet
		if len(tags) == 0 {
			t, err := db.GetTweets(user.ID, offset, MAX_TWEETS_PER_QUERY, os.Getenv("WRITER"))
			if err != nil {
				utils.WriteError(rw, http.StatusInternalServerError, err.Error())
				return
			}
			tweets = t
		} else {
			t, err := db.SearchTweetsByTags(user.ID, tags,
				MAX_TWEETS_PER_QUERY, os.Getenv("WRITER"))
			if err != nil {
				utils.WriteError(rw, http.StatusInternalServerError, err.Error())
				return
			}
			tweets = t
		}

		utils.WriteData(rw, http.StatusOK, tweets)
	}
}

func CreatePostTweetsHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := models.PostTweetsRequest{}
		if err := json.Unmarshal(data, &req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt parse:"+err.Error())
			return
		}

		userID := r.Header.Get("TRIBIST_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		tweet, _, err := db.InsertTweet(userID, req.Tweet, "", os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, err.Error())
			return
		}
		utils.WriteData(rw, http.StatusOK, tweet)
	}
}
