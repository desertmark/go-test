package api

import (
	"blog/models"
	repo "blog/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

var r = repo.NewPostPostgresRepository()

func Posts(group *gin.RouterGroup) {
	group.GET("/:id", getById)
	group.GET("/", getPosts)
	group.POST("/", createPost)
	group.DELETE("/:id", deletePost)
}

func getById(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	post := r.GetById(postID)
	if post.ID == 0 {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(200, post)
}

func getPosts(c *gin.Context) {
	posts := r.Get()
	c.JSON(200, posts)
}

func createPost(c *gin.Context) {
	var input models.CreatePostInput
	if c.ShouldBindBodyWithJSON(&input) != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	post := r.Create(input)
	c.JSON(201, post)
}

func deletePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	r.Delete(postID)
	c.JSON(204, nil)
}
