package api

import (
	repo "blog/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Posts(group *gin.RouterGroup) {
	group.GET("/", getPosts)
	group.POST("/", createPost)
	group.DELETE("/:id", deletePost)
}

func getPosts(c *gin.Context) {
	posts := repo.Get()
	c.JSON(200, posts)
}

func createPost(c *gin.Context) {
	var input repo.CreatePostInput
	if c.ShouldBindBodyWithJSON(&input) != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	post := repo.Create(input)
	c.JSON(201, post)
}

func deletePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	repo.Delete(postID)
	c.JSON(204, nil)
}
