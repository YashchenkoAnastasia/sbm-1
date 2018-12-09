// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"../db"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func ConnectMiddleware(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}

func AuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.Split(tokenString, "Bearer ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("c2VjcmV0Cg=="), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user"] != "admin" {
			c.JSON(http.StatusForbidden, nil)
			return
		}
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
		return
	}
	c.Next()
}
