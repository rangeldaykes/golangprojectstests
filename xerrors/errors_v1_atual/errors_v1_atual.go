package main

import (
	"errors_v1_atual/db"
	"fmt"
)

func main() {
	d := db.Mydb{}
	var val string

	if err := d.Get("my-key", &val); err == db.ErrNotFound {
		fmt.Println("no value")
		fmt.Println(err)
	} else if err != nil {
		fmt.Println("orther error")
	} else if err != nil {
		fmt.Println(val)
	}
}
