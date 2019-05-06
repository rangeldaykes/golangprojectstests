package main

import "fmt"

type StatefulIterator interface {
	Value() int
	Next() bool
	Err() error
}

type EvenNumberIterator struct {
	max       int
	currValue int
	err       error
}

func NewEvenNumberIterator(max int) *EvenNumberIterator {
	var err error
	if max < 0 {
		err = fmt.Errorf("'max' is %d, should be >= 0", max)
	}
	return &EvenNumberIterator{
		max:       max,
		currValue: 0,
		err:       err,
	}
}

func (i *EvenNumberIterator) Next() bool {
	if i.err != nil {
		return false
	}
	i.currValue += 2
	return i.currValue <= i.max
}

func (i *EvenNumberIterator) Value() int {
	if i.err != nil || i.currValue > i.max {
		panic("Value is not valid after iterator finished")
	}
	return i.currValue
}

func (i *EvenNumberIterator) Err() error {
	return i.err
}
