package db

import "fmt"

type Mydb struct {
}

//var ErrNotFound = errors.New("mydb: key not found")

func (md Mydb) Get(key string, val interface{}) error {
	return KeyNotFoundError{Name: key}
}

type KeyNotFoundError struct {
	Name string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Errorf("taildb: key %q not found", e.Name).Error()
}
