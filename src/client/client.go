package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = []byte("myimportantsecret")

func getEnvVariable(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

func GenerateJWT() (string, error) {

	key := getEnvVariable("JWT_TKN_0")
	mySigningKey := []byte(key)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "tinolebat"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func main() {
	fmt.Println("Client")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Errorf("Error generating token: %s", err.Error())
	}

	fmt.Println(tokenString)
}
