package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	ud "github.com/mahdifarbehi/starme/auth/dtos"
	um "github.com/mahdifarbehi/starme/auth/models"
	uu "github.com/mahdifarbehi/starme/auth/utils"
	cs "github.com/mahdifarbehi/starme/core"
	in "github.com/mahdifarbehi/starme/initializers"
)

var ErrLogin = errors.New("incorrect username or password")

func LoginUserHandler(data ud.UserLoginRequest) (string, error) {
	user := um.User{Username: data.Username}
	result := in.DB.First(&user)
	passwordVerified := uu.CheckPasswordHash(data.Password, user.Password)
	if result.Error != nil || !passwordVerified {
		log.Printf("error in LoginUserHandler: %v", result.Error)
		return "", ErrLogin
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cs.JWT_SECRET))
	if err != nil {
		log.Printf("error in LoginUserHandler: %v", err)
		return "", ErrLogin
	}
	return tokenString, nil
}
