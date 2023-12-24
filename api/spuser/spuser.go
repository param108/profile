package spuser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
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

		user, err := db.GetSPUserByID(userID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "bad request: "+err.Error())
			return
		}

		if req.Name == "" || req.PhotoURL == "" {
			utils.WriteError(rw, http.StatusBadRequest, "name and photoURL cannot be empty")
			return
		}

		user.Name = req.Name
		user.PhotoURL = req.PhotoURL
		user.ProfileUpdated = true

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

type GetPutImageUrlResponse struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type GetPutImageUrlRequest struct {
	APIToken string `json:"api_token"`
	Suffix   string `json:"suffix"`
}

func CanAllocateResources(userID string, db store.Store, aws *utils.AWS) (bool, error) {
	totalSize, _, err := aws.GetSPBucketSize(os.Getenv("AWS_IMAGE_BUCKET"), userID)
	if err != nil {
		return false, err
	}

	res, err := db.SetResources(userID, "images", int(totalSize), os.Getenv("WRITER"))
	if err != nil {
		return false, err
	}

	if res.Value >= res.Max {
		return false, nil
	}

	return true, nil
}

func CreatePutImageSignedUrlHandler(db store.Store, aws *utils.AWS) http.HandlerFunc {
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

		req := GetPutImageUrlRequest{}
		if err := json.Unmarshal(data, &req); err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt parse:"+err.Error())
			return
		}

		if req.APIToken != os.Getenv("IMAGE_UPLOAD_API_KEY") {
			utils.WriteError(rw, http.StatusForbidden, "invalid API Key")
			return
		}

		if allowed, err := CanAllocateResources(userID, db, aws); (err != nil) || !allowed {
			if err != nil {
				utils.WriteError(rw, http.StatusInternalServerError, "couldnt check resources")
				return
			}

			if !allowed {
				utils.WriteError(rw, http.StatusTooManyRequests, "too many resources")
				return
			}
		}

		suffix := strings.TrimSpace(req.Suffix)
		// add .<suffix> to the key if a suffix is provided.
		if len(suffix) > 0 {
			suffix = "." + suffix
		}

		bucket := os.Getenv("AWS_IMAGE_BUCKET")
		u, err := uuid.NewUUID()
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "failed to create uuid")
			return
		}
		url, headers, err := aws.CreateSignedPutUrl(
			bucket,
			"sp_data_"+userID+"_"+u.String()+suffix,
			time.Second*600)
		if err != nil {
			fmt.Println("failed to create url", err.Error())
			utils.WriteError(rw, http.StatusInternalServerError, "failed to create url")
			return
		}

		ret := GetPutImageUrlResponse{
			URL:     url,
			Headers: headers,
		}

		utils.WriteData(rw, http.StatusOK, ret)
	}
}
