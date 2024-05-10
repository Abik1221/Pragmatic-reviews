package main

import (
	"github.com/gin-gonic/gin"
"github.com/dummy/gin-logger"
"github.com/dummy/gin-recovery"
"github.com/dummy/gin-gonic/gin/binding"
"github.com/dummy/gin-gonic/gin/render",
)

func main() {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hellow world",
		},
		)
	})

	server.Run(":8080")
}
