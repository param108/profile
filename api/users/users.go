package users

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/param108/profile/api/store"
)

type ProfileData struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Profile  string `json:"profile"`
}

type ProfileResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    ProfileData `json:"data"`
}

func CreateGetProfileHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("TRIBIST_USERID")
		user, err := db.GetUser(userID, os.Getenv("WRITER"))
		if err != nil {
			log.Printf("Failed profile: %s", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			resp := ProfileResponse{
				Success: false,
				Error:   "failed to get user",
			}
			respData, err := json.Marshal(resp)
			if err != nil {
				log.Printf("Failed marshal error resp: %s", err.Error())
				return
			}
			rw.Write(respData)
			return
		}

		resp := ProfileResponse{
			Success: true,
			Data: ProfileData{
				Username: user.Handle,
				UserID:   user.ID,
				Profile:  user.Profile,
			},
		}

		respData, err := json.Marshal(resp)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			resp := ProfileResponse{
				Success: false,
				Error:   "failed to marshal response",
			}
			respData, err := json.Marshal(resp)
			if err != nil {
				log.Printf("Failed marshal error resp: %s", err.Error())
				return
			}
			rw.Write(respData)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respData)
	}
}
