package monitor_mutex

import (
	"context"
	"fmt"
)

func NewCounterMnitor(ctx context.Context) chan<- int {
	ch := make(chan int)

	go func() {
		counter := 0

		for {
			select {
			case i := <-ch:
				counter += i
			case <-ctx.Done():
				fmt.Printf("final_count: %d\n", counter)
				return
			}
		}
	}()

	return ch
}
