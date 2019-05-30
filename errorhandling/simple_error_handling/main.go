// https://medium.com/gett-engineering/error-handling-in-go-53b8a7112d04
package main

import (
	"errors"
	"fmt"
	"time"

	//"golang.org/x/xerrors"
	pkgerrors "github.com/pkg/errors"
)

func main() {
	e1 := errors.New(fmt.Sprintf("Could not open file"))
	e2 := fmt.Errorf("Could not open file")

	fmt.Println(fmt.Sprintf("Type of error 1: %T", e1))
	fmt.Println(fmt.Sprintf("Type of error 2: %T", e2))

	ret, err := process(4, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println(pkgerrors.Cause(err))
	}
	fmt.Println("ret = ", ret)
}

func process(a int, b int) (float64, error) {
	fmt.Println("processing")
	time.Sleep(1 * time.Second)
	ret, err := divide(a, b)
	if err != nil {
		return 0, pkgerrors.Wrap(err, "process problem")
	}

	return ret, nil
}

func divide(a int, b int) (float64, error) {
	if b <= 0 {
		return 0, fmt.Errorf("zeroooo")
	} else {
		return float64(a) / float64(b), nil
	}
}
