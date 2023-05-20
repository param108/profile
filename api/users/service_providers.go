package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/users/login/twitter"
)

func CreateServiceProviderLoginRedirect(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		serviceProvider := r.URL.Query().Get("source")

		resp := models.Response{}

		switch serviceProvider {
		case "twitter":
			tlp := twitter.NewTwitterLoginProvider(db)
			tlp.HandleLogin(rw, r)
			return
		default:
			rw.WriteHeader(http.StatusBadRequest)
			resp.Success = false
			resp.Errors = []string{"invalid source: " + serviceProvider}
			b, _ := json.Marshal(resp)
			rw.Write(b)
			return
		}

	}
}

func CreateServiceProviderAuthorizeRedirect(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		resp := models.Response{}

		v := mux.Vars(r)

		serviceProvider := v["service_provider"]

		switch serviceProvider {
		case "twitter":
			tlp := twitter.NewTwitterLoginProvider(db)
			tlp.HandleAuthorize(rw, r)
			return
		default:
			rw.WriteHeader(http.StatusBadRequest)
			resp.Success = false
			resp.Errors = []string{"invalid source: " + serviceProvider}
			b, _ := json.Marshal(resp)
			rw.Write(b)
			return
		}

	}
}
