package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("Even numbers up to 8:\n")
	t1 := time.Now()

	printEvenNumbers(9999)

	t2 := time.Now()
	diff := t2.Sub(t1)
	diffmili := int64(diff / time.Millisecond)
	fmt.Println(diffmili)
}

func printEvenNumbers(max int) {
	err := iterateEvenNumbers(max, func(n int) error {
		//fmt.Printf("%d\n", n)
		return nil
	})
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}

func iterateEvenNumbers(max int, cb func(n int) error) error {
	if max < 0 {
		return fmt.Errorf("max is %d, must be >= 0", max)
	}
	for i := 2; i < max; i += 2 {
		err := cb(i)
		if err != nil {
			return err
		}
	}
	return nil
}
