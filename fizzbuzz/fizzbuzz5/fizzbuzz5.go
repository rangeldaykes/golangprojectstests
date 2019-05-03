package main

import (
	"fmt"
	"strconv"
)

type predicate func(i int) bool

func main() {
	r := map[string]predicate{
		"Fizz": func(i int) bool { return i%3 == 0 },
		"Buzz": func(i int) bool { return i%5 == 0 },
		"Zazz": func(i int) bool { return i < 10 },
	}

	fizzbuzz(1, 100, r)
}

func fizzbuzz(ini int, end int, interval map[string]predicate) {
	for i := ini; i <= end; i++ {
		var text string
		for k, v := range interval {
			if v(i) {
				text += k
			}
		}
		if text == "" {
			text = strconv.Itoa(i)
		}
		fmt.Println(text)
	}
}
