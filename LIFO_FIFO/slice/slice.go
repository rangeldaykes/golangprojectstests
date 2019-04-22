// Warning: The memory allocated for the array is never returned.
// For a long-living queue you should probably use a dynamic data
//  structure, such as a linked list.
package main

import "fmt"

func main() {
	var queue []string
	queue = append(queue, "Hello ")
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Println(queue[0])
		queue = queue[1:]
	}
}
