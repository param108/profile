package spotps

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

func CreateMakeOTPHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := &models.CreateOTPRequest{}

		if err := json.Unmarshal(data, req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		// empty API KEY in env is an error
		if os.Getenv("OTP_API_KEY") == "" || req.APIKey != os.Getenv("OTP_API_KEY") {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		if err := db.CreateOTP(req.Phone, time.Now(), os.Getenv("WRITER")); err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "Internal Error:"+err.Error())
			return
		}

		utils.WriteData(rw, http.StatusOK, "ok")
	}
}

func CreateCheckOTPHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := &models.CheckOTPRequest{}

		if err := json.Unmarshal(data, req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		// empty API KEY in env is an error
		if os.Getenv("OTP_API_KEY") == "" || req.APIKey != os.Getenv("OTP_API_KEY") {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		spOtp, err := db.CheckOTP(req.Phone, req.Code, time.Now(), os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		// TODO delete the OTP now.

		// successful validation of otp
		// 1. Create the spuser if not exists
		// 2. Create the access and refresh tokens for that user
		spUser, err := db.FindOrCreateSPUser(spOtp.Phone, os.Getenv("WRITER"))
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "Failed finding sp user:"+err.Error())
			return
		}

		accessToken, refreshToken, err := utils.CreateSignedSPTokens(spUser.Phone, spUser.ID)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "Failed creating tokens:"+err.Error())
			return
		}
		resp := models.CheckOTPResponse{
			SpUser:       spUser,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		utils.WriteData(rw, http.StatusOK, resp)
	}
}
