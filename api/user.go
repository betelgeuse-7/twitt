package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/betelgeuse-7/twitt/api/helpers"
	"github.com/betelgeuse-7/twitt/db"
	chi "github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var apiError ApiError
	// get user data
	userData := struct {
		Username string
		Email    string
		Password string
		Handle   string
	}{}
	json.NewDecoder(r.Body).Decode(&userData)
	// check user data
	if err := helpers.CheckNewUserCreds(userData.Username, userData.Password, userData.Email, userData.Handle); err != nil {
		apiError = ApiError{
			Title:   "invalid input",
			Message: "check your input",
			Code:    400,
		}
		apiError.Give(w)
		return
	}
	// hash password
	pwd, err := helpers.HashPassword(userData.Password)
	if err != nil {
		log.Println(err)
		apiError = ApiError{
			Title:   "internal error",
			Message: "internal server error",
			Code:    500,
		}
		apiError.Give(w)
		return
	}
	userData.Password = pwd
	// save to db
	err = db.NewUser(userData.Username, userData.Password, userData.Email, userData.Handle)
	if err != nil {
		log.Println(err)
		apiError = ApiError{
			Title:   "internal error",
			Message: "error while registering",
			Code:    500,
		}
		apiError.Give(w)
		return
	}
	helpers.JSON(w, map[string]string{
		"message": "successfully registered",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var apiError ApiError
	// get user data (email, password)
	var user struct {
		Email    string
		Password string
	}
	json.NewDecoder(r.Body).Decode(&user)
	// check user input
	if err := helpers.CheckLoginInput(user.Email, user.Password); err != nil {
		apiError = ApiError{
			Title:   "login failed",
			Message: "login failed - check your inputs",
			Code:    401,
		}
		apiError.Give(w)
		return
	}
	// find user with email
	userFromDb, err := db.GetUserByEmail(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			apiError = ApiError{
				Title:   "no such user",
				Message: "invalid input",
				Code:    401,
			}
		} else {
			apiError = ApiError{
				Title:   "internal error",
				Message: "internal server error",
				Code:    500,
			}
		}
		apiError.Give(w)
		return
	}
	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password)); err != nil {
		apiError = ApiError{
			Title:   "incorrect password",
			Message: "incorrect password.",
			Code:    401,
		}
		apiError.Give(w)
		return
	}
	// generate jwt
	jwt, err := NewJWT(int(userFromDb.Id))
	if err != nil {
		apiError = ApiError{
			Title:   "internal error",
			Message: "internal server error",
			Code:    500,
		}
		apiError.Give(w)
		return
	}
	// give jwt
	helpers.JSON(w, map[string]string{
		"token": jwt,
	})
}

func LikedTweets(w http.ResponseWriter, r *http.Request) {
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")
	idStr := chi.URLParam(r, "id")
	if offsetStr == "" || limitStr == "" {
		fmt.Fprint(w, "please include offset and limit parameters.\n")
		return
	}
	offset, _ := helpers.StrToInt(offsetStr)
	limit, _ := helpers.StrToInt(limitStr)
	id, err := helpers.StrToInt(idStr)
	if err != nil {
		http.Error(w, "please provide an id (int)", http.StatusBadRequest)
		return
	}
	tweets, err := db.GetUserLikedTweets(id, offset, limit)
	if err != nil {
		ApiError{
			Title:   "internal error",
			Message: "internal server error",
			Code:    500,
		}.Give(w)
		return
	}
	helpers.JSON(w, tweets)
}
