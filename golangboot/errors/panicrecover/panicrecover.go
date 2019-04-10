package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")

	g(0)
	g(1)
	g(2)
	g(3)
	g(4)

	fmt.Println("Returned normally from g.")
}

func g(i int) {

	if i > 3 {
		//fmt.Println("Panicking!")
		log.Panicln("deu ruim!")
		//panic("deu ruim") //(fmt.Sprintf("%v", i))
	}

	//defer fmt.Println("Defer in g", i)

	fmt.Println("Printing in g", i)

	//g(i + 1)
}
