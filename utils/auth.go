package utils

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// AuthMiddleware a middleware to check if the request has a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			logrus.Error("Error loading .env file")
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			logrus.Error("JWT_SECRET not found in .env file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			logrus.Error("Error parsing JWT token: ", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			logrus.Error("Invalid JWT token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		context.Set(r, "user", token.Claims.(jwt.MapClaims)["user"])
		next.ServeHTTP(w, r)
	})
}

// GetUserIDFromToken returns the user ID from the JWT token
func GetUserIDFromToken(r *http.Request) uint {
	user := context.Get(r, "user")
	userID := uint(user.(float64))
	return userID
}

// GetUsernameFromToken returns the username from the JWT token
func GetUsernameFromToken(r *http.Request) string {
	user := context.Get(r, "user")
	username := user.(string)
	return username
}

// GetRoleFromToken returns the role from the JWT token
func GetRoleFromToken(r *http.Request) string {
	user := context.Get(r, "user")
	role := user.(jwt.MapClaims)["role"].(string)
	return role
}

// GetTokenFromRequest returns the JWT token from the request
func GetTokenFromRequest(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	return tokenString
}

