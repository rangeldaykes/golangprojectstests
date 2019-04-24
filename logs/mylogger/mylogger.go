package main

import (
	"fmt"
	"mylogger/mylogger1"
	"time"
)

var log *mylogger1.Logger

func main() {

	log = mylogger1.NewLogger(
		mylogger1.NewFileLogger(
			"Logteste",
			"/home/rangelsantos/discod/golangprojectstests/"),
		mylogger1.LEVEL_LOG)

	t1 := time.Now()

	//time.Sleep(3 * time.Second)
	testLog()
	//time.Sleep(3 * time.Second)
	testLog()
	//time.Sleep(3 * time.Second)
	testLog()

	t2 := time.Now()
	diff := t2.Sub(t1)
	diffmili := int64(diff / time.Millisecond)
	fmt.Println(diffmili)

	fmt.Scanln()
}

func testLog() {
	for i := 0; i < 50000; i++ {

		log.Log("TestLogger1")

		//str := "ab"
		//i, err := strconv.Atoi(str)
		//if err != nil {
		//log.Error(err, "conversão ruin")
		//}
		//fmt.Println(i)

		log.Assert(i == 1, "i não tem o valor 1")

		log.ContextStart("start1")
		log.Log("after start context 1")

		log.ContextStart("start2")
		log.Log("after start context 2")
		log.ContextEnd()

		log.ContextEnd()

		log.Log("after end context 1")
	}
}
