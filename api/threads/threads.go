package threads

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

func CreateMakeThreadHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("TRIBIST_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := models.CreateThreadRequest{}
		if err := json.Unmarshal(data, &req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt parse:"+err.Error())
			return
		}

		thread, err := db.CreateThread(userID, req.Name, os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "failed to create:"+err.Error())
			return
		}

		utils.WriteData(rw, http.StatusOK, thread)
	}
}

func CreateGetThreadHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		threadID := strings.TrimSpace(v["thread_id"])

		if len(threadID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid thread")
			return
		}

		userStr := strings.TrimSpace(v["username"])
		if len(userStr) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		user, err := db.GetUserByHandle(userStr, os.Getenv("WRITER"))
		if err != nil {
			log.Println("failed get userbyhandle:", err.Error())
			utils.WriteError(rw, http.StatusBadRequest, "invalid user")
			return
		}

		threadData, err := db.GetThread(user.ID, threadID, os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "failed to get:"+err.Error())
			return
		}

		utils.WriteData(rw, http.StatusOK, threadData)
	}
}

func CreateDeleteThreadHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("TRIBIST_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := models.DeleteThreadRequest{}
		if err := json.Unmarshal(data, &req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt parse:"+err.Error())
			return
		}

		thread, err := db.DeleteThread(userID, req.ThreadID, os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "failed to delete:"+err.Error())
			return
		}

		utils.WriteData(rw, http.StatusOK, thread)
	}
}
