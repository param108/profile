package serve

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Serve(port int) {
	router := mux.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("127.0.0.1:%d",port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

    log.Fatal(srv.ListenAndServe())
}
