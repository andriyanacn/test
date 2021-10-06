package middleware

import (
	"phone_review/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(constants.Secret_JWT))
}
