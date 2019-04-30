package main

import "fmt"

type Person struct {
	name []byte
}

func (p *Person) Write(data []byte) (n int, err error) {
	p.name = data
	return len(data), nil
}

func main() {
	b := []byte("Dave")
	person := Person{}
	fmt.Fprintf(&person, "%s", b)
	fmt.Printf("Person name: %s\n", person.name)
}
