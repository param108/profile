package spmsg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
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

func CreateGetSPGroupMessagesHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")

		v := mux.Vars(r)

		groupID := strings.TrimSpace(v["group_id"])

		if len(groupID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid group")
			return
		}

		// check if user is part of this group
		groupUser, err := db.GetSPGroupUser(userID, groupID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		// user is valid so get the CreateGetSPGroupMessagesHandler
		if groupUser.Deleted {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
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

		data, err := db.GetSPGroupMessagesByDay(groupID, start, tz, limit, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError,
				fmt.Sprintf("didnt get data %s", err))
			return
		}

		utils.WriteData(rw, http.StatusOK, data)
	}
}

func CreateGetSPGroupUserMessagesHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")

		v := mux.Vars(r)

		groupID := strings.TrimSpace(v["group_id"])

		if len(groupID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid group")
			return
		}

		reqUserID := strings.TrimSpace(v["user_id"])

		if len(reqUserID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user id")
			return
		}

		// check if user is part of this group
		groupUser, err := db.GetSPGroupUser(userID, groupID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		// check if reqUser is part of the group
		_, err = db.GetSPGroupUser(reqUserID, groupID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "invalid request user")
			return
		}

		if groupUser.Deleted {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		if groupUser.Role != "admin" {
			utils.WriteError(rw, http.StatusForbidden, "must be admin")
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

		data, err := db.GetSPUserMessagesByDay(reqUserID, start, tz, limit, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError,
				fmt.Sprintf("didnt get data %s", err))
			return
		}

		utils.WriteData(rw, http.StatusOK, data)
	}
}

func CreateGetSPGroupUserOTPHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")

		v := mux.Vars(r)

		groupID := strings.TrimSpace(v["group_id"])

		if len(groupID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid group")
			return
		}

		reqUserID := strings.TrimSpace(v["user_id"])

		if len(reqUserID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user id")
			return
		}

		// check if user is part of this group
		groupUser, err := db.GetSPGroupUser(userID, groupID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		// check if reqUser is part of the group
		_, err = db.GetSPGroupUser(reqUserID, groupID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "invalid request user")
			return
		}

		if groupUser.Deleted {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		if groupUser.Role != "admin" {
			utils.WriteError(rw, http.StatusForbidden, "must be admin")
			return
		}

		reqUser, err := db.GetSPUserByID(reqUserID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		err = db.CreateOTP(reqUser.Phone, time.Now(), writer)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError,
				fmt.Sprintf("didnt get data %s", err))
			return
		}

		otp, err := db.GetOTP(reqUser.Phone, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError,
				fmt.Sprintf("didnt get data %s", err))
			return
		}

		utils.WriteData(rw, http.StatusOK, otp.Code)
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

		var tz = "Asia/Kolkata"

		req.Writer = os.Getenv("WRITER")

		req.CreatedAt = time.Now().UTC()
		// FIXME validate

		msg, err := db.AddSpMessage(&req, tz, req.Writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt save msg:"+err.Error())
			return
		}

		ret := []*models.SpGroupMsgData{msg}
		utils.WriteData(rw, http.StatusOK, ret)
	}
}
