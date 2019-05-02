package db

import (
	"fmt"

	"golang.org/x/xerrors"
)

type Mydb struct {
}

//var ErrNotFound = errors.New("mydb: key not found")

func (md Mydb) Get(key string, val interface{}) error {
	return KeyNotFoundError{Name: key}
}

func (md Mydb) AccessCheck(key string) error {
	var val string
	if err := md.Get(key, &val); err != nil {
		// AccessCheck change the type of error that was previously KeyNotFoundError
		//return fmt.Errorf("access check: %v", err)

		// Solution with golang.org/x/xerrors
		return xerrors.Errorf("access check: %w", err)
	}
	if val != "AccessGranted" {
		fmt.Println("access")
	}
	return nil
}

type KeyNotFoundError struct {
	Name string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Errorf("taildb: key %q not found", e.Name).Error()
}
