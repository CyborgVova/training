package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	id    int64
	Human struct {
		FirstName string `json:"name"`
		LastName  string `json:"lastname"`
	} `json:"human"`
	Role string `json:"role"`
	City string `json:"city"`
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee: %s %s,\nRole: %s,\nCity: %s;\n", e.Human.FirstName, e.Human.LastName, e.Role, e.City)
}

func main() {
	employee := Employee{
		id: 23,
		Human: struct {
			FirstName string `json:"name"`
			LastName  string `json:"lastname"`
		}{
			FirstName: "Vladimir",
			LastName:  "Erokhin",
		},
		City: "Moscow",
		Role: "Developer",
	}

	b, err := json.MarshalIndent(employee, "", "	")
	if err != nil {
		log.Fatalf("marshall error: %v", err)
	}
	fmt.Println(string(b))
	fmt.Println()

	backoff := Employee{}
	json.Unmarshal(b, &backoff)
	fmt.Println(backoff)
}
