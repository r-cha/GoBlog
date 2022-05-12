package controllers

import (
	"net/http"
	"r-cha/goblog/db"
	"r-cha/goblog/models"

	"github.com/gin-gonic/gin"
)

func (r routes) addPostsController(rg *gin.RouterGroup) {
	controller := rg.Group("/posts")

	controller.POST("/", CreatePost)
	controller.GET("/", ListPosts)
	controller.GET("/:id", GetPost)
	controller.PATCH("/:id", UpdatePost)
}

func CreatePost(c *gin.Context) {
	// Validate Input
	var input models.CreatePost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new post
	post := models.Post{
		Title:  input.Title,
		Author: input.Author,
		Text:   input.Text,
	}
	db.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func ListPosts(c *gin.Context) {
	var posts []models.Post
	db.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post

	if err := db.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	// Check if it exists
	var post models.Post
	if err := db.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate Input
	var input models.UpdatePost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the record
	db.DB.Model(&post).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": post})
}
