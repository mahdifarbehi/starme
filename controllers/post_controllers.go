package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdifarbehi/starme/initializers"
	"github.com/mahdifarbehi/starme/models"
)

func PostCreate(c *gin.Context) {
	type PostCreateInput struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	var input PostCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Title: input.Title, Content: input.Content}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func PostRead(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.Status(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func PostReadAll(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.Status(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func PostUpdate(c *gin.Context) {
	type PostUpdateInput struct {
		Title   string `json:"title"   binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	var input PostUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var post models.Post
	if result := initializers.DB.First(&post, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	post.Title = input.Title
	post.Content = input.Content

	if result := initializers.DB.Save(&post); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if result := initializers.DB.First(&post, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	if result := initializers.DB.Delete(&post); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}
