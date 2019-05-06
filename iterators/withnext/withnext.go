package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("Even numbers up to 8:\n")
	t1 := time.Now()

	printEvenNumbers(9999999)

	t2 := time.Now()
	diff := t2.Sub(t1)
	diffmili := int64(diff / time.Millisecond)
	fmt.Println(diffmili)
}

func printEvenNumbers(max int) {
	iter := NewEvenNumberIterator(max)
	for iter.Next() {
		//fmt.Printf("n: %d\n", iter.Value())
	}
	if iter.Err() != nil {
		log.Fatalf("error: %s\n", iter.Err())
	}
}
