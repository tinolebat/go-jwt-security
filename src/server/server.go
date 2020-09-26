package main

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// GetEnvVariable Get the Environment variables
func getEnvVariable(key string) string {
	err := godotenv.Load("../config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

func homepage(c *gin.Context) {
	fmt.Println("Welcome to homepage")
	c.JSON(200, gin.H{
		"message": "Welcome",
	})
}

func isAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := getEnvVariable("JWT_TKN_0")
		mySigningKey := []byte(key)
		if c.GetHeader("Token") != "" {

			token, err := jwt.Parse(c.GetHeader("Token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There is an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Println("Token parsing error")
			}

			if token.Valid {
				endpoint(c)
			}

		} else {
			fmt.Println("Not Authorized")
			w := "Not Authorized"
			c.JSON(200, w)
		}
	}

}

func handleRequest() {
	router := gin.Default()
	router.GET("/", isAuthorized(homepage))
	router.Run()
}

func main() {
	fmt.Println("Server side")
	handleRequest()

}
