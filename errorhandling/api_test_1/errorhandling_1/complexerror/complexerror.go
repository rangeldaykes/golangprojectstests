package complexerror

import (
	"fmt"

	"golang.org/x/xerrors"
)

type ComplexError struct {
	Message string
	Code    int
	Frame   xerrors.Frame
}

func (ce ComplexError) FormatError(p xerrors.Printer) error {
	p.Printf("%d %s", ce.Code, ce.Message)
	ce.Frame.Format(p)
	return nil
}

func (ce ComplexError) Format(f fmt.State, c rune) {
	xerrors.FormatError(ce, f, c)
}

func (ce ComplexError) Error() string {
	return fmt.Sprint(ce)
}
