package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Even numbers up to 8:\n")
	printEvenNumbers(8)
}

func printEvenNumbers(max int) {
	iter := NewEvenNumberIterator(max)
	for iter.Next() {
		fmt.Printf("n: %d\n", iter.Value())
	}
	if iter.Err() != nil {
		log.Fatalf("error: %s\n", iter.Err())
	}
}
