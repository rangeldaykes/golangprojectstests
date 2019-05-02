package db

import "errors"

type Mydb struct {
}

var ErrNotFound = errors.New("mydb: key not found")

func (md Mydb) Get(key string, val interface{}) error {
	return ErrNotFound
}
