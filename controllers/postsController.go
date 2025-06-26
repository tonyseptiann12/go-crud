package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyseptiann12/go-crud/initializer"
	"github.com/tonyseptiann12/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializer.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializer.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializer.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off url
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializer.DB.First(&post, id)

	// Update it
	initializer.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializer.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
