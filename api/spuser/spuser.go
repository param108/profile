package spuser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

func CreateGetSPUserHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		spUser, err := db.GetSPUserByID(userID, os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "did not find")
			return
		}

		utils.WriteData(rw, http.StatusOK, spUser)
	}
}

func CreateUpdateSPUserHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := models.SpUser{}
		if err := json.Unmarshal(data, &req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt parse:"+err.Error())
			return
		}

		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		if userID != req.ID {
			utils.WriteError(rw, http.StatusForbidden, "wrong user")
			return
		}

		writer := os.Getenv("WRITER")

		if writer != req.Writer {
			utils.WriteError(rw, http.StatusForbidden, "wrong writer")
			return
		}

		spUser, err := db.UpdateSPUser(&req)
		if err != nil {
			log.Printf("failed to update user" + err.Error())
			utils.WriteError(rw, http.StatusBadRequest, "Error saving")
			return
		}

		utils.WriteData(rw, http.StatusOK, spUser)
	}
}

func CreateRefreshSPUserTokenHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("SP_USERID")
		phone := r.Header.Get("SP_PHONE")

		if len(userID) == 0 || len(phone) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		spUser, err := db.FindOrCreateSPUser(phone, os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "Failed finding sp user:"+err.Error())
			return
		}

		accessToken, refreshToken, err := utils.CreateSignedSPTokens(spUser.Phone, spUser.ID)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "Failed creating tokens:"+err.Error())
			return
		}

		resp := models.RefreshTokenResponse{
			SpUser:       spUser,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		utils.WriteData(rw, http.StatusOK, resp)
	}
}
