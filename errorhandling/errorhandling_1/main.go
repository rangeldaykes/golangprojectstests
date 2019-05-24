package main

import (
	"errorhandling_1/complexerror"
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	cerr := someComplexErrorHappens()
	var originalErr complexerror.ComplexError
	if xerrors.As(cerr, &originalErr) {
		// deal with the complex error
		// we can now directly interrogate originalErr.Code
		// and originalErr.Message!

		fmt.Println(originalErr.Code)
		fmt.Println(originalErr.Message)
		fmt.Println(originalErr.Frame)
	}
}

func someComplexErrorHappens() error {
	complexErr := complexerror.ComplexError{
		Code:    1234,
		Message: "there was way too much tune",
		Frame:   xerrors.Caller(1), // skip the first frame)
	}

	return xerrors.Errorf("uh oh! something terribly complex happened: %w", complexErr)
}
