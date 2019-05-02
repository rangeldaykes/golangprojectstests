package main

import (
	"errors_v2_type_error/db"
	"fmt"
)

func main() {
	d := db.Mydb{}
	var val string

	err := d.Get("my-key", &val)
	if err != nil {
		if _, isNotFound := err.(db.KeyNotFoundError); isNotFound {
			fmt.Println("no value")
			fmt.Println(err)
		} else {
			fmt.Println("other error")
		}
	} else {
		fmt.Println(val)
	}
}
