package middleware

import (
	"fmt"
	initializers "go-expense-app/initilizers"
	"go-expense-app/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// get the token from the request
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(401)
		}

		// Find the user and attach it to request
		var user models.User
		initializers.DB.First(&user, claims["id"])

		if user.ID == 0 {
			c.AbortWithStatus(401)
		}

		c.Set("user", user)

		// if the token is valid, call the next handler
		c.Next()

	} else {
		c.AbortWithStatus(401)
	}

}
