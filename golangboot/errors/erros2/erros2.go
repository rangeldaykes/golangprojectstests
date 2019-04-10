package main

import "os"

func main() {
	f, err := os.Open("name.txt")

	if err != nil {
		return nil, err
	}
}
