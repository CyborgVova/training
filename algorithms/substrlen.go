package main

import "fmt"

func FindIndex(arr []rune, elem rune) int {
	for i := range arr {
		if arr[i] == elem {
			return i
		}
	}
	return -1
}

func SubStrLen(slice []rune) int {
	length := len(slice)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return 1
	}

	res, tmp, idx := 0, 0, 0
	for i := 1; i < length; i++ {
		s := slice[idx:i]
		tmp = len(s)
		if find := FindIndex(s, slice[i]); find != -1 {
			idx = find + 1
		}
		if res < tmp {
			res = tmp
		}
	}

	return res
}

// 'a', 'b', 'c', 'b', 'a', 'd', 'a'

func main() {
	a := []rune{'a', 'b', 'c', 'b', 'a', 'd', 'a'}
	b := []rune{'a', 'x', 'b', 'x', 'c', 'x', 'd'}
	c := []rune{'a', 'a', 'a', 'a', 'a', 'a', 'a'}
	d := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	fmt.Println(SubStrLen(a))
	fmt.Println(SubStrLen(b))
	fmt.Println(SubStrLen(c))
	fmt.Println(SubStrLen(d))
}
