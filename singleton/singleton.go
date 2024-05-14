package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	s string
}

var singleton *Singleton

func GetInstance() {
	singleton = &Singleton{s: "singleton"}
}

func main() {
	var once sync.Once

	once.Do(GetInstance)
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)

	once.Do(GetInstance)
	singleton.s = "Only Once"
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)

	once.Do(GetInstance)
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)
}
