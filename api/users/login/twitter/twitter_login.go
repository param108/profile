package twitter

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/param108/profile/api/store"
)

const (
	jsonContentType = "application/json"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

type TwitterLoginProvider struct {
	http.Client
	DB store.Store
}

func NewTwitterLoginProvider(db store.Store) *TwitterLoginProvider {
	tlp := &TwitterLoginProvider{
		Client: http.Client{
			Timeout: time.Second * 30,
		},
		DB: db,
	}

	return tlp
}

const (
	redirectAuthorizeURL = "https://twitter.com/i/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s&code_challenge=%s&code_challenge_method=plain"
	baseAuthorizeURL     = "/users/authorize/twitter"
)

func (tlp *TwitterLoginProvider) Periodic() {
	tlp.DB.DeleteOldTwitterChallenges(time.Hour)
}

func (tlp *TwitterLoginProvider) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("TWITTER_CLIENT_ID")

	challenge := RandStringBytesMask(16)

	key, err := tlp.DB.CreateTwitterChallenge(challenge, os.Getenv("WRITER"))
	if err != nil {
		log.Println("login failed:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("login failed"))
		return
	}

	// The redirect url when authentication has succeeded.
	// scope: users.read
	// code_challenge: 16 char long random string
	redirectURL := fmt.Sprintf(redirectAuthorizeURL,
		url.QueryEscape(clientID),
		url.QueryEscape("https://"+os.Getenv("HOST")+baseAuthorizeURL),
		url.QueryEscape("users.read"),
		url.QueryEscape(key),
		url.QueryEscape(challenge))

	log.Println("redirect success", redirectURL)

	http.Redirect(rw, r, redirectURL, 302)
}

const (
	getTokenURL = "https://api.twitter.com/2/oauth2/token"
)

func (tlp *TwitterLoginProvider) HandleAuthorize(rw http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if len(code) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	state := r.URL.Query().Get("state")
	if len(state) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	challenge, err := tlp.DB.GetTwitterChallenge(state, os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	req, err := http.NewRequest(http.MethodPost, getTokenURL, nil)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Authentication Failed"))
		return
	}
	values := req.PostForm
	values.Add("code", code)
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", fmt.Sprintf("https://%s%s", os.Getenv("HOST"), baseAuthorizeURL))
	values.Add("code_verifier", challenge)

	req.SetBasicAuth(os.Getenv("TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"))
	resp, err := tlp.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(d)
}
