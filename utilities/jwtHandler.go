package utilities

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const secret = "!$aFF_WQDAX>>!(#*["

var UserID float64

var hmacSampleSecret []byte

func JWT(userId uint) (string, string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))

	if err != nil {
		return "", "", err
	}

	return t, rt, nil
}

//check if token es verify
//receive token from header, return bool if correct or not and user id
func CheckToken(tokenReceived string) (bool, float64) {
	token, err := jwt.Parse(tokenReceived, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return true, claims["user_id"].(float64)
	}

	return false, 0
}
