package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWTAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {

		var checkAdmin string
		if c.Request.Header.Get("Authorization") != "" {

			token, err := jwt.Parse(c.Request.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("MY_SECRET_KEY")), nil

			})

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				checkAdmin = claims["isAdmin"].(string)
				if checkAdmin == "N" {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not admin"})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not admin"})
				c.Abort()
				return
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

	}

}

func AuthLogin() gin.HandlerFunc {

	return func(c *gin.Context) {

		var isLogin bool
		if c.Request.Header.Get("Authorization") != "" {

			token, err := jwt.Parse(c.Request.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("MY_SECRET_KEY")), nil

			})

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				isLogin = claims["isLogin"].(bool)
				if !isLogin {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not login"})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not login"})
				c.Abort()
				return
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

	}

}
