package api

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const EXP = int64(time.Hour * 3)

type JWT struct {
	Claims jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewJWT(userId int) (string, error) {
	j := JWT{
		Claims: jwt.StandardClaims{
			ExpiresAt: EXP,
		},
		UserId: userId,
	}
	jwtTokenToSign := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Claims)
	jwtToken, err := jwtTokenToSign.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
