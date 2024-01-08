package spservices

import (
	"net/http"
	"os"

	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

func CreateGetServices(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "unknown user")
			return
		}

		writer := os.Getenv("WRITER")
		data, err := db.GetSPServices(writer)
		if err != nil {
			utils.WriteError(rw, http.StatusInternalServerError, "find services")
			return
		}

		utils.WriteData(rw, http.StatusOK, data)
	}
}
