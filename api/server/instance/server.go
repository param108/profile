package serve

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/param108/profile/api/users"
	"gorm.io/gorm"
)

type Server struct {
	port int
	r    *mux.Router
	s    *http.Server
	db   *gorm.DB
}

//go:embed version.txt
var version []byte

func NewServer(port int) (*Server, error) {
	server := &Server{}
	server.r = mux.NewRouter()

	server.RegisterHandlers()

	server.s = &http.Server{
		Handler: server.r,
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server, nil
}

func (s *Server) Serve() {
	log.Fatal(s.s.ListenAndServe())
}

func (s *Server) Quit() {
	s.s.Close()
}

func (s *Server) RegisterHandlers() {
	s.r.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write(version)
	})

	s.r.HandleFunc("/users/login", users.ServiceProviderRedirect).Methods(http.MethodGet)
}
