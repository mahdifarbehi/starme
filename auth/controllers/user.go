package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ud "github.com/mahdifarbehi/starme/auth/dtos"
	us "github.com/mahdifarbehi/starme/auth/services"
)

func CreateUserAPI(c *gin.Context) {
	var request ud.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := us.CreateUserHandler(request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func LoginUserAPI(c *gin.Context) {
	var request ud.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := us.LoginUserHandler(request)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
