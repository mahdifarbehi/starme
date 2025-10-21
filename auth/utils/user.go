package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	cs "github.com/mahdifarbehi/starme/core"
	"golang.org/x/crypto/bcrypt"
)

var ErrDecodeJWT = errors.New("error in DecodeJWTToken")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func DecodeJWTToken(authToken string) (uint, error) {

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (any, error) {
		return []byte(cs.JWT_SECRET), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		fmt.Printf("error in DecodeJWTToken: %v", err)
		return 0, ErrDecodeJWT
	}

	var userID uint
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		idFloat, ok := claims["user_id"].(float64)
		if !ok {
			fmt.Println("error in ")
			return 0, ErrDecodeJWT
		}
		userID = uint(idFloat)
	} else {
		fmt.Printf("error in DecodeJWTToken: %v", err)
		return 0, ErrDecodeJWT
	}
	return userID, nil
}
