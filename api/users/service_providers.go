package users

import (
	"encoding/json"
	"net/http"

	"github.com/param108/profile/api/models"
)

func ServiceProviderRedirect(rw http.ResponseWriter, r *http.Request) {
	serviceProvider := r.URL.Query().Get("source")

	resp := models.Response{}

	switch serviceProvider {
	default:
		rw.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.Errors = []string{"invalid source: " + serviceProvider}
		b, _ := json.Marshal(resp)
		rw.Write(b)
		return
	}

}
