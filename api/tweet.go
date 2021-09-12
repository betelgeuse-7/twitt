package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/betelgeuse-7/twitt/api/helpers"
	"github.com/betelgeuse-7/twitt/db"
	"github.com/go-chi/chi/v5"
)

const CONSTRAINT_LIKE_ONCE = "likes_who_liked_fkey"

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
		GiveInternalServerError(w)
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
			GiveInternalServerError(w)
			return
		}
		helpers.JSON(w, lastInsertedId)
	}
}

func LikeTweet(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		fmt.Println("type assertion failed..")
		GiveInternalServerError(w)
		return
	}
	tweetId, err := helpers.StrToInt(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id must be an int", http.StatusBadRequest)
		return
	}
	err = db.LikeTweet(tweetId, userId)
	if !strings.Contains(err.Error(), CONSTRAINT_LIKE_ONCE) {
		GiveInternalServerError(w)
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

func GetCommentsForTweet(w http.ResponseWriter, r *http.Request) {}

func NewCommentForTweet(w http.ResponseWriter, r *http.Request) {}

func GetTweetsWithHashtag(w http.ResponseWriter, r *http.Request) {}
