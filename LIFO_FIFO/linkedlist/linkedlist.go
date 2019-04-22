package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushBack("Hello ") // Enqueue
	queue.PushBack("world!")

	for queue.Len() > 0 {
		e := queue.Front() // First element
		fmt.Print(e.Value)

		queue.Remove(e) // Dequeue
	}
}
