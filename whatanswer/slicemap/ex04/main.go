package main

import (
	"fmt"
)

func main() {
	var m map[string]int // нельзя писать в неинициализированную карту
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++ // тут будет паника
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
