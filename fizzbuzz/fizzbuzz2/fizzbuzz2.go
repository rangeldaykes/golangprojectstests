// Dry dont repeat i%3 and i%5
package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 1; i <= 100; i++ {
		fizz := func() bool { return i%3 == 0 }
		buzz := func() bool { return i%5 == 0 }

		ret := func() string {
			switch {
			case fizz() && buzz():
				return "FizzBuzz"
			case fizz():
				return "Fizz"
			case buzz():
				return "Fizz"
			default:
				return strconv.Itoa(i)
			}
		}()

		fmt.Println(ret)
	}
}
