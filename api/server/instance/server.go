package serve

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/param108/profile/api/common"
	"github.com/param108/profile/api/spgroups"
	"github.com/param108/profile/api/spmsg"
	"github.com/param108/profile/api/spotps"
	"github.com/param108/profile/api/spservices"
	"github.com/param108/profile/api/spuser"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/threads"
	"github.com/param108/profile/api/tweets"
	"github.com/param108/profile/api/users"
	"github.com/param108/profile/api/users/login/twitter"
	"github.com/param108/profile/api/utils"
)

type Server struct {
	port int
	r    *mux.Router
	s    *http.Server
	DB   store.Store
	aws  *utils.AWS
	// periodicQuit External call to quit
	periodicQuit chan struct{}

	// periodicDone broadcast that periodic is done
	periodicDone chan struct{}
}

//go:embed version.txt
var version []byte

func NewServer(port int) (*Server, error) {
	server := &Server{}
	server.r = mux.NewRouter()

	server.periodicDone = make(chan struct{})
	server.periodicQuit = make(chan struct{})

	// FIXME: Turning off cors for development. Need to turn it back on.
	//allowedClients := strings.Split(os.Getenv("ALLOWED_CLIENTS"), ",")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "TRIBIST_JWT"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	server.s = &http.Server{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(server.r),
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if db, err := store.NewStore(); err != nil {
		return nil, err
	} else {
		server.DB = db
	}

	if aws, err := utils.NewAWS(); err != nil {
		return nil, err
	} else {
		server.aws = aws
	}

	// Must be done at the end
	server.RegisterHandlers()

	return server, nil
}

// Simple Hourly Periodic Jobs
func (s *Server) StartPeriodic() {
	ticker := time.NewTicker(time.Hour)
END_PERIODIC:
	for {
		select {
		case <-s.periodicQuit:
			break END_PERIODIC
		case <-ticker.C:
			twitter.NewTwitterLoginProvider(s.DB).Periodic()
			store.Periodic(s.DB, os.Getenv("WRITER"))
		}
	}
	ticker.Stop()
	close(s.periodicDone)
}

func (s *Server) Serve() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	go s.StartPeriodic()
	log.Fatal(s.s.ListenAndServe())
}

func (s *Server) Quit() {
	close(s.periodicQuit)
	<-s.periodicDone
	s.s.Close()
}

func (s *Server) RegisterHandlers() {
	s.r.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write(
			append(
				version,
				[]byte("\nwriter:"+os.Getenv("WRITER")+
					"\nHost:"+os.Getenv("HOST"))...))
	})

	s.r.HandleFunc("/users/login",
		utils.CheckM(users.CreateServiceProviderLoginRedirect(s.DB)).ServeHTTP).
		Methods(http.MethodGet)

	s.r.HandleFunc("/users/authorize/{service_provider}",
		users.CreateServiceProviderAuthorizeRedirect(s.DB))

	s.r.HandleFunc("/onetime",
		common.CreateGetOneTimeHandler(s.DB)).Methods(http.MethodGet)

	s.r.HandleFunc("/profile",
		utils.AuthM(
			users.CreateGetProfileHandler(s.DB)).ServeHTTP).
		Methods(http.MethodGet)

	s.r.HandleFunc("/tweets", tweets.CreateGetTweetsHandler(s.DB)).
		Methods(http.MethodGet)

	s.r.HandleFunc("/tweet/{tweet_id}", tweets.CreateGetATweetHandler(s.DB)).
		Methods(http.MethodGet)

	s.r.HandleFunc("/tweets",
		utils.AuthM(
			tweets.CreatePostTweetsHandler(s.DB),
		).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/tweets", utils.AuthM(
		tweets.CreateUpdateTweetHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPut)

	s.r.HandleFunc("/tweets/delete", utils.AuthM(
		tweets.CreateDeleteTweetHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/user/{username}/threads/{thread_id}",
		(threads.CreateGetThreadHandler(s.DB)).ServeHTTP).
		Methods(http.MethodGet)

	s.r.HandleFunc("/user/{username}/threads", utils.AuthM(
		threads.CreateMakeThreadHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/user/{username}/threads/{thread_id}/delete", utils.AuthM(
		threads.CreateDeleteThreadHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/sp/otp", (spotps.CreateMakeOTPHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	// Creates the sp user.
	s.r.HandleFunc("/sp/otp/check", (spotps.CreateCheckOTPHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/sp/users", utils.AuthSP(spuser.CreateGetSPUserHandler(s.DB)).ServeHTTP).
		Methods(http.MethodGet)

	s.r.HandleFunc("/sp/users/update", utils.AuthSP(spuser.CreateUpdateSPUserHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/sp/refresh", utils.AuthRefreshSP(spuser.CreateRefreshSPUserTokenHandler(s.DB)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/sp/upload_url", utils.AuthSP(
		spuser.CreatePutImageSignedUrlHandler(s.DB, s.aws)).ServeHTTP).
		Methods(http.MethodPost)

	s.r.HandleFunc("/sp/users/messages", utils.AuthSP(
		spmsg.CreateGetUserMessages(s.DB)).ServeHTTP).Methods(http.MethodGet)

	s.r.HandleFunc("/sp/users/messages", utils.AuthSP(
		spmsg.CreatePostUserMessages(s.DB)).ServeHTTP).Methods(http.MethodPost)

	s.r.HandleFunc("/sp/services", utils.AuthSP(
		spservices.CreateGetServices(s.DB)).ServeHTTP).Methods(http.MethodGet)

	s.r.HandleFunc("/sp/groups", utils.AuthSP(
		spgroups.CreatePostSPGroupHandler(s.DB)).ServeHTTP).Methods(http.MethodPost)

	s.r.HandleFunc("/sp/groups", utils.AuthSP(
		spgroups.CreateGetSPGroupsHandler(s.DB)).ServeHTTP).Methods(http.MethodGet)

	s.r.HandleFunc("/sp/group/{group_id}/users", utils.AuthSP(
		spgroups.CreateAddGroupUserHandler(s.DB)).ServeHTTP).Methods(http.MethodPost)

	s.r.HandleFunc("/sp/group/{group_id}/users", utils.AuthSP(
		spgroups.CreateGetSPGroupUsersHandler(s.DB)).ServeHTTP).Methods(http.MethodGet)

	s.r.HandleFunc("/sp/group/{group_id}/messages", utils.AuthSP(
		spmsg.CreateGetSPGroupMessagesHandler(s.DB)).ServeHTTP).Methods(http.MethodGet)

}
