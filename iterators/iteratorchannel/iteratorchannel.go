package main

import (
	"context"
	"fmt"
	"log"
)

type IntWithError struct {
	Int int
	Err error
}

func generateEvenNumbers(max int) chan IntWithError {
	ch := make(chan IntWithError)
	go func() {
		defer close(ch)
		if max < 0 {
			ch <- IntWithError{
				Err: fmt.Errorf("'max' is %d and should be >= 0", max),
			}
			return
		}

		for i := 2; i <= max; i += 2 {
			ch <- IntWithError{
				Int: i,
			}
		}
	}()
	return ch
}

func printEvenNumbers(max int) {
	for val := range generateEvenNumbers(max) {
		if val.Err != nil {
			log.Fatalf("Error: %s\n", val.Err)
		}
		fmt.Printf("%d\n", val.Int)
	}
}

func main() {
	fmt.Printf("Even numbers up to 8:\n")
	printEvenNumbers(8)

	fmt.Printf("Even numbers up to 20, cancel at 8:\n")
	printEvenNumbersCancellable(20, 8)
}

func generateEvenNumbersWithCtx(ctx context.Context, max int) chan IntWithError {
	ch := make(chan IntWithError)
	go func() {
		defer close(ch)
		if max < 0 {
			ch <- IntWithError{
				Err: fmt.Errorf("'max' is %d and should be >= 0", max),
			}
			return
		}

		for i := 2; i <= max; i += 2 {
			if ctx != nil {
				// if context was cancelled, we stop early
				select {
				case <-ctx.Done():
					return
				default:
				}
			}
			ch <- IntWithError{
				Int: i,
			}
		}
	}()
	return ch
}

func printEvenNumbersCancellable(max int, stopAt int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := generateEvenNumbersWithCtx(ctx, max)
	for val := range ch {
		if val.Err != nil {
			log.Fatalf("Error: %s\n", val.Err)
		}
		if val.Int > stopAt {
			cancel()
			// notice we keep going in order to drain the channel
			continue
		}
		// process the value
		fmt.Printf("%d\n", val.Int)
	}
}
