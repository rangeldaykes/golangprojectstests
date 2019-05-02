package main

import (
	"errors_v3_type_error_intermediary/db"
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	d := db.Mydb{}
	var val string

	// AccessCheck change the type of error that was previously KeyNotFoundError
	err := d.AccessCheck("my-key")
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
	// AccessCheck change the type of error that was previously KeyNotFoundError

	// Solution with golang.org/x/xerrors
	if err := d.AccessCheck("my-key"); err != nil {
		var notFoundErr db.KeyNotFoundError
		if xerrors.As(err, &notFoundErr) {
			fmt.Println("no value")
			fmt.Println(err)
		} else {
			fmt.Println("other error")
		}
	} else {
		fmt.Println(val)
	}
	// Solution with golang.org/x/xerrors
}
