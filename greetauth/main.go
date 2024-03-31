package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var Database = map[string]string{}

func NewClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, password, ok := r.BasicAuth()
	if !ok || user == "" || password == "" {
		log.Println("bad data: user or password", r.URL.Host, ":", r.URL.Port(), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := Database[user]; ok {
		log.Println("such user exist yet")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("error generate password", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Database[user] = string(hashPass)
	b := []byte(user)
	b = append(b, byte(':'))
	b = append(b, []byte(hashPass)...)
	w.Write(b)
}

func Start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, password, ok := r.BasicAuth()
	if !ok {
		log.Println("bad authorization")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hashPass, ok := Database[user]
	if !ok {
		log.Println("go to http://localhost:8080/signup")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		log.Println("password is wrong")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Hello, %s !!!", user)
}

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/signup", NewClient)
	fmt.Println("Server on port :8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("server not started:", err)
	}
}
