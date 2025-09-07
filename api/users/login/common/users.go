package common

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

func FindOrCreateTPUser(db store.Store, handle, profile, writer string) (*models.User, error) {
	return db.FindOrCreateUser(handle, profile, models.RoleUser, writer)
}

// LoginUser Logs in the user and redirects to the configured
// redirect url with onetime and final redirect_url as parameters
func LoginUser(rw http.ResponseWriter, r *http.Request, db store.Store,
	handle, userID, savedRedirect string) {
	token, err := utils.CreateSignedToken(handle, userID)
	if err != nil {
		log.Printf("Failed to create token %s\n", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("create token failure"))
		return
	}

	// Now we have a token but we don't want to send it to
	// a frontend as url parameter, so we send a onetime token
	// instead which can be exchanged for the actual token.

	oneTime, err := db.SetOneTime(token, os.Getenv("WRITER"))
	if err != nil {
		log.Printf("Failed to create OneTime %s\n", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("create onetime failure"))
		return
	}

	// TODO this url should not be hardcoded.
	redirectURL, err := url.Parse(os.Getenv("AUTH_REDIRECT_URL"))
	if err != nil {
		log.Printf("Invalid redirectURL %s\n", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("redirect url"))
		return
	}

	// Add oneTime as onetime parameter
	params := redirectURL.Query()

	params.Set("onetime", oneTime.ID)
	params.Set("redirect_url", savedRedirect)
	redirectURL.RawQuery = params.Encode()

	// Set CORS header before redirecting
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	http.Redirect(rw, r, redirectURL.String(), http.StatusFound)
}
