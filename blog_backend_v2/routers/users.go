package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersApi(router *gin.RouterGroup) {
	router.POST("/authenticate", UserAuthenticateHandler)
}

type LoginCredentials struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Authenticate user by login/password
func UserAuthenticateHandler(c *gin.Context) {
	var LoginCredentials LoginCredentials
	var mySigningKey = []byte("c2VjcmV0Cg==")

	if err := c.BindJSON(&LoginCredentials); err != nil {
		c.Error(err)
		return
	}

	if LoginCredentials.Login != "admin" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if LoginCredentials.Password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": "admin",
		})
		tokenString, _ := token.SignedString(mySigningKey)
		c.JSON(http.StatusCreated, gin.H{
			"jwt": tokenString,
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Invalid password",
		})
		return
	}
}
