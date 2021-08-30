package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/betelgeuse-7/twitt/api/helpers"
	"github.com/betelgeuse-7/twitt/db"
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
