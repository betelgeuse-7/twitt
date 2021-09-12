package api_middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/betelgeuse-7/twitt/api"
	"github.com/dgrijalva/jwt-go"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		lastOfPath := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
		// "new" may cause bugs
		// because we can use it later, for creating a new tweet.
		if lastOfPath == "new" || lastOfPath == "login" {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		if len(strings.Split(authHeader, " ")) != 2 {
			http.Error(w, "not authorized (bad Authorization header)", http.StatusUnauthorized)
			return
		}
		bearerToken := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(bearerToken, keyFunc)
		if err != nil || !token.Valid {
			fmt.Println("auth middleware ->", err)
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		id, err := api.GetUserIdFromJWT(bearerToken)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, "userId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
