package serve

import (
	"fmt"
	"log"
	"net/http"
	"time"
	_ "embed"
	"github.com/gorilla/mux"
)

type Server struct {
	port int
	r *mux.Router
	s *http.Server
}


//go:embed version.txt
var version []byte

func NewServer(port int) (*Server,error) {
	server := &Server{}
	server.r = mux.NewRouter()

	server.RegisterHandlers()

	server.s = &http.Server{
		Handler: server.r,
		Addr:    fmt.Sprintf("127.0.0.1:%d",port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server, nil
}

func (s *Server) Serve() {
	log.Fatal(s.s.ListenAndServe())
}

func (s *Server) RegisterHandlers() {
	s.r.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write(version)
	})
}
