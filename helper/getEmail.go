package helper

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetEmail(c *gin.Context) string {

	var userEmail string

	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return ""
	} else {

		token, err := jwt.Parse(c.Request.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("MY_SECRET_KEY")), nil

		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return ""
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userEmail = claims["email"].(string)
			return userEmail
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not admin"})
			c.Abort()
			return ""
		}
	}

}
