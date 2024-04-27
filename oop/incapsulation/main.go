package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	FirstName string
	LastName  string
	Login     string
	password  string
}

func (u *User) JSONData() (string, error) {
	b, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return "", fmt.Errorf("impossible serialized struct: %v", err)
	}
	return string(b), nil
}

func main() {
	user := &User{
		FirstName: "Vova",
		LastName:  "Ivanov",
		Login:     "admin",
		password:  "password",
	}

	b, err := user.JSONData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
