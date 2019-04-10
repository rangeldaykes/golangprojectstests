package main

import (
	"fmt"
)

func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

func tryAssertInt(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	var s interface{} = 56
	assert(s)

	var s2 interface{} = 56
	tryAssertInt(s2)

	var s3 interface{} = "56"
	tryAssertInt(s3)

	findType("Naveen")
	findType(77)
	findType(89.98)
}
