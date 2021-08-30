package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/betelgeuse-7/twitt/api/helpers"
	"github.com/betelgeuse-7/twitt/db"
	"github.com/go-chi/chi/v5"
)

func GetTweet(w http.ResponseWriter, r *http.Request) {
	apiError := ApiError{}
	id64, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		apiError = ApiError{
			Code:    500,
			Title:   "error",
			Message: "internal server error",
		}
		apiError.Give(w)
		return
	}
	id := int(id64)
	// get tweet from db
	tweet := db.GetTweetById(id)

	if tweet.Id == 0 {
		apiError = ApiError{
			Code:    404,
			Title:   "no such tweet",
			Message: fmt.Sprintf("tweet with id %d does not exist\n", id),
		}
		apiError.Give(w)
		return
	}

	helpers.JSON(w, tweet)
}

func NewTweet(w http.ResponseWriter, r *http.Request) {
	var apiError ApiError
	var body struct {
		Content string
		Author  int
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Println(err)
		apiError = ApiError{
			Code:    500,
			Title:   "internal error",
			Message: "internal server error",
		}
		apiError.Give(w)
		return
	}
	if body.Content == "" || body.Author == 0 {
		apiError = ApiError{
			Code:    417,
			Title:   "error",
			Message: "check body",
		}
		apiError.Give(w)
	} else {
		lastInsertedId, err := db.NewTweet(body.Content, body.Author)
		if err != nil {
			fmt.Println(err)
			apiError = ApiError{
				Title:   "server error",
				Message: "internal server error",
				Code:    500,
			}
			apiError.Give(w)
			return
		}
		helpers.JSON(w, lastInsertedId)
	}
}
