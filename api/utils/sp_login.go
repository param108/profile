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

type SPClaims struct {
	Phone  string `json:"phone"`
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func parseSPToken(jwtStr string) (*SPClaims, error) {
	claims := &SPClaims{}

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

func CreateSignedSPTokens(phone, userID string) (string, string, error) {
	timeNow := time.Now()
	claims := SPClaims{
		Phone:  phone,
		UserID: userID,
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
	accessToken, err := token.SignedString([]byte(os.Getenv("TRIBIST_JWT_KEY")))
	if err != nil {
		return "", "", err
	}

	refreshClaims := SPClaims{
		Phone:  phone,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			IssuedAt:  jwt.NewNumericDate(timeNow),
			NotBefore: jwt.NewNumericDate(timeNow),
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(24 * 7 * 24 * time.Hour)),
			Issuer:    "tribist",
			Subject:   "refresh",
			ID:        userID,
			Audience:  []string{"public"},
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshToken, err := token.SignedString([]byte(os.Getenv("TRIBIST_JWT_KEY")))

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// AuthSP delete SP_USER header and then
// repopulate with an empty value if unauthenticated,
// returns 401 or 403 if failure and does not proceed to handler.
func AuthSP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("SP_PHONE")
		r.Header.Del("SP_USERID")

		jwtStr := r.Header.Get("TRIBIST_JWT")

		claims, err := parseSPToken(jwtStr)
		if err != nil {
			if err.Error() == "unauthorized" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
			}

			if err.Error() == "forbidden" {
				http.Error(w, "forbidden", http.StatusForbidden)
			}
			return
		}

		// refresh token cannot be used for normal APIs
		if claims.Subject == "refresh" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		r.Header.Set("SP_USERID", claims.UserID)
		r.Header.Set("SP_PHONE", claims.Phone)
		next.ServeHTTP(w, r)
	})
}

// AuthRefreshSP delete SP_USER header and then
// repopulate with an empty value if unauthenticated,
// returns 401 or 403 if failure and does not proceed to handler.
// This middleware expects a refresh token,
// It will return 403 for access token.
func AuthRefreshSP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("SP_PHONE")
		r.Header.Del("SP_USERID")

		jwtStr := r.Header.Get("TRIBIST_JWT")

		claims, err := parseSPToken(jwtStr)
		if err != nil {
			if err.Error() == "unauthorized" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
			}

			if err.Error() == "forbidden" {
				http.Error(w, "forbidden", http.StatusForbidden)
			}
			return
		}

		// access token cannot be used for refresh
		if claims.Subject == "access" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		r.Header.Set("SP_USERID", claims.UserID)
		r.Header.Set("SP_PHONE", claims.Phone)
		next.ServeHTTP(w, r)
	})
}

// CheckSP will delete the SP_USER and will try and parse
// the SP_JWT header. If successful will set TRIBIST_USER.
// Even if authentication fails this middleware will call the handler.
func CheckSP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("SP_PHONE")
		r.Header.Del("SP_USERID")
		jwtStr := r.Header.Get("SP_JWT")
		if len(jwtStr) > 0 {
			claims, err := parseSPToken(jwtStr)
			if err == nil {
				r.Header.Set("SP_PHONE", claims.Phone)
				r.Header.Set("SP_USERID", claims.UserID)
			}
		}
		// Call the next handler anyway. It is expected that the next handler
		// will check the header itself.
		next.ServeHTTP(w, r)
	})
}
