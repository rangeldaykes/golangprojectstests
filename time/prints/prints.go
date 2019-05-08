package main

import (
	"fmt"
	"time"
)

type JsonIso8601 time.Time

func main() {
	//var t time.Time

	var t2 JsonIso8601
	fmt.Println(t2)
}
