package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/users/login/common"
	"github.com/param108/profile/api/users/login/twitter"
)

func loginDevUser(rw http.ResponseWriter, r *http.Request, db store.Store) {
	fmt.Println("Logging in dev user")
	common.LoginUser(rw, r, db, "param108",
		"a592e6ab-91d1-49a7-9435-ab3c04f77ab9", r.URL.Query().Get("redirect_url"))
}

func CreateServiceProviderLoginRedirect(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		env := os.Getenv("ENV")

		if env == "dev" {
			// We have only one dev user.
			loginDevUser(rw, r, db)
			return
		}

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
