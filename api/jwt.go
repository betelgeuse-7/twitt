package api

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const EXP = int64(time.Hour * 3)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func NewJWT(userId int) (string, error) {
	j := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: EXP,
		},
	}
	jwtTokenToSign := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	jwtToken, err := jwtTokenToSign.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func GetUserIdFromJWT(token string) (int, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
	})
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return claims.UserId, nil
}