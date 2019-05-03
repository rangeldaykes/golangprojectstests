package main

import (
	"fmt"
	"strconv"
)

type predicate func(i int) bool

func main() {
	m := map[string]predicate{
		"Fizz": func(i int) bool { return i%3 == 0 },
		"Buzz": func(i int) bool { return i%5 == 0 },
	}

	fmt.Println(fizzbuzz(1, 100, m))
}

func fizzbuzz(ini int, end int, interval map[string]predicate) []string {
	var result []string
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
		result = append(result, text)
	}
	return result
}
