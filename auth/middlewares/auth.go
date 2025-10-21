package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	uu "github.com/mahdifarbehi/starme/auth/utils"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		authTokenValue := c.GetHeader("Authorization")
		parts := strings.SplitN(authTokenValue, " ", 2)
		if authTokenValue == "" || len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth token is missing or invalid"})
			return
		}
		authToken := parts[1]

		userID, err := uu.DecodeJWTToken(authToken)
		if err != nil {
			fmt.Printf("error in AuthRequired: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode token"})
		}

		c.Set("userID", userID)

		c.Next()
	}
}
