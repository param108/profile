package twitter

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)


const (
	jsonContentType="application/json"
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

}

func setAuthorizationHeader(rw http.ResponseWriter) {

}

func NewTwitterLoginProvider() *TwitterLoginProvider {
	tlp := &TwitterLoginProvider{
		Client: http.Client{
			Timeout: time.Second * 30,
		},
	}

	return tlp
}

const (
	redirectAuthorizeURL=
		"https://twitter.com/i/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s"
	baseAuthorizeURL=
		"/users/authorize/twitter"
)

func (tlp *TwitterLoginProvider) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("TWITTER_CLIENT_ID")

	// The redirect url when authentication has succeeded.
	// scope: users.read
	// code_challenge: 16 char long random string
	myRedirectURI := fmt.Sprintf(
		"https://%s%s&scope=%s&state=state&code_challenge=%s&code_challenge_method=plain",
		os.Getenv("HOST"),
		baseAuthorizeURL,
		"users.read",
		RandStringBytesMask(16),
	)

	redirectURL := fmt.Sprintf(redirectAuthorizeURL,
		clientID,
		myRedirectURI,
	)

	http.Redirect(rw, r, redirectURL, 302)
}

const (
	getTokenURL="https://api.twitter.com/2/oauth2/token"
)
func (tlp *TwitterLoginProvider) HandlerAuthorize(rw http.ResponseWriter, r *http.Request) {

	request, err := tlp.Post(getTokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		rw.Header(http.StatusInternalServerError)
		rw.Write("Authentication Failed")
	}
	tlp.Post(url string, contentType string, body io.Reader)
}
