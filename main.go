package main

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	postGroup := router.Group("/posts")
	api.Posts(postGroup)
	router.Run()
}
