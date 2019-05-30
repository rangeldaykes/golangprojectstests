package monitor_mutex_test

import (
	"context"
	"fmt"
	"monitor_mutex"
	"sync"
	"testing"
)

func TestNewCounterMonitor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := monitor_mutex.NewCounterMnitor(ctx)
	in <- 1
}

func TestMutex(t *testing.T) {
	sc := &monitor_mutex.SafeCounter{
		Mu: &sync.Mutex{},
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("sc.count = ", sc.Count)
	// sc.count = 10
}
