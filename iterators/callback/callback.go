package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("Even numbers up to 8:\n")
	printEvenNumbers(8)
}

func printEvenNumbers(max int) {
	err := iterateEvenNumbers(max, func(n int) error {
		fmt.Printf("%d\n", n)
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
		time.Sleep(1 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
