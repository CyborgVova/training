package main

import "fmt"

type Inner struct{}

type Outer struct {
	Inner
}

func (i *Inner) Printer() {
	fmt.Println("Printer")
}

func main() {
	outer := Outer{}
	outer.Printer()
}
