package context5

import (
	"context"
	"time"
)

func main() {
	//ctx := context.Background()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	go Perform(ctx)

}

func Perform(ctx context.Context) error {
	for {
		SomeFunction(ctx)

		time.Sleep(1 * time.Second)
		select {

		case <-ctx.Done():
			// ctx is canceled
			return ctx.Err()

		case <-time.After(time.Second):
			// wait for 1 second
		}
	}
	return nil
}

fucn SomeFunction(ctx) {
	
	time.Sleep(3 * time.Second)
}