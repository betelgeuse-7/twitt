package api

import (
	"net/http"

	"github.com/betelgeuse-7/twitt/api/helpers"
)

type ApiError struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Code    int    `json:"status_code"`
}

// give api error
// set status code and send message
func (ae ApiError) Give(w http.ResponseWriter) {
	w.WriteHeader(ae.Code)
	helpers.JSON(w, ae)
}
