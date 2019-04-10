package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	//for i := 0; i < len(s); i++ {
	//fmt.Printf("%c ", s[i])
	//}

	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
