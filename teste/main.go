package main

import "fmt"

func main() {
	logPrefix := "aabbc"
	logPrefix = logPrefix[2:]

	fmt.Println(logPrefix)
}
