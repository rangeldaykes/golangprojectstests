package main

import (
	"errors_v4_type_error_final/db"
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	d := db.Mydb{}
	var val string

	if err := d.AccessCheck("my-key"); xerrors.Is(err, db.ErrNotFound) {
		fmt.Println("no value")
		fmt.Println(err)
	} else if err != nil {
		fmt.Println("other error")
	} else {
		fmt.Println(val)
	}
}
