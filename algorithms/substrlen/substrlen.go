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
	if length < 2 {
		return length
	}

	m := map[rune]int{slice[0]: 0}

	res, tmp := 0, 1
	for i := 1; i < length; i++ {
		if ii, ok := m[slice[i]]; ok {
			m = map[rune]int{}
			res = tmp
			tmp = 0
			i = ii
			continue
		}
		tmp++
		m[slice[i]] = i
		if res < tmp {
			res = tmp
		}
	}

	return res
}

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
