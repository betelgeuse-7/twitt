package api_middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		bearerToken := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(bearerToken, keyFunc)
		if err != nil || !token.Valid {
			fmt.Println("auth middleware ->", err)
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
