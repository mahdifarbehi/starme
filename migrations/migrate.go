package main

import (
	"fmt"

	am "github.com/mahdifarbehi/starme/auth/models"
	"github.com/mahdifarbehi/starme/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	if err := initializers.DB.AutoMigrate(&am.User{}); err != nil {
		fmt.Println("Failed to migrate")
	}
}
