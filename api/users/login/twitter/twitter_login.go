package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/users/login/common"
	"github.com/param108/profile/api/utils"
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

	jwtString := r.Header.Get("TRIBIST_USER")
	// try and login using jwt
	if len(jwtString) > 0 {
		back := r.URL.Query().Get("redirect_url")

		// the redirect_url must begin with /
		if len(back) == 0 || !strings.HasPrefix(back, "/") {
			back = "/"
		}

		http.Redirect(rw, r, back, http.StatusTemporaryRedirect)
		return
	}

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
		url.QueryEscape("users.read tweet.read"),
		url.QueryEscape(key),
		url.QueryEscape(challenge))

	log.Println("redirect success", redirectURL)

	http.Redirect(rw, r, redirectURL, 302)
}

const (
	getTokenURL = "https://api.twitter.com/2/oauth2/token"
)

func (tlp *TwitterLoginProvider) getUsername(accessToken string) (string, error) {
	userDataURL := "https://api.twitter.com/2/users/me"

	req, err := http.NewRequest(http.MethodGet, userDataURL, strings.NewReader(""))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := tlp.Do(req)
	if err != nil {
		return "", err
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	val := map[string]interface{}{}
	if err := json.Unmarshal(d, &val); err != nil {
		return "", err
	}

	m, ok := val["data"].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid user data")
	}

	username, ok := m["username"].(string)
	if !ok {
		return "", errors.New("invalid user data")
	}

	return username, nil
}

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

	form := url.Values{}
	form.Add("code", code)
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", fmt.Sprintf("https://%s%s", os.Getenv("HOST"), baseAuthorizeURL))
	form.Add("code_verifier", challenge)

	req, err := http.NewRequest(http.MethodPost, getTokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	req.SetBasicAuth(os.Getenv("TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	loginData := map[string]interface{}{}
	// extract the access token
	err = json.Unmarshal(d, &loginData)
	if err != nil {
		log.Printf("failed to unmarshall loginData: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	accessToken, ok := loginData["access_token"].(string)
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Authentication Failed"))
		return
	}

	// got the access token with appropriate scope.
	// now we need to use that to get the user_id
	// after this we don't care for the accessToken
	// (atleast for now)
	username, err := tlp.getUsername(accessToken)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("failure"))
		return
	}

	// check if the user is already in the database.
	u, err := common.FindOrCreateTPUser(tlp.DB, username, "twitter", os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("failure"))
		return
	}

	token, err := utils.CreateSignedToken(u.Handle, u.ID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("create token failure"))
		return
	}

	cred := map[string]interface{}{
		"data": map[string]interface{}{
			"access_token": token,
		},
	}

	d, err = json.Marshal(cred)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("create token failure"))
		return
	}

	// userID, err := users.FindOrCreateTPAUser(username, "twitter")
	rw.WriteHeader(http.StatusOK)
	rw.Write(d)
}
