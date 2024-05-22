package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Next  *Node
	Value int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (l *LinkedList) PushBack(val int) {
	if l.Head == nil {
		l.Head = &Node{nil, val}
		return
	}
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{nil, val}
}

func (l *LinkedList) Print() {
	tmp := l.Head
	for tmp != nil {
		fmt.Print(tmp.Value, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}

func (list *LinkedList) ReverseLinkedList() {
	arr := []int{}
	tmp := list.Head
	for tmp != nil {
		arr = append(arr, tmp.Value)
		tmp = tmp.Next
	}

	tmp = list.Head
	for i := len(arr) - 1; i >= 0; i-- {
		tmp.Value = arr[i]
		tmp = tmp.Next
	}
}

func (list *LinkedList) Reverse2LinkedList() {
	var prev, next *Node = nil, nil
	cursor := list.Head

	for cursor != nil {
		next, cursor.Next = cursor.Next, prev
		prev, cursor = cursor, next
	}

	list.Head = prev
}

func main() {
	list := NewLinkedList()
	list.PushBack(50)
	list.PushBack(100)
	list.PushBack(150)
	list.PushBack(200)
	list.Print()

	list.ReverseLinkedList()
	list.Print()

	list.Reverse2LinkedList()
	list.Print()
}
