package main

import (
	"fmt"
	"sort"
)

func main() {
	//unsortmap()
	sortmap()
}

func unsortmap() {
	// Generates items map with 10 elements.
	items := make(map[int]string)
	for i := 0; i < 10; i++ {
		items[i] = fmt.Sprintf("This is item %d", i)
	}

	// Processes items from map
	for _, item := range items {
		performItem(item)
	}
}

func performItem(item string) {
	fmt.Println(item)
}

func sortmap() {
	items := make(map[int]string)
	for i := 0; i < 10; i++ {
		items[i] = fmt.Sprintf("This is item %d", i)
	}

	// Generates keys slice.
	keys := make([]int, len(items))
	for k := range items {
		keys[k] = k
	}

	sort.Ints(keys)

	// Iterates over items map, using sorted keys.
	for _, k := range keys {
		performItem(items[k])
	}
}
