package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func recovery(x string) {
	if r := recover(); r != nil {
		fmt.Println(x, " - recovered:", r)
		debug.PrintStack()
	}
}

func a() {
	defer recovery("A")
	fmt.Println("Inside A")
	//panic("oh! A panicked")
}

func b() {
	defer recovery("B")
	fmt.Println("Inside B")

	n := []int{5, 7, 4}
	fmt.Println(n[3])

}

func main() {
	go b()
	go a()

	time.Sleep(1 * time.Second)
	fmt.Println("normally returned from main")
}
