package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ac "github.com/mahdifarbehi/starme/auth/controllers"
	"github.com/mahdifarbehi/starme/controllers"
	"github.com/mahdifarbehi/starme/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api is working",
		})
	})

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostReadAll)
	r.GET("/posts/:id", controllers.PostRead)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.POST("/users", ac.CreateUserAPI)
	r.POST("/login", ac.LoginUserAPI)

	if err := r.Run(); err != nil {
		fmt.Println("Failed to start the server")
	}
}
