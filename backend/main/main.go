package main

import (
	"../db"
	"../middlewares"
	"../routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func init() {
	db.Connect()
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // TODO change to other port
	}

	router := gin.Default()
	router.RedirectTrailingSlash = false

	// Middlewares
	router.Use(middlewares.ConnectMiddleware)
	router.Use(middlewares.ErrorHandler)

	//CORS
	router.Use(CORS())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	api := router.Group("/api")
	routers.UsersApi(api.Group("/users"))
	routers.PostApi(api.Group("/posts"))
	router.Run(":" + port)
}
