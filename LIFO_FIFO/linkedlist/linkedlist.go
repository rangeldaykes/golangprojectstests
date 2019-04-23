package main

import (
	"container/list"
	"fmt"
)

var queue *list.List

func main() {
	queue = list.New()
	queue.PushBack("Hello ") // Enqueue
	queue.PushBack("world!")
	queue.PushBack("world!")

	printList()
	e := queue.Back()
	queue.Remove(e)
	printList()
}

func printList() {
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
