package email

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/users/login/common"
	"golang.org/x/crypto/bcrypt"
)

type EmailPasswordProvider struct {
	DB store.Store
}

func NewEmailPasswordProvider(db store.Store) *EmailPasswordProvider {
	return &EmailPasswordProvider{
		DB: db,
	}
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var validUsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func isValidRedirectURL(redirectURL string) bool {
	return strings.HasPrefix(redirectURL, "/")
}

func (epp *EmailPasswordProvider) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	user := strings.Split(r.Header.Get("TRIBIST_USER"), ":")
	userID := user[0]

	redirectURL := r.URL.Query().Get("redirect_url")

	// If user is already logged in
	if len(userID) > 0 {
		if isValidRedirectURL(redirectURL) {
			// Add CORS headers for cross-origin redirects
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS")
			rw.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, TRIBIST_JWT")
			
			http.Redirect(rw, r, redirectURL, http.StatusTemporaryRedirect)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("invalid redirect"))
		return
	}

	// If user is not logged in, set up redirect flow
	oneTime, err := epp.DB.SetOneTime(redirectURL, os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("failed to create session"))
		return
	}

	loginURL := fmt.Sprintf("https://ui.tribist.com/login?key=%s", oneTime.ID)
	
	// Add CORS headers for cross-origin redirects
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, TRIBIST_JWT")
	
	http.Redirect(rw, r, loginURL, http.StatusTemporaryRedirect)
}

func (epp *EmailPasswordProvider) HandleAuthorize(rw http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	code := r.URL.Query().Get("code")

	// If key or code are empty, redirect to home
	if key == "" || code == "" {
		// Add CORS headers for cross-origin redirects
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, TRIBIST_JWT")
		
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Handle email login flow with key and code
	epp.handleEmailLoginAuthorize(rw, r, key, code)
}

func (epp *EmailPasswordProvider) handleRegistrationForm(rw http.ResponseWriter, r *http.Request) {
	redirectURL := r.URL.Query().Get("redirect_url")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Register</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<style>
			body { font-family: Arial, sans-serif; max-width: 400px; margin: 50px auto; padding: 20px; }
			.form-group { margin-bottom: 15px; }
			label { display: block; margin-bottom: 5px; font-weight: bold; }
			input { width: 100%%; padding: 8px; border: 1px solid #ddd; border-radius: 4px; }
			button { background: #007cba; color: white; padding: 10px 20px; border: none; border-radius: 4px; cursor: pointer; }
			button:hover { background: #005a87; }
			.error { color: red; margin-top: 10px; }
		</style>
	</head>
	<body>
		<h2>Create Account</h2>
		<form method="POST" action="/users/authorize/email?redirect_url=%s">
			<div class="form-group">
				<label for="username">Username (alphanumeric and underscore only):</label>
				<input type="text" id="username" name="username" pattern="[a-zA-Z0-9_]+" required>
			</div>
			<div class="form-group">
				<label for="password">Password:</label>
				<input type="password" id="password" name="password" required minlength="6">
			</div>
			<button type="submit">Create Account</button>
		</form>
		
		<p>Already have an account? <a href="/users/login?source=email&redirect_url=%s">Sign in</a></p>
	</body>
	</html>
	`, redirectURL, redirectURL)

	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte(html))
}

func (epp *EmailPasswordProvider) handleRegistration(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Invalid form data"},
		})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"All fields are required"},
		})
		return
	}

	if !validUsernameRegex.MatchString(username) {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Username must contain only alphanumeric characters and underscores"},
		})
		return
	}

	if len(password) < 6 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Password must be at least 6 characters"},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Failed to process password"},
		})
		return
	}

	emailUser, err := epp.DB.CreateEmailUser(username, string(hashedPassword), os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"User creation failed: " + err.Error()},
		})
		return
	}

	// Find or create user with username as handle
	user, err := common.FindOrCreateTPUser(epp.DB, emailUser.UserName, "email", os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Login failed"},
		})
		return
	}

	back := r.URL.Query().Get("redirect_url")
	common.LoginUser(rw, r, epp.DB, user.Handle, user.ID, back)
}

func (epp *EmailPasswordProvider) handleEmailLoginAuthorize(rw http.ResponseWriter, r *http.Request, key, code string) {

	// Get the original redirect URL using the key
	keyOneTime, err := epp.DB.GetOneTime(key, time.Hour, os.Getenv("WRITER"))
	if err != nil {
		log.Printf("couldnt get key %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Cant get key" + err.Error()},
		})
		return
	}

	redirectURL := keyOneTime.Data

	// Get the user info using the code
	codeOneTime, err := epp.DB.GetOneTime(code, time.Hour, os.Getenv("WRITER"))
	if err != nil {
		log.Printf("couldnt get code %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Cant get code" + err.Error()},
		})
		return
	}

	// Parse the user JSON from the code
	var userPayload map[string]string
	if err := json.Unmarshal([]byte(codeOneTime.Data), &userPayload); err != nil {
		log.Printf("couldnt get payload %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Cant get payload" + err.Error()},
		})
		return
	}

	username, ok := userPayload["user"]
	if !ok {
		log.Println("couldnt get user")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Cant get user"},
		})
		return
	}

	// Find or create user with username as handle
	user, err := common.FindOrCreateTPUser(epp.DB, username, "email", os.Getenv("WRITER"))
	if err != nil {
		log.Printf("couldnt create TP User %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Cant get TP user"},
		})
		return
	}

	// Log the user in and redirect to original URL
	common.LoginUser(rw, r, epp.DB, user.Handle, user.ID, redirectURL)
}

func (epp *EmailPasswordProvider) HandleEmailLogin(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Method not allowed"},
		})
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Invalid request body"},
		})
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Missing key parameter"},
		})
		return
	}

	if !validUsernameRegex.MatchString(loginReq.UserName) {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Invalid username format"},
		})
		return
	}

	emailUser, err := epp.DB.GetEmailUserByUserName(loginReq.UserName, os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Invalid credentials"},
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(emailUser.PasswordHash), []byte(loginReq.Password)); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Invalid credentials"},
		})
		return
	}

	// Create JSON payload with username
	userPayload := map[string]string{"user": emailUser.UserName}
	userJSON, err := json.Marshal(userPayload)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Failed to create user payload"},
		})
		return
	}

	// Store the user JSON as a OneTime code
	codeOneTime, err := epp.DB.SetOneTime(string(userJSON), os.Getenv("WRITER"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(models.Response{
			Success: false,
			Errors:  []string{"Failed to create code"},
		})
		return
	}

	// Redirect to /users/authorize/email with key and code parameters
	redirectURL := fmt.Sprintf("/users/authorize/email?key=%s&code=%s", key, codeOneTime.ID)
	
	// Add CORS headers for cross-origin redirects
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, TRIBIST_JWT")
	
	http.Redirect(rw, r, redirectURL, http.StatusTemporaryRedirect)
}
