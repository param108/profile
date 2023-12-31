package spmsg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

const (
	DEFAULT_MESSAGE_LIMIT = 50
)

func CreateGetUserMessages(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		timeStr := r.URL.Query().Get("time")
		limitStr := r.URL.Query().Get("limit")
		tzStr := r.URL.Query().Get("tz")

		var start time.Time

		if len(timeStr) > 0 {
			st, err := time.Parse(time.RFC3339, timeStr)
			if err != nil {
				utils.WriteError(rw, http.StatusBadRequest,
					fmt.Sprintf("Invalid time: %s", err.Error()))
				return
			}

			start = st
		} else {
			start = time.Now()
		}

		var limit int
		if len(limitStr) > 0 {
			l, err := strconv.Atoi(limitStr)
			if err != nil {
				utils.WriteError(rw, http.StatusBadRequest,
					fmt.Sprintf("Invalid limit: %s", err))
				return
			}

			limit = l
		} else {
			limit = DEFAULT_MESSAGE_LIMIT
		}

		var tz = "Asia/Kolkata"

		if len(tzStr) > 0 {
			_, err := time.LoadLocation(tzStr)
			if err != nil {
				utils.WriteError(rw, http.StatusBadRequest,
					fmt.Sprintf("Invalid location: %s %s", tzStr, err))
				return
			}
			tz = tzStr
		}

		writer := os.Getenv("WRITER")

		data, err := db.GetSPUserMessagesByDay(userID, start, tz, limit, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError,
				fmt.Sprintf("didnt get data %s", err))
			return
		}

		utils.WriteData(rw, http.StatusOK, data)
	}
}

func CreatePostUserMessages(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := models.SpMessage{}
		if err := json.Unmarshal(data, &req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt parse:"+err.Error())
			return
		}

		if req.SpUserID != userID {
			utils.WriteError(rw, http.StatusBadRequest, "wrong user")
			return
		}

		req.Writer = os.Getenv("WRITER")

		req.CreatedAt = time.Now().UTC()
		// FIXME validate

		msg, err := db.AddSpMessage(&req, req.Writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt save msg:"+err.Error())
			return
		}

		utils.WriteData(rw, http.StatusOK, msg)
	}
}
