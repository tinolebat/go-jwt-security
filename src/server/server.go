package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func homepage(c *gin.Context) {
	fmt.Println("Welcome to homepage")
	c.JSON(200, gin.H{
		"message": "Welcome",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Welcome to homepage")
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})

	fmt.Println("Hi")
	router.Run()
}
