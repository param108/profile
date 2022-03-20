package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/param108/profile/api/models"
)

func ServiceProviderRedirect(rw http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	resp := models.Response{}

	if serviceProvider, ok := v["source"]; ok {
		switch serviceProvider {
		default:
			rw.WriteHeader(http.StatusBadRequest)
			resp.Success = false
			resp.Errors = []string{"invalid source: " + serviceProvider}
			b, _ := json.Marshal(resp)
			rw.Write(b)
			return
		}
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.Errors = []string{"missing source"}
		b, _ := json.Marshal(resp)
		rw.Write(b)
		return
	}

}
