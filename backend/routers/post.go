package routers

import (
	"../middlewares"
	"../models"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func PostApi(router *gin.RouterGroup) {
	router.GET("", PostListHandler)
	router.GET("/:id", PostShowHandler)
	router.Use(middlewares.AuthMiddleware)
	{
		router.POST("", PostCreateHandler)
		router.PUT("/:id", PostUpdateHandler)
		router.DELETE("/:id", PostDeleteHandler)
	}
}

type PostCredentials struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// Get post list
func PostListHandler(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	posts := make([]models.Post, 0)
	err := db.C(models.CollectionPost).Find(nil).All(&posts)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// Get an post
func PostShowHandler(c *gin.Context) {
	var flag int
	id := c.Param("id")
	if id != "" && bson.IsObjectIdHex(id) {
		db := c.MustGet("db").(*mgo.Database)
		post := models.Post{}
		err := db.C(models.CollectionPost).FindId(bson.ObjectIdHex(id)).One(&post)
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Post not found",
				})
			} else {
				c.Error(err)
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"post": post,
		})
	} else {
		c.Error(errors.New("Invalid post's id"))
	}
}

// Create an user
func PostCreateHandler(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	post := models.Post{}
	err := c.Bind(&post)
	post.Id = bson.NewObjectId()
	if err != nil {
		c.Error(err)
		return
	}
	err = post.ValidateNewPost()
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionPost).Insert(post)

	if err != nil {
		switch err.(type) {
		case *mgo.QueryError:
			c.Error(errors.New("Database query error"))
		case *mgo.LastError:
			c.Error(err)
		default:
			c.Error(err)
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

// Update a post
func PostUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	if id != "" && bson.IsObjectIdHex(id) {

		// Searching for an existing user in DB
		db := c.MustGet("db").(*mgo.Database)
		postFromDB := models.Post{}
		err := db.C(models.CollectionPost).FindId(bson.ObjectIdHex(id)).One(&postFromDB)
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, nil)
			} else {
				c.Error(err)
			}
			return
		}

		// Binding new values from incoming JSON
		var postJSON models.Post
		err = c.BindJSON(&postJSON)
		if err != nil {
			c.Error(err)
			return
		}

		postJSON.Id = postFromDB.Id

		// Merging two structs: existing post and binded fields
		postFromDB.MergeInPlace(postJSON)

		// Updating user's record in DB
		err = db.C(models.CollectionPost).UpdateId(bson.ObjectIdHex(id), postFromDB)
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, nil)
			} else {
				c.Error(err)
			}
			return
		}
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.Error(errors.New("Invalid post's id"))
	}
}

// Delete an user
func PostDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	if id != "" && bson.IsObjectIdHex(id) {
		db := c.MustGet("db").(*mgo.Database)
		err := db.C(models.CollectionPost).RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, nil)
			} else {
				c.Error(err)
			}
			return
		}
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.Error(errors.New("Invalid post's id"))
	}
}
