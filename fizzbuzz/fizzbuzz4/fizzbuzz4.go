package main

import (
	"fmt"
	"strconv"
)

func main() {
	var keys = map[string]int{
		"Fizz": 3,
		"Buzz": 5,
	}

	fizzbuzz(1, 15, keys)
}

func fizzbuzz(ini int, end int, triggers map[string]int) {
	for i := ini; i <= end; i++ {
		var text string
		for k, v := range triggers {
			if i%v == 0 {
				text += k
			}
		}
		if text == "" {
			text = strconv.Itoa(i)
		}
		fmt.Println(text)
	}
}
