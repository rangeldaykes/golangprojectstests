package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	result := make(chan int)

	// Create cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	// Make request in separate go routine
	go request(ctx, result)

	// Wait for 100 ms timeout and cancel the request with context
	select {
	case <-time.After(10000 * time.Millisecond):
		cancel()
		fmt.Println("Request has been canceled")
		fallback()
	case StatusCode := <-result:
		fmt.Println("Request has been done. Status Code : ", StatusCode)
	}
}

func request(ctx context.Context, result chan int) {
	req, _ := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	client.Transport = &http.Transport{}
	res, _ := client.Do(req)

	result <- res.StatusCode
}

func fallback() {
	fmt.Println("Fallback")
}
