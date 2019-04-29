package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func (p *Person) Write(w io.Writer) {
	b, _ := json.Marshal(*p)
	w.Write(b)
}

func main() {
	me := Person{
		Id:   1,
		Name: "Me",
		Age:  64,
	}

	var b bytes.Buffer

	me.Write(&b)

	fmt.Printf("%s", b.String())
}
