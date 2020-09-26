package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = []byte("myimportantsecret")

// GetEnvVariable Get the Environment variables
func GetEnvVariable(key string) string {

	err := godotenv.Load("../config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

func generateJWT() (string, error) {

	key := GetEnvVariable("JWT_TKN_0")
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
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Client")
	handleRequest()

}
