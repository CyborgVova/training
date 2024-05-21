package main

import "fmt"

func Perm(arr []rune) {
	perm(arr, 0)
}

func perm(arr []rune, length int) {
	if length > len(arr) {
		fmt.Println(string(arr))
		return
	}
	perm(arr, length+1)
	for i := length + 1; i < len(arr); i++ {
		arr[i], arr[length] = arr[length], arr[i]
		perm(arr, length+1)
		arr[i], arr[length] = arr[length], arr[i]
	}
}

func main() {
	Perm([]rune("abc"))
	fmt.Println()
	Perm([]rune("abcd"))
}
