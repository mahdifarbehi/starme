package services

import (
	"errors"
	"log"

	ud "github.com/mahdifarbehi/starme/auth/dtos"
	um "github.com/mahdifarbehi/starme/auth/models"
	uu "github.com/mahdifarbehi/starme/auth/utils"
	in "github.com/mahdifarbehi/starme/initializers"
)

var ErrDB = errors.New("database operation failed")

func CreateUserHandler(data ud.UserCreateRequest) (um.User, error) {
	hashed_password, _ := uu.HashPassword(data.Password)
	user := um.User{Username: data.Username, Password: hashed_password}
	result := in.DB.Create(&user)
	if result.Error != nil {
		log.Printf("DB error in CreateUserHandler: %v", result.Error)
		return user, ErrDB
	}
	return user, nil
}
