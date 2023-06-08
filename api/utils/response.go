package utils

import (
	"encoding/json"
	"net/http"
)

func WriteError(rw http.ResponseWriter, status int, errorStr string) {
	rw.WriteHeader(status)
	resp := map[string]interface{}{
		"success": false,
		"error":   errorStr,
	}
	d, err := json.Marshal(resp)
	if err != nil {
		return
	}
	rw.Write(d)
}

func WriteData(rw http.ResponseWriter, status int, data interface{}) {
	resp := map[string]interface{}{
		"success": true,
		"error":   "",
		"data":    data,
	}
	d, err := json.Marshal(resp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(status)
	rw.Write(d)
}
