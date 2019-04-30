package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Size: %d MB.", 85)
	s := buf.String()
	fmt.Println(s)
}
