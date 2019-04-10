package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//nextInt := intSeq()
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//newInts := intSeq()
	//fmt.Println(newInts())

	// tests with struct ////////////////////////
	// opt := DialDatabase(11)
	// do := dialOptions{}
	// fmt.Printf("\n %p", &do)
	// opt.f(&do)
	// fmt.Printf("\n %p", &do)
	// fmt.Print("\n")
	// fmt.Println(do.db)
	/////////////////////////////////////////////

	opt := DialDatabase2(11)
	var do int
	do = opt.f()
	fmt.Println(do)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type DialOption2 struct {
	f func() int
}

func DialDatabase2(db int) DialOption2 {
	op := DialOption2{}

	op.f = func() int {
		return db
	}

	return op
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type DialOption struct {
	f func(*dialOptions)
}

type dialOptions struct {
	db int
}

func DialDatabase(db int) DialOption {

	op := DialOption{}

	op.f = func(do *dialOptions) {
		fmt.Printf("\n %p", do)
		do.db = db
		fmt.Printf("\n %p", do)
	}

	return op
}
