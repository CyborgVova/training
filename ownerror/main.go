package main

import (
	"fmt"
	"log"
)

type MyError string

const (
	ErrorNotFound   MyError = "not found"
	ErrorOutOfRange MyError = "out of range"
)

func (e MyError) Error() string {
	return string(e)
}

func CheckOwnError(arr []int, idx int) (int, error) {
	if idx >= len(arr) {
		return 0, ErrorOutOfRange
	}
	return arr[idx], nil
}

func main() {
	var err error
	err = ErrorOutOfRange
	err = ErrorNotFound

	arr := []int{1, 2, 3, 4, 5, 6}
	item, err := CheckOwnError(arr, 6)
	if err != nil {
		log.Fatal("check own error:", err)
	}
	fmt.Println(item)
}
