package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/param108/profile/api/store"
)

func CreateGetOneTimeHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{}

		id := r.URL.Query().Get("id")
		if len(id) == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("bad code"))
			return
		}

		onetime, err := db.GetOneTime(id, time.Second*60, os.Getenv("WRITER"))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("no one time"))
			return
		}

		resp["success"] = true
		resp["data"] = map[string]interface{}{
			"value": onetime.Data,
		}

		respStr, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Failed marshal %s\n", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("failed marshall"))
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respStr)
	}
}
