package services

import (
	"errors"
	"log"

	um "github.com/mahdifarbehi/starme/auth/models"
	in "github.com/mahdifarbehi/starme/initializers"
	"gorm.io/gorm"
)

var ErrGetMe = errors.New("failed to get me")

func GetMeHandler(userID uint) (um.User, error) {
	user := um.User{Model: gorm.Model{ID: userID}}
	result := in.DB.First(&user)
	if result.Error != nil {
		log.Printf("error in AuthRequired: %v", result.Error)
		return user, ErrGetMe
	}
	return user, nil
}
