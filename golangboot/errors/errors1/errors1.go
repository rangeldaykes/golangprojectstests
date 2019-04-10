package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	//r := calc(3, "a")
	//fmt.Println(r)
	//r2 := calc(3, "3")
	//fmt.Println(r2)

	//b()
	fmt.Println(c())

	fmt.Println("Fim!!!")
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func calc(a int, b string) int {
	r := convert(b)
	return a * r
}

func convert(b string) int {

	r, err := strconv.Atoi(b)
	if err != nil {
		//log.Panicln(err)
		log.Fatalln(err)

		//log.Println(err)

		//fmt.Println(err)
		//var e = fmt.Errorf("errouuu - ", err)
		//fmt.Println(e)
	}
	return r
}
