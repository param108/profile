package tweets

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

type GetPutImageUrlResponse struct {
	URL string `json:"url"`
	Headers map[string]string `json:"headers"`
}

func CreatePutImageSignedUrlHandler(db store.Store, aws *utils.AWS) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("TRIBIST_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		suffix := strings.TrimSpace(r.URL.Query().Get("suffix"))
		// add .<suffix> to the key if a suffix is provided.
		if len(suffix) > 0 {
			suffix = "." + suffix
		}

		bucket := os.Getenv("AWS_IMAGE_BUCKET")
		u,err := uuid.NewUUID()
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "failed to create uuid")
			return
		}
		url, headers, err := aws.CreateSignedPutUrl(
			bucket,
			userID+"_"+u.String()+suffix,
			time.Second*600)
		if err != nil {
			fmt.Println("failed to create url", err.Error())
			utils.WriteError(rw, http.StatusInternalServerError, "failed to create url")
			return
		}

		ret := GetPutImageUrlResponse{
			URL: url,
			Headers: headers,
		}

		utils.WriteData(rw, http.StatusOK, ret)
	}
}
