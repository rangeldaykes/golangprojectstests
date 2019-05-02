package db

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

type Mydb struct {
}

var ErrNotFound = errors.New("mydb: key not found")

func (md Mydb) Get(key string, val interface{}) error {
	return xerrors.Errorf("mydb: %q: %w", key, ErrNotFound)
}

func (md Mydb) AccessCheck(key string) error {
	var val string
	if err := md.Get(key, &val); err != nil {
		return xerrors.Errorf("access check: %w", err)
	}
	if val != "AccessGranted" {
		fmt.Println("access")
	}
	return nil
}
