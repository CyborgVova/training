package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	var age int
	val, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err == nil {
		age = val
	}
	b, err := json.MarshalIndent(Human{Name: name, Age: age}, "", "    ")
	if err != nil {
		log.Printf("Marshal error: %v\n", err)
	}
	fmt.Fprint(w, string(b))
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req: %s, path: %s ", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	}
}

var credentials = map[string]string{"Vova": "123"}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if ok && credentials[user] == password {
			next.ServeHTTP(w, r)
		} else {
			log.Printf("status: %d 'user=%s password=%s'\n", http.StatusUnauthorized, user, password)
			fmt.Fprintf(w, "%v Not authorized", http.StatusUnauthorized)
		}
	}
}

func main() {
	http.HandleFunc("/hello", Logger(Auth(Hello)))

	fmt.Println("Server listen 8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
