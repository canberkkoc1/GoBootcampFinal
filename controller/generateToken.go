package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email, role string) (string, error) {

	// token key from env

	tokenKey := os.Getenv("MY_TOKEN_KEY")

	// create token

	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["isAdmin"] = role
	claims["isLogin"] = false
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// generate token
	tokenString, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		fmt.Errorf("something went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func GenerateLoginJWT(email, role string) (string, error) {

	tokenKey := os.Getenv("MY_TOKEN_KEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["isLogin"] = true
	claims["isAdmin"] = role
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		fmt.Errorf("something went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil

}
