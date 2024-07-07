package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var email = "email@email.ru"

const secretKey = "secret_key"

func main() {
	payload := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(2 * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("signed string:", t)

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Fatal("error get payload from token")
	}

	fmt.Println("email:", payload["sub"])
	fmt.Println("expire:", payload["exp"])

	tok, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("header:", tok.Header)
	integer, decimal := math.Modf(tok.Claims.(jwt.MapClaims)["exp"].(float64))
	fmt.Println("datetime to token expiration:", time.Unix(int64(integer), int64(decimal)))
	for k, v := range tok.Claims.(jwt.MapClaims) {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Println(tok.Valid)
	time.Sleep(time.Second * 5)
	tok1, err1 := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err1 != nil {
		log.Print(err1)
	}
	fmt.Println(tok1.Valid)
}
