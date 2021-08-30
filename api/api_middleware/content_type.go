package api_middleware

import (
	"net/http"
	"strings"
)

/*
Middleware to check whether POST and PUT requests' headers include Content-Type=application/json or not.
*/
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := strings.ToLower(r.Header.Get("Content-Type"))
		if contentType != "" {
			if contentType != "application/json" {
				http.Error(w, "invalid content type", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
