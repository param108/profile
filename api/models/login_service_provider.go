package models

import "net/http"

type LoginServiceProvider interface {
	HandleLogin(rw http.ResponseWriter, r *http.Request)
	HandleAuthorize(rw http.ResponseWriter, r *http.Request)
}
