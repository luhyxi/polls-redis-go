// Package internal contains internal code and helpers for runnning miscelanious things
package internal

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateJwtToken(username string)  (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf": time.Now().Unix(),
	})
	
	tokenString, err := token.SignedString(GetJwtToken)
	
	if err != nil {
		log.Fatal("could not get jwt token")
		return "", errors.New("could not get jwt token")
	}

	return tokenString, nil
}
