package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/betelgeuse-7/twitt/api/helpers"
	"github.com/betelgeuse-7/twitt/db"
	"github.com/go-chi/chi/v5"
)

func GetTweet(w http.ResponseWriter, r *http.Request) {
	apiError := ApiError{}
	id, err := helpers.StrToInt(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "please provide an id (int)", http.StatusBadRequest)
		return
	}
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
	userId := r.Context().Value("userId").(int)
	var apiError ApiError
	var body struct {
		Content string
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
	if body.Content == "" {
		apiError = ApiError{
			Code:    417,
			Title:   "error",
			Message: "check body",
		}
		apiError.Give(w)
	} else {
		lastInsertedId, err := db.NewTweet(body.Content, userId)
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

// {"userid": 1}
func LikeTweet(w http.ResponseWriter, r *http.Request) {
	tweetId, err := helpers.StrToInt(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id must be an int", http.StatusBadRequest)
		return
	}
	var apiError ApiError
	var body struct {
		UserId int
	}
	json.NewDecoder(r.Body).Decode(&body)
	err = db.LikeTweet(tweetId, body.UserId)
	if err != nil {
		apiError = ApiError{
			Title:   "internal error",
			Message: "server error",
			Code:    500,
		}
		apiError.Give(w)
		return
	}
	helpers.JSON(w, map[string]string{
		"message": fmt.Sprintf("liked tweet with id %d", tweetId),
	})
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)
	id, err := helpers.StrToInt(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "provide an id (int)", http.StatusBadRequest)
		return
	}
	tweet := db.GetTweetById(id)
	if tweet.Author != userId {
		http.Error(w, "you are not permissible to delete this tweet", http.StatusUnauthorized)
		return
	}
	err = db.DeleteTweet(id)
	if err != nil {
		ApiError{
			Title:   "internal error",
			Message: "error while deleting the tweet",
			Code:    500,
		}.Give(w)
		return
	}
	helpers.JSON(w, map[string]string{
		"message": "deleted tweet!",
	})
}
