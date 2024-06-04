package main

import "fmt"

func main() {
	c := []string{"A", "B", "D", "E"}
	b := c[1:2]
	b = append(b, "TT")
	fmt.Println(c) // [A B TT E] len = 2 cap = 4
	fmt.Println(b) // [B TT] len = 2 cap = 3
}
