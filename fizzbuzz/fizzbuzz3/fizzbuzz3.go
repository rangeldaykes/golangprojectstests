// Constants Fizz - Buzz, and funciton divisibleby
package main

import (
	"fmt"
	"strconv"
)

const (
	FIZZ = "Fizz"
	BUZZ = "Buzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		fizz := divisibleby(i, 3)
		buzz := divisibleby(i, 5)

		ret := func() string {
			switch {
			case fizz && buzz:
				return FIZZ + BUZZ
			case fizz:
				return FIZZ
			case buzz:
				return BUZZ
			default:
				return strconv.Itoa(i)
			}
		}

		fmt.Println(ret())
	}
}

func divisibleby(numerator int, denominator int) bool {
	return numerator%denominator == 0
}
