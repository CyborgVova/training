package main

import (
	"fmt"
)

func Palindrom(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func RecursivePalindrom(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}
	return RecursivePalindrom(str[1 : len(str)-1])
}

func main() {
	fmt.Println(Palindrom("ab  ba"))
	fmt.Println(RecursivePalindrom("ab ba"))
}
