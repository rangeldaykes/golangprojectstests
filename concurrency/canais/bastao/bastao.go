package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	time.Sleep(2 * time.Second)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)

	time.Sleep(2 * time.Second)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)

		go Runner(baton)
	}

	if runner == 4 {
		fmt.Printf("Runner %d Finished, RaceOver\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}
