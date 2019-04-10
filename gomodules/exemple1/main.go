package main

import (
	"exemple1/configuration"
	"fmt"

	"github.com/donvito/hellomod"
)

func main() {

	fmt.Println("Start gomodules example1")

	hellomod.SayHello()

	fmt.Println(configuration.Calculate(2))

	fmt.Scanln()
}
