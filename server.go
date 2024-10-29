package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hellow world",
		},
		)
	})

	server.Run(":8080")

}
