package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	jwt.RegisteredClaims
}

func parseToken(jwtStr string) (*Claims, error) {
	claims := &Claims{}

	t, err := jwt.ParseWithClaims(
		strings.TrimSpace(jwtStr),
		claims,
		func(token *jwt.Token) (interface{}, error) {
			// Signing Key will come from env
			jwtKey := os.Getenv("TRIBIST_JWT_KEY")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.New("unauthorized")
	}

	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}

	validErr := t.Claims.Valid()
	if err != nil || !t.Valid || validErr != nil {
		log.Printf("Invalid token %s - err: %s claims err: %s",
			jwtStr,
			err.Error(),
			validErr.Error(),
		)
		return nil, errors.New("forbidden")
	}

	return claims, nil
}

// AuthM delete TRIBIST_USER header and then
// repopulate with an empty value if unauthenticated,
// returns 401 or 403 if failure and does not proceed to handler.
func AuthM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("TRIBIST_USER")
		r.Header.Del("TRIBIST_USERID")

		jwtStr := r.Header.Get("TRIBIST_JWT")

		claims, err := parseToken(jwtStr)
		if err != nil {
			if err.Error() == "unauthorized" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
			}

			if err.Error() == "forbidden" {
				http.Error(w, "forbidden", http.StatusForbidden)
			}
			return
		}
		r.Header.Set("TRIBIST_USERID", claims.UserID)
		r.Header.Set("TRIBIST_USER", claims.Username)
		next.ServeHTTP(w, r)
	})
}

// CheckM will delete the TRIBIST_USER and will try and parse
// the TRIBIST_JWT header. If successful will set TRIBIST_USER.
// Even if authentication fails this middleware will call the handler.
func CheckM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("TRIBIST_USER")
		r.Header.Del("TRIBIST_USERID")
		jwtStr := r.Header.Get("TRIBIST_JWT")
		if len(jwtStr) > 0 {
			claims, err := parseToken(jwtStr)
			if err == nil {
				r.Header.Set("TRIBIST_USER", claims.Username)
				r.Header.Set("TRIBIST_USERID", claims.UserID)
			}
		}
		// Call the next handler anyway. It is expected that the next handler
		// will check the header itself.
		next.ServeHTTP(w, r)
	})
}

func CreateSignedToken(username, userID string) (string, error) {
	timeNow := time.Now()
	claims := Claims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			IssuedAt:  jwt.NewNumericDate(timeNow),
			NotBefore: jwt.NewNumericDate(timeNow),
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(24 * time.Hour * 7)),
			Issuer:    "tribist",
			Subject:   "access",
			ID:        userID,
			Audience:  []string{"public"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(os.Getenv("TRIBIST_JWT_KEY")))
}
