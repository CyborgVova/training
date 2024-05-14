package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	s string
}

var (
	singleton *Singleton
	once      sync.Once
)

func GetInstance() {
	once.Do(func() {
		singleton = &Singleton{s: "singleton"}
	})
}

func main() {
	GetInstance()
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)

	GetInstance()
	singleton.s = "Only Once"
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)

	GetInstance()
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)
}
