package tweets

import (
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
		users := strings.Split(r.URL.Query().Get("users"), ",")
		tags := strings.Split(r.URL.Query().Get("tags"), ",")

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
